package exporttavanir

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cast"
	"models"
)

func runSheetTwo(resXlsx *excelize.File, sheetName string, x map[string]companyData, from, to string) {
	for in, r := range getSortedCompanies() {
		sheetTwoSetValues(r, resXlsx, sheetName, cast.ToString(4+in), x, from, to)
	}
}

func sheetTwoSetValues(cid string, resXlsx *excelize.File, sheetName string, row string, x map[string]companyData, from, to string) {
	title := fmt.Sprintf(`فرم 2- گزارش تجمیعی عملکرد بیمه مشترکین تجاری (سایر مصارف) از تاریخ %s لغایت %s`, queryTimeTostring(from), queryTimeTostring(to))

	resXlsx.SetCellValue(sheetName, "A1", title)
	com := x[cid]
	// spew.Dump(com)

	curCount := 'D'
	curAmount := 'K'

	for i := 0; i < 6; i++ {
		curComDamage := com.iterate()
		resXlsx.SetCellValue(sheetName, string(curCount)+row, curComDamage.count)
		resXlsx.SetCellValue(sheetName, string(curAmount)+row, curComDamage.sumDamageAmount)
		curCount = nextChar(curCount)
		curAmount = nextChar(curAmount)
	}

	resXlsx.SetCellValue(sheetName, "R"+row, com.statusData[models.RequestStatusPayed].count)
	resXlsx.SetCellValue(sheetName, "S"+row, com.statusData[models.RequestStatusPayed].sumDamageAmount)

	resXlsx.SetCellValue(sheetName, "T"+row, com.statusData[models.RequestStatusReadyToPay].count)
	resXlsx.SetCellValue(sheetName, "U"+row, com.statusData[models.RequestStatusReadyToPay].sumDamageAmount)

	resXlsx.SetCellValue(sheetName, "V"+row, com.statusData[models.RequestStatusClosed].count)
	resXlsx.SetCellValue(sheetName, "W"+row, com.statusData[models.RequestStatusClosed].sumDamageAmount)

	resXlsx.SetCellValue(sheetName, "X"+row, com.statusData[models.RequestStatusInProgress].count)
	resXlsx.SetCellValue(sheetName, "Y"+row, com.statusData[models.RequestStatusInProgress].sumDamageAmount)

}
