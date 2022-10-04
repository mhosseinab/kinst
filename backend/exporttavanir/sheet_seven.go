package exporttavanir

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"models"
)

func runSheetSeven(resXlsx *excelize.File, sheetName string, x damageReport, from, to string) {
	house := x.DamageTypeHouse
	// biz := x.DamageTypeBiz
	// all := x.DamageTypeAll

	// curA := 'C'
	// curB := 'D'

	// for i := 4; i < 6+4; i++ {
	curDamage := house.iterate()
	log.Println(curDamage.count)

	cell := fmt.Sprintf("C%s", "4")
	resXlsx.SetCellValue(sheetName, cell, curDamage.count)

	cell = fmt.Sprintf("D%s", "4")
	resXlsx.SetCellValue(sheetName, cell, curDamage.sumDamageAmount)

	cell = fmt.Sprintf("E%s", "5")
	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusInProgress].count)

	cell = fmt.Sprintf("F%s", "5")
	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusInProgress].sumDamageAmount)
	// }

	// 	curA = nextChar(curA)
	// 	curB = nextChar(curB)

	// 	log.Println(cell)
	// 	cell = fmt.Sprintf("%c%s", curA, "7")
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusIncomplete].count)

	// 	log.Println(cell)
	// 	cell = fmt.Sprintf("%c%s", curB, "7")
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusIncomplete].sumDamageAmount)

	// 	curA = nextChar(curA)
	// 	curB = nextChar(curB)

	// 	log.Println(cell)
	// 	cell = fmt.Sprintf("%c%s", curA, "8")
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].count)

	// 	log.Println(cell)
	// 	cell = fmt.Sprintf("%c%s", curB, "8")
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].sumDamageAmount)

	// 	curA = nextChar(curA)
	// 	curB = nextChar(curB)

	// 	log.Println(cell)
	// 	cell = fmt.Sprintf("%c%s", curA, "9")
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].count)

	// 	log.Println(cell)
	// 	cell = fmt.Sprintf("%c%s", curB, "9")
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].sumDamageAmount)

	// 	curA = nextChar(curA)
	// 	curB = nextChar(curB)

	// 	log.Println(cell)
	// 	cell = fmt.Sprintf("%c%s", curA, "10")
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusClosed].count)

	// 	log.Println(cell)
	// 	cell = fmt.Sprintf("%c%s", curB, "10")
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusClosed].sumDamageAmount)
	// 	log.Println(cell)

	// 	// curA = nextChar(curA)
	// 	// curB = nextChar(curB)

	// 	log.Printf("%c%c", curA, curB)

	// }

	// for i := 17; i < 6+17; i++ {
	// 	curDamage := biz.iterate()

	// 	cell := fmt.Sprintf("%s", "C"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.count)

	// 	cell = fmt.Sprintf("%s", "D"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.sumDamageAmount)

	// 	cell = fmt.Sprintf("%s", "G"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].count)

	// 	cell = fmt.Sprintf("%s", "H"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].sumDamageAmount)

	// 	cell = fmt.Sprintf("%s", "K"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].count)

	// 	cell = fmt.Sprintf("%s", "L"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].sumDamageAmount)
	// }

	// for i := 30; i < 6+30; i++ {
	// 	curDamage := all.iterate()

	// 	cell := fmt.Sprintf("%s", "C"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.count)

	// 	cell = fmt.Sprintf("%s", "D"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.sumDamageAmount)

	// 	cell = fmt.Sprintf("%s", "G"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].count)

	// 	cell = fmt.Sprintf("%s", "H"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusPayed].sumDamageAmount)

	// 	cell = fmt.Sprintf("%s", "K"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].count)

	// 	cell = fmt.Sprintf("%s", "L"+cast.ToString(i))
	// 	resXlsx.SetCellValue(sheetName, cell, curDamage.statusData[models.RequestStatusReadyToPay].sumDamageAmount)
	// }

}
