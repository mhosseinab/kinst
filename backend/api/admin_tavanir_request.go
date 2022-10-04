package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cast"

	"tavanir"

	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff"
	"models"
)

const (
	tavanirDamageTypeFiring     = "6"
	tavanirDamageTypeInstrument = "5"
)

func adminTavanirUpdateRequest(c *gin.Context) {
	id := c.Param("id")

	var m tavanir.Case

	d := db.Model(&tavanir.Case{})

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch || claim.Role == models.UserRoleTavanir {
			p := strings.Split(claim.State, ",")
			d = d.Where("company_id in (?)", p)
		}
	}

	if err := d.
		Find(&m, "tavanir_id=?", id).Error; err != nil {
		log.Println(err.Error())
		c.JSON(404, "404 page not found")
		return
	}

	beforeChange := m

	curStatus := m.Status
	//curExpertStatus := m.ExpertStatus

	if err := c.ShouldBind(&m); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error_msg": err.Error()})
		return
	}

	var autoAcceptMaxAmount int64 = 50_000_000
	if m.CompensationTypeID == tavanirDamageTypeFiring {
		autoAcceptMaxAmount = 400_000_000
	}
	if m.CompensationTypeID == tavanirDamageTypeInstrument {
		autoAcceptMaxAmount = 260_752_000
	}

	if m.TrackingID == "9717839566733" {
		autoAcceptMaxAmount = 5_400_000_000
	}
	//if m.Province == "خوزستان" {
	//	autoAcceptMaxAmount = 50_000_000
	//}

	// if user, ex := c.Get("user"); ex {
	// 	claim := user.(*KinsClaims)
	// 	if claim.Role == models.UserRoleBranch {
	// 		if curStatus != m.Status {
	// 			if m.ExpertAcceptedAmount > autoAcceptMaxAmount && m.Status == models.RequestStatusPayed {
	// 				c.JSON(400, gin.H{"error_msg": "برای این مبالغ امکان تغییر وضعیت توسط کاربران شعب وجود ندارد."})
	// 				return
	// 			}
	// 		}
	// 	}
	// }

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch {
			if beforeChange.AcceptedAmount != m.AcceptedAmount && beforeChange.AcceptedAmount > autoAcceptMaxAmount {
				c.JSON(400, gin.H{"error_msg": "برای این مبالغ امکان تغییر وضعیت توسط کاربران شعب وجود ندارد."})
				return
			}
		}
	}

	if beforeChange.AcceptedAmount == 0 && m.ExpertAcceptedAmount != 0 && m.ExpertAcceptedAmount < autoAcceptMaxAmount {
		m.AcceptedAmount = m.ExpertAcceptedAmount
	}

	if curStatus != m.Status {
		if m.Status == models.RequestStatusPayed && m.Sheba == "" {
			c.JSON(400, gin.H{"error_msg": "باید شبا وارد شده باشد"})
			return
		}

		//if !(m.DamageType == models.RequestDamageTypeDeath && m.Status == models.RequestStatusPayed) {
		//	sendSmsChaneStatus(m.MobileNumber)
		//}

	}

	// expert status changed
	//if curExpertStatus != m.ExpertStatus {
	//	m.ExpertUpdatedAt = time.Now()
	//	if m.ExpertStatus == models.RequestExpertStatusAccepted {
	//		if m.DamageType == models.RequestDamageTypeInstrument {
	//			if m.ExpertAcceptedAmount <= autoAcceptMaxAmount {
	//				m.Status = models.RequestStatusReadyToPay
	//				m.AcceptedAmount = m.ExpertAcceptedAmount
	//			}
	//		}
	//	}
	//}

	switch m.Status {
	case models.RequestStatusIncomplete:
		if m.MissingDocuments == "2000" {
			db.Model(&tavanir.StatusUpdateQueue{}).Save(&tavanir.StatusUpdateQueue{
				CaseID:    cast.ToUint64(m.TavanirID),
				NewStatus: tavanir.MethodPaymentInfoUpdateRequest,
			})
		} else {
			m.MissingDocuments = strings.Replace(m.MissingDocuments, "2000,", "", -1)
			m.MissingDocuments = strings.Replace(m.MissingDocuments, ",2000", "", -1)
			db.Model(&tavanir.StatusUpdateQueue{}).Save(&tavanir.StatusUpdateQueue{
				CaseID:    cast.ToUint64(m.TavanirID),
				NewStatus: m.Status,
			})
		}

		m.ExpertStatus = m.Status
		sendSmsTavanirChaneStatus(m.MobileNo)
		break

	case models.RequestStatusPayed:
		if m.AcceptedAmount == 0 {
			c.JSON(400, gin.H{"error_msg": "مبلغ خسارت تایید نشده است"})
			return
		}
		db.Model(&tavanir.StatusUpdateQueue{}).Save(&tavanir.StatusUpdateQueue{
			CaseID:    cast.ToUint64(m.TavanirID),
			NewStatus: m.Status,
		})
		m.ExpertStatus = m.Status
		sendSmsTavanirChaneStatus(m.MobileNo)
		break

	case models.RequestStatusRejected, models.RequestStatusReadyToPay, models.RequestStatusNeedToVisitInPerson:
		db.Model(&tavanir.StatusUpdateQueue{}).Save(&tavanir.StatusUpdateQueue{
			CaseID:    cast.ToUint64(m.TavanirID),
			NewStatus: m.Status,
		})
		m.ExpertStatus = m.Status
		sendSmsTavanirChaneStatus(m.MobileNo)
		break
	}

	if err := db.Save(&m).Error; err != nil {
		log.Println(err.Error())
	}

	changelog, _ := diff.Diff(beforeChange, m)
	if changelog != nil && len(changelog) > 1 {
		changelogByte, _ := json.Marshal(changelog)
		cl := tavanir.CaseChangelog{
			ChangeLogs: string(changelogByte),
			CreatedAt:  time.Now(),
			CaseID:     cast.ToUint(m.Id),
		}
		cl.UserID = getUserFromSession(c).UserID
		db.Create(&cl)
	}

	//if client, err := connections.GetElasticsearch(); err == nil {
	//	es.StoreRequestItem(client, m)
	//}

	c.JSON(http.StatusOK, m)
}

