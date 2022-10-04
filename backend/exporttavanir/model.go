package exporttavanir

import "models"

type statusData struct {
	name            string
	count           int64
	sumDamageAmount int64
}

type damageData struct {
	count           int64
	sumDamageAmount int64
	statusData      map[string]statusData
}

type companyData struct {
	instrumentDamage damageData
	firingDamage     damageData
	explosionDamage  damageData
	deathDamage      damageData
	lackDamage       damageData
	medicalDamage    damageData
	statusData       map[string]statusData
	cursor           int
	cursorBiz        int
}

type damagesType struct {
	instrumentDamage damageData
	firingDamage     damageData
	explosionDamage  damageData
	deathDamage      damageData
	lackDamage       damageData
	medicalDamage    damageData
	cursor           int
	cursorBiz        int
}

func (cd *damagesType) iterate() damageData {
	switch cd.cursor {
	case 0:
		cd.cursor++
		return cd.instrumentDamage
	case 1:
		cd.cursor++
		return cd.firingDamage
	case 2:
		cd.cursor++
		return cd.explosionDamage
	case 3:
		cd.cursor++
		return cd.medicalDamage
	case 4:
		cd.cursor++
		return cd.deathDamage
	case 5:
		cd.cursor = 0
		return cd.lackDamage
	}
	panic("unhandled company data")
}

func (cd *damagesType) iterateBiz() damageData {
	switch cd.cursorBiz {
	case 0:
		cd.cursorBiz++
		return cd.firingDamage
	case 1:
		cd.cursorBiz++
		return cd.explosionDamage
	case 2:
		cd.cursorBiz++
		return cd.medicalDamage
	case 3:
		cd.cursorBiz++
		return cd.deathDamage
	case 4:
		cd.cursorBiz = 0
		return cd.lackDamage
	}
	panic("unhandled company data")
}

func (cd *companyData) getPayedCount() (count int64) {
	for i := 0; i < 6; i++ {
		curComDamage := cd.iterate()
		count += curComDamage.statusData[models.RequestStatusPayed].count
	}
	return count
}

func (cd *companyData) iterate() damageData {
	switch cd.cursor {
	case 0:
		cd.cursor++
		return cd.instrumentDamage
	case 1:
		cd.cursor++
		return cd.firingDamage
	case 2:
		cd.cursor++
		return cd.explosionDamage
	case 3:
		cd.cursor++
		return cd.medicalDamage
	case 4:
		cd.cursor++
		return cd.deathDamage
	case 5:
		cd.cursor = 0
		return cd.lackDamage
	}
	panic("unhandled company data")
}

func (cd *companyData) iterateBiz() damageData {
	switch cd.cursorBiz {
	case 0:
		cd.cursorBiz++
		return cd.firingDamage
	case 1:
		cd.cursorBiz++
		return cd.explosionDamage
	case 2:
		cd.cursorBiz++
		return cd.medicalDamage
	case 3:
		cd.cursorBiz++
		return cd.deathDamage
	case 4:
		cd.cursorBiz = 0
		return cd.lackDamage
	}
	panic("unhandled company data")
}

// var companiesData map[string]companyData

// func init() {
// companiesData = make(map[string]companyData)
// }
