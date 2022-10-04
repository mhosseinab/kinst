package export

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cast"
	"models"
)

func runSheetFour(resXlsx *excelize.File, sheetName string, x map[string]companyData, from, to string) {
	for in, r := range getSortedCompanies() {
		sheetFourSetValues(r, resXlsx, sheetName, cast.ToString(4+(in*6)), x, from, to)

	}
}

func sheetFourSetValues(cid string, resXlsx *excelize.File, sheetName string, row string, x map[string]companyData, from, to string) {

	title := fmt.Sprintf(`فرم 4- خسارات مشترکین تجاری (سایر مصارف) شهری و روستایی  از تاریخ %s لغایت %s`, queryTimeTostring(from), queryTimeTostring(to))

	log.Println("row:", row)
	resXlsx.SetCellValue(sheetName, "A1", title)
	com := x[cid]

	startRow := cast.ToInt(row)
	for i := 0; i < 5; i++ {
		curComDamage := com.iterateBiz()
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

		cell = fmt.Sprintf("%s%d", "K", startRow)
		resXlsx.SetCellValue(sheetName, cell, curComDamage.statusData[models.RequestStatusReadyToPay].count)

		cell = fmt.Sprintf("%s%d", "L", startRow)
		resXlsx.SetCellValue(sheetName, cell, curComDamage.statusData[models.RequestStatusReadyToPay].sumDamageAmount)

		startRow = startRow + 1
	}

}