// adminRequestItems returns request items
func adminTavanirRequestItems(c *gin.Context) {
	var reqs tavanir.CaseSlice
	var count int
	offset, limit, order := getRequestFilters(c)
	d := db.Model(&tavanir.Case{})

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch || claim.Role == models.UserRoleTavanir {
			p := strings.Split(claim.State, ",")
			d = d.Where("company_id in (?)", p)
		}
	}

	filterKeys := []string{
		"status",
		"state_name",
		"city_name",
		"amount",
		"event_date",
		"tracking_id",
		"tavanir_id",
		"id",
		"compensation_type_id",
		"national_id",
		"reference_code",
		"bill_id",
		"expert_status",
	}

	for _, fk := range filterKeys {
		k := fmt.Sprintf("%s[]", fk)

		if values, ok := c.Request.URL.Query()[k]; ok {
			q := fmt.Sprintf("%s in (?)", fk)
			d = d.Where(q, values)
		}
	}

	if q, ok := c.Request.URL.Query()["query"]; ok {
		if isValidTavanirRawQuery(q[0]) {
			d = d.Where(q[0])
		}
	}

	err := d.Count(&count).
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&reqs).Error

	for i, c := range reqs {
		reqs[i].IsDuplicate = false
		var item models.Request
		db.Model(&models.Request{}).
			Where("status != ?", models.RequestStatusNotCompleted).
			Where(models.Request{
				BillIdentifier: c.BillID,
				NationalCode:   c.NationalID}).
			Find(&item)
		if item.ID != 0 {
			reqs[i].IsDuplicate = true
			log.Println(item.ID, item.BillIdentifier, item.NationalCode)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"objects": reqs,
		"meta": gin.H{
			"total_count": count,
			"err":         err,
		},
	})
}

