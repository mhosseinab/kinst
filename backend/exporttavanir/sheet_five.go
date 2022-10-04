package exporttavanir

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cast"
	"models"
)

func runSheetFive(resXlsx *excelize.File, sheetName string, x damageReport, from, to string) {
	house := x.DamageTypeHouse
	biz := x.DamageTypeBiz
	all := x.DamageTypeAll

	for i := 4; i < 6+4; i++ {
		curDamage := house.iterate()

		cell := fmt.Sprintf("%s", "C"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.count)

		cell = fmt.Sprintf("%s", "D"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.sumDamageAmount)

		cell = fmt.Sprintf("%s", "G"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].count)

		cell = fmt.Sprintf("%s", "H"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].sumDamageAmount)

		cell = fmt.Sprintf("%s", "K"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].count)

		cell = fmt.Sprintf("%s", "L"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].sumDamageAmount)
	}

	for i := 15; i < 6+15; i++ {
		curDamage := biz.iterate()

		cell := fmt.Sprintf("%s", "C"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.count)

		cell = fmt.Sprintf("%s", "D"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.sumDamageAmount)

		cell = fmt.Sprintf("%s", "G"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].count)

		cell = fmt.Sprintf("%s", "H"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].sumDamageAmount)

		cell = fmt.Sprintf("%s", "K"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].count)

		cell = fmt.Sprintf("%s", "L"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].sumDamageAmount)
	}

	for i := 27; i < 6+27; i++ {
		curDamage := all.iterate()

		cell := fmt.Sprintf("%s", "C"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.count)

		cell = fmt.Sprintf("%s", "D"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.sumDamageAmount)

		cell = fmt.Sprintf("%s", "G"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].count)

		cell = fmt.Sprintf("%s", "H"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].sumDamageAmount)

		cell = fmt.Sprintf("%s", "K"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].count)

		cell = fmt.Sprintf("%s", "L"+cast.ToString(i))
		resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].sumDamageAmount)
	}

}
