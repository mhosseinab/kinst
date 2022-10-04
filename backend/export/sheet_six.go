package export

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/olivere/elastic/v7"
	"models"
)

func runSheetSix(resXlsx *excelize.File, sheetName string, qState, dateQuery *elastic.BoolQuery, from, to string) {
	client := getEsClient()

	q := elastic.NewBoolQuery().Must(
		elastic.NewTermQuery("status.keyword", models.RequestStatusClosed),
	).Must(elastic.NewExistsQuery("reference_code")).Must(qState).Must(dateQuery)

	res, _ := client.Search().
		Index(elasticIndexName).
		Query(q).
		Size(10000).
		Do(context.Background())

	for row, h := range res.Hits.Hits {
		curRow := row + 5
		r := models.Request{}
		json.Unmarshal(h.Source, &r)

		cell := fmt.Sprintf("%s%d", "A", curRow)
		resXlsx.SetCellValue(sheetName, cell, row+1)

		cell = fmt.Sprintf("%s%d", "B", curRow)
		resXlsx.SetCellValue(sheetName, cell, r.Firstname)

		cell = fmt.Sprintf("%s%d", "C", curRow)
		resXlsx.SetCellValue(sheetName, cell, r.Surname)

		cell = fmt.Sprintf("%s%d", "D", curRow)
		resXlsx.SetCellValue(sheetName, cell, r.BillIdentifier)

		cell = fmt.Sprintf("%s%d", "E", curRow)
		resXlsx.SetCellValue(sheetName, cell, r.NationalCode)

		cell = fmt.Sprintf("%s%d", "K", curRow)
		resXlsx.SetCellValue(sheetName, cell, r.GetStatusMessage())

		cell = fmt.Sprintf("%s%d", "L", curRow)
		resXlsx.SetCellValue(sheetName, cell, r.ExpertDescription)

	}

}