func adminTavanirSyncQueueItems(c *gin.Context) {
	var reqs tavanir.StatusUpdateQueueSlice
	var count int
	offset, limit, order := getRequestFilters(c)
	d := db.Model(&tavanir.StatusUpdateQueue{})

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role != models.UserRoleAdmin {
			c.JSON(403, "403 unauthorized")
			return
		}
	}

	filterKeys := []string{
		"status",
		"case_id",
		"new_status",
		"success",
		"is_done",
	}

	for _, fk := range filterKeys {
		k := fmt.Sprintf("%s[]", fk)
		if values, ok := c.Request.URL.Query()[k]; ok {
			q := fmt.Sprintf("%s in (?)", fk)
			d = d.Where(q, values)
		}
	}

	err := d.Count(&count).
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&reqs).Error

	c.JSON(http.StatusOK, gin.H{
		"objects": reqs,
		"meta": gin.H{
			"total_count": count,
			"err":         err,
		},
	})
}

func adminTavanirMessageItems(c *gin.Context) {
	var reqs tavanir.MessageSlice
	var count int
	offset, limit, order := getRequestFilters(c)

	d := db.Model(&tavanir.MessageSlice{})

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role != models.UserRoleAdmin {
			c.JSON(403, "403 unauthorized")
			return
		}
	}

	filterKeys := []string{
		"message_id",
		"case_id",
		"ref_id",
		"date_time",
		"seen",
		"message",
	}

	for _, fk := range filterKeys {
		k := fmt.Sprintf("%s[]", fk)
		if values, ok := c.Request.URL.Query()[k]; ok {
			q := fmt.Sprintf("%s in (?)", fk)
			d = d.Where(q, values)
		}
	}

	err := d.Count(&count).
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&reqs).Error

	c.JSON(http.StatusOK, gin.H{
		"objects": reqs,
		"meta": gin.H{
			"total_count": count,
			"err":         err,
		},
	})
}

// adminRequestChangelog returns request items with changelog
func adminTavanirRequestChangelog(c *gin.Context) {
	var reqs tavanir.CaseChangelogSlice
	var count int
	offset, limit, order := getRequestFilters(c)
	d := db.Model(&tavanir.CaseChangelog{})

	filterKeys := []string{
		"case_id",
		"user_id",
		"created_at",
	}

	for _, fk := range filterKeys {
		k := fmt.Sprintf("%s[]", fk)
		if values, ok := c.Request.URL.Query()[k]; ok {
			q := fmt.Sprintf("%s in (?)", fk)
			d = d.Where(q, values)
		}
	}

	err := d.Count(&count).
		Limit(limit).
		Offset(offset).
		Order(order).
		Find(&reqs).Error

	for index := range reqs {
		var cd diff.Changelog
		err := json.Unmarshal([]byte(reqs[index].ChangeLogs), &cd)
		if err != nil {
			log.Println(err.Error())
		}
		reqs[index].ChangelogData = cd
	}

	c.JSON(http.StatusOK, gin.H{
		"objects": reqs,
		"meta": gin.H{
			"total_count": count,
			"err":         err,
		},
	})
}

// adminTavanirRequestSimilar returns similar requests by bill_id and national_id
func adminTavanirRequestSimilar(c *gin.Context) {
	id := c.Param("id")
	var item tavanir.Case

	d := db.Model(&tavanir.Case{})

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch || claim.Role == models.UserRoleTavanir {
			p := strings.Split(claim.State, ",")
			d = d.Where("company_id in (?)", p)
		}
	}

	if err := d.
		Select("tavanir_id, national_id, bill_id").
		Find(&item, "tavanir_id=?", id).Error; err != nil {
		log.Println(err.Error())
		c.JSON(404, "404 page not found")
		return
	}

	var items tavanir.CaseSlice
	if err := d.
		Where(&tavanir.Case{BillID: item.BillID, NationalID: item.NationalID}).
		Not(&tavanir.Case{TavanirID: item.TavanirID}).
		Find(&items).Error; err != nil {
		log.Println(err.Error())
		c.JSON(500, "500 internal server error")
		return
	}

	c.JSON(http.StatusOK, items)
}

