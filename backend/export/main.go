package export

import (
	"context"
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cast"
	"models"
	"tools"
)

const (
	sheetOne   = "عملکرد بیمه مشترکین خانگی"
	sheetTwo   = "عملکرد بیمه مشترکین تجاری"
	sheetThree = "پرداختی بیمه مشترکین خانگی"
	sheetFour  = "پرداختی بیمه مشترکین تجاری"
	sheetFive  = "جدول کلی پرداختی ها "
	sheetSix   = "پرونده های مختومه (عدم تایید)"
	sheetSeven = "جدول تجمیعی رسیدگی خسارات"
)

type damageReport struct {
	companiesHouse  map[string]companyData
	companiesBiz    map[string]companyData
	DamageTypeAll   damagesType
	DamageTypeBiz   damagesType
	DamageTypeHouse damagesType
}

// Runner runs export
func Runner(resultFileName, stateList, from, to string) {
	resXlsx, err := excelize.OpenFile("./result.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	stateQuery := elastic.NewBoolQuery()
	if stateList != "" {
		sShould := elastic.NewBoolQuery()
		for _, s := range strings.Split(stateList, ",") {
			sShould.Should(
				elastic.NewTermQuery(
					"company_id.keyword", s),
			)
		}
		stateQuery.Must(sShould)
	}

	dateQuery := elastic.NewBoolQuery()
	if from != "" {
		dateQuery.Must(elastic.NewRangeQuery("casuality_date").Gte(from))
	}
	if to != "" {
		dateQuery.Must(elastic.NewRangeQuery("casuality_date").Lte(to))
	}

	reportData := damageReport{}

	qLocationHouse := elastic.NewBoolQuery().Must(
		elastic.NewBoolQuery().Should(
			elastic.NewTermQuery("location_usage", "1"),
		),
	).Must(elastic.NewExistsQuery("reference_code")).Must(stateQuery).Must(dateQuery)

	qLocationBiz := elastic.NewBoolQuery().Must(
		elastic.NewBoolQuery().Should(
			elastic.NewTermQuery("location_usage", "2"),
			elastic.NewTermQuery("location_usage", "3")),
	).Must(elastic.NewExistsQuery("reference_code")).Must(stateQuery).Must(dateQuery)

	qAll := elastic.NewBoolQuery().Must(
		elastic.NewBoolQuery().Should(
			elastic.NewTermQuery("location_usage", "1"),
			elastic.NewTermQuery("location_usage", "2"),
			elastic.NewTermQuery("location_usage", "3")),
	).Must(elastic.NewExistsQuery("reference_code")).Must(stateQuery).Must(dateQuery)

	reportData.companiesHouse = makeCompanisData(qLocationHouse)
	reportData.companiesBiz = makeCompanisData(qLocationBiz)
	reportData.DamageTypeHouse = makeGeneralReportData(qLocationHouse)
	reportData.DamageTypeBiz = makeGeneralReportData(qLocationBiz)
	reportData.DamageTypeAll = makeGeneralReportData(qAll)

	runSheetOne(resXlsx, sheetOne, reportData.companiesHouse, from, to)
	runSheetTwo(resXlsx, sheetTwo, reportData.companiesBiz, from, to)
	runSheetThree(resXlsx, sheetThree, reportData.companiesHouse, from, to)
	runSheetFour(resXlsx, sheetFour, reportData.companiesBiz, from, to)
	runSheetFive(resXlsx, sheetFive, reportData, from, to)
	runSheetSix(resXlsx, sheetSix, stateQuery, dateQuery, from, to)
	// runSheetSeven(resXlsx, sheetSeven, reportData, from, to)

	resXlsx.UpdateLinkedValue()

	mediaPrefix := tools.GetEnv("MEDIA_ROOT", "./media")
	err = resXlsx.SaveAs(mediaPrefix + "/storage/" + resultFileName)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func makeCompanisData(q *elastic.BoolQuery) map[string]companyData {
	companiesData := make(map[string]companyData)
	client := getEsClient()

	sumDamageAmountAgg := elastic.NewSumAggregation().Field("sum_damage_amount")

	statusAgg := elastic.NewTermsAggregation().
		Field("status.keyword").
		MinDocCount(0).
		Size(100).
		SubAggregation("status_sum_damage_amount", sumDamageAmountAgg)

	damageTypeAgg := elastic.NewTermsAggregation().
		Field("damage_type.keyword").
		Size(100).
		MinDocCount(0).
		SubAggregation("damage_type_sum_damage_amount", sumDamageAmountAgg).
		SubAggregation("status", statusAgg)

	companyAgg := elastic.NewTermsAggregation().
		Field("company_id.keyword").
		Size(50).
		MinDocCount(1).
		SubAggregation("damage_type", damageTypeAgg).
		SubAggregation("status", statusAgg)

	res, _ := client.Search().
		Index(elasticIndexName).
		Aggregation("companies", companyAgg).
		Aggregation("damage_type", damageTypeAgg).
		Query(q).
		Size(0).
		Do(context.Background())

	coms, _ := res.Aggregations.Terms("companies")
	for _, r := range coms.Buckets {
		curCom := r.Key.(string)
		cd := companyData{
			statusData: make(map[string]statusData),
		}

		dt, _ := r.Aggregations.Terms("damage_type")
		for _, t := range dt.Buckets {
			switch t.Key.(string) {
			case models.RequestDamageTypeInstrument:
				das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
				sumDage := cast.ToInt64(das.Value)

				cd.instrumentDamage = damageData{
					count:           t.DocCount,
					sumDamageAmount: sumDage,
					statusData:      getStatusAgg(t),
				}
			case models.RequestDamageTypeFiring:
				das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
				sumDage := cast.ToInt64(das.Value)
				cd.firingDamage = damageData{
					count:           t.DocCount,
					sumDamageAmount: sumDage,
					statusData:      getStatusAgg(t),
				}
			case models.RequestDamageTypeExplosion:
				das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
				sumDage := cast.ToInt64(das.Value)
				cd.explosionDamage = damageData{
					count:           t.DocCount,
					sumDamageAmount: sumDage,
					statusData:      getStatusAgg(t),
				}
			case models.RequestDamageTypeMedical:
				das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
				sumDage := cast.ToInt64(das.Value)
				cd.medicalDamage = damageData{
					count:           t.DocCount,
					sumDamageAmount: sumDage,
					statusData:      getStatusAgg(t),
				}
			case models.RequestDamageTypeDeath:
				das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
				sumDage := cast.ToInt64(das.Value)
				cd.deathDamage = damageData{
					count:           t.DocCount,
					sumDamageAmount: sumDage,
					statusData:      getStatusAgg(t),
				}
			case models.RequestDamageTypeLack:
				das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
				sumDage := cast.ToInt64(das.Value)
				cd.lackDamage = damageData{
					count:           t.DocCount,
					sumDamageAmount: sumDage,
					statusData:      getStatusAgg(t),
				}
			}
		}

		ds, _ := r.Aggregations.Terms("status")
		for _, t := range ds.Buckets {
			das, _ := t.Aggregations.Sum("status_sum_damage_amount")
			sumDage := cast.ToInt64(das.Value)
			cd.statusData[t.Key.(string)] = statusData{
				name:            t.Key.(string),
				count:           t.DocCount,
				sumDamageAmount: sumDage,
			}
		}

		// log.Println(curCom)
		companiesData[curCom] = cd
	}

	return companiesData
}

func makeGeneralReportData(q *elastic.BoolQuery) damagesType {
	damageTypeRes := damagesType{}
	client := getEsClient()

	sumDamageAmountAgg := elastic.NewSumAggregation().Field("sum_damage_amount")

	statusAgg := elastic.NewTermsAggregation().
		Field("status.keyword").
		MinDocCount(0).
		Size(100).
		SubAggregation("status_sum_damage_amount", sumDamageAmountAgg)

	damageTypeAgg := elastic.NewTermsAggregation().
		Field("damage_type.keyword").
		Size(100000).
		MinDocCount(0).
		SubAggregation("damage_type_sum_damage_amount", sumDamageAmountAgg).
		SubAggregation("status", statusAgg)

	res, _ := client.Search().
		Index(elasticIndexName).
		Aggregation("damage_type", damageTypeAgg).
		Query(q).
		Size(0).
		Do(context.Background())

	dt, _ := res.Aggregations.Terms("damage_type")
	for _, t := range dt.Buckets {
		switch t.Key.(string) {
		case models.RequestDamageTypeInstrument:
			das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
			sumDage := cast.ToInt64(das.Value)

			damageTypeRes.instrumentDamage = damageData{
				count:           t.DocCount,
				sumDamageAmount: sumDage,
				statusData:      getStatusAgg(t),
			}
		case models.RequestDamageTypeFiring:
			das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
			sumDage := cast.ToInt64(das.Value)
			damageTypeRes.firingDamage = damageData{
				count:           t.DocCount,
				sumDamageAmount: sumDage,
				statusData:      getStatusAgg(t),
			}
		case models.RequestDamageTypeExplosion:
			das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
			sumDage := cast.ToInt64(das.Value)
			damageTypeRes.explosionDamage = damageData{
				count:           t.DocCount,
				sumDamageAmount: sumDage,
				statusData:      getStatusAgg(t),
			}
		case models.RequestDamageTypeMedical:
			das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
			sumDage := cast.ToInt64(das.Value)
			damageTypeRes.medicalDamage = damageData{
				count:           t.DocCount,
				sumDamageAmount: sumDage,
				statusData:      getStatusAgg(t),
			}
		case models.RequestDamageTypeDeath:
			das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
			sumDage := cast.ToInt64(das.Value)
			damageTypeRes.deathDamage = damageData{
				count:           t.DocCount,
				sumDamageAmount: sumDage,
				statusData:      getStatusAgg(t),
			}
		case models.RequestDamageTypeLack:
			das, _ := t.Aggregations.Sum("damage_type_sum_damage_amount")
			sumDage := cast.ToInt64(das.Value)
			damageTypeRes.lackDamage = damageData{
				count:           t.DocCount,
				sumDamageAmount: sumDage,
				statusData:      getStatusAgg(t),
			}
		}
	}

	return damageTypeRes
}

func getSumAgg(t *elastic.AggregationBucketKeyItem, key string) int64 {
	das, ex := t.Aggregations.Sum(key)
	if !ex {
		return 0
	}
	return cast.ToInt64(das.Value)

}

func getStatusAgg(t *elastic.AggregationBucketKeyItem) map[string]statusData {
	sds := make(map[string]statusData)
	ts, ex := t.Aggregations.Terms("status")
	if !ex {
		panic("!ex")
	}

	for _, status := range ts.Buckets {
		sum, _ := status.Aggregations.Sum("status_sum_damage_amount")
		vsum := cast.ToInt64(sum.Value)

		sd := statusData{
			name:            status.Key.(string),
			count:           status.DocCount,
			sumDamageAmount: vsum,
		}
		sds[sd.name] = sd
	}

	return sds
}
