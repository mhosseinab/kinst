package export

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cast"
	"models"
)

func runSheetThree(resXlsx *excelize.File, sheetName string, x map[string]companyData, from, to string) {
	for in, r := range getSortedCompanies() {
		sheetThreeSetValues(r, resXlsx, sheetName, cast.ToString(4+(in*7)), x, from, to)

	}
}

func sheetThreeSetValues(cid string, resXlsx *excelize.File, sheetName string, row string, x map[string]companyData, from, to string) {

	title := fmt.Sprintf(`فرم 3- خسارات مشترکین خانگی (شهری و روستایی) از تاریخ %s لغایت %s`, queryTimeTostring(from), queryTimeTostring(to))

	resXlsx.SetCellValue(sheetName, "A1", title)
	com := x[cid]

	startRow := cast.ToInt(row)
	for i := 0; i < 6; i++ {
		curComDamage := com.iterate()
		cell := fmt.Sprintf("%s%d", "E", startRow)
		resXlsx.SetCellValue(sheetName, cell, curComDamage.count)

		cell = fmt.Sprintf("%s%d", "F", startRow)
		resXlsx.SetCellValue(sheetName, cell, curComDamage.sumDamageAmount)

		cell = fmt.Sprintf("%s%d", "G", startRow)
		resXlsx.SetCellValue(sheetName, cell, 0)

		cell = fmt.Sprintf("%s%d", "H", startRow)
		resXlsx.SetCellValue(sheetName, cell, 0)

		cell = fmt.Sprintf("%s%d", "I", startRow)
		resXlsx.SetCellValue(sheetName, cell, curComDamage.statusData[models.RequestStatusPayed].count)

		cell = fmt.Sprintf("%s%d", "J", startRow)
		resXlsx.SetCellValue(sheetName, cell, curComDamage.statusData[models.RequestStatusPayed].sumDamageAmount)

		cell = fmt.Sprintf("%s%d", "M", startRow)
		resXlsx.SetCellValue(sheetName, cell, curComDamage.statusData[models.RequestStatusReadyToPay].count)

		cell = fmt.Sprintf("%s%d", "N", startRow)
		resXlsx.SetCellValue(sheetName, cell, curComDamage.statusData[models.RequestStatusReadyToPay].sumDamageAmount)

		startRow = startRow + 1
	}

}