// adminRequestItem returns request item
func adminTavanirRequestItem(c *gin.Context) {
	id := c.Param("id")
	var item tavanir.Case

	d := db.Model(&tavanir.Case{})

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch || claim.Role == models.UserRoleTavanir {
			p := strings.Split(claim.State, ",")
			d = d.Where("company_id in (?)", p)
		}
	}

	if err := d.
		Find(&item, "tavanir_id=?", id).Error; err != nil {
		log.Println(err.Error())
		c.JSON(404, "404 page not found")
		return
	}

	item.Documents, _ = getCaseDocs(item.Id)
	//var docs []tavanir.Document
	//db.Model(&tavanir.Document{}).Where("case_id=?", item.Id).Find(&docs)
	//for i := range docs {
	//	docs[i].Content = ""
	//	docs[i].FileType = ""
	//	docs[i].FileName = ""
	//}

	//item.Documents = docs

	c.JSON(http.StatusOK, item)
}

// adminRequestItem returns request item
func adminTavanirRequestItemById(c *gin.Context) {
	id := c.Param("id")
	var item tavanir.Case

	d := db.Model(&tavanir.Case{})

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch || claim.Role == models.UserRoleTavanir {
			p := strings.Split(claim.State, ",")
			d = d.Where("company_id in (?)", p)
		}
	}

	if err := d.
		Find(&item, "id=?", id).Error; err != nil {
		log.Println(err.Error())
		c.JSON(404, "404 page not found")
		return
	}

	item.Documents, _ = getCaseDocs(item.Id)
	//var docs []tavanir.Document
	//db.Model(&tavanir.Document{}).Where("case_id=?", item.Id).Find(&docs)
	//for i := range docs {
	//	docs[i].Content = ""
	//	docs[i].FileType = ""
	//	docs[i].FileName = ""
	//}

	//item.Documents = docs

	c.JSON(http.StatusOK, item)
}

// adminTavanirRequestDocuments returns request docs
func adminTavanirRequestDocuments(c *gin.Context) {
	id := c.Param("id")
	var doc tavanir.Document
	var cs tavanir.Case

	d := db.Model(&tavanir.Document{})

	if err := d.
		Find(&doc, "id=?", id).Error; err != nil {
		log.Println(err.Error())
		c.JSON(404, "404 page not found")
		return
	}

	if user, ex := c.Get("user"); ex {
		claim := user.(*KinsClaims)
		if claim.Role == models.UserRoleBranch || claim.Role == models.UserRoleTavanir {
			p := strings.Split(claim.State, ",")

			d = db.Model(&tavanir.Case{})
			d = d.Where("company_id in (?)", p)

			if err := d.
				Find(&cs, "id=?", doc.CaseId).Error; err != nil {
				log.Println(err.Error())
				c.JSON(406, "406 integrity failure")
				return
			}
		}
	}

	data, err := base64toFile(doc.Content)

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}

	c.Header("Content-Type", doc.FileType)
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", doc.FileName))
	c.Writer.Write(data)

}

func getCaseDocs(caseId int64) ([]tavanir.Document, error) {

	docs := make([]tavanir.Document, 0)

	err := db.
		Model(&tavanir.Document{}).
		Select("id, file_name, case_id, document_type_id, file_type").
		Where("case_id=?", caseId).
		Find(&docs).Error

	if err != nil {
		return docs, err
	}

	return docs, nil

}

//Given a base64 string of a File, writes it to a temp file
func base64toFile(data string) ([]byte, error) {

	d, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Println("[!] base64 decode failed")
		return nil, err
	}
	return d, nil

}

func tempFileName(prefix, suffix string) string {
	randBytes := make([]byte, 32)
	rand.Read(randBytes)
	return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)
}

func isValidTavanirRawQuery(q string) bool {
	re, _ := regexp.Compile("^amount = (\\d+)$")
	if re.MatchString(q) {
		return true
	}

	re, _ = regexp.Compile("^amount BETWEEN (\\d+) AND (\\d+)$")
	if re.MatchString(q) {
		return true
	}
	log.Println("not valid advanved query")
	return false
}
