package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"connections"
	"es"
	"models"
	"tools"
)

func step6(c *gin.Context) {
	if v := tools.GetEnv("DISABLE_STEPS", "0"); v == "1" {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}
	type fstep6 struct {
		// DamageAmount                    int64  `json:"damage_amount"`
		FiringDamageAmount              string `json:"firing_damage_amount"`
		FiringPlace1Photo               string `json:"firing_place_1_photo"`
		FiringPlace2Photo               string `json:"firing_place_2_photo"`
		FiringStationReportPhoto        string `json:"firing_station_report_photo"`
		FiringPoliceReportPhoto         string `json:"firing_police_report_photo"`
		FiringCourtReportPhoto          string `json:"firing_court_report_photo"`
		FiringInvoice1Photo             string `json:"firing_invoice_1_photo"`
		FiringInvoice2Photo             string `json:"firing_invoice_2_photo"`
		FiringInvoice3Photo             string `json:"firing_invoice_3_photo"`
		FiringInvoice4Photo             string `json:"firing_invoice_4_photo"`
		FiringInvoice5Photo             string `json:"firing_invoice_5_photo"`
		InstrumentDamageAmount          string `json:"instrument_damage_amount"`
		InstrumentInvoicePhoto          string `json:"instrument_invoice_photo"`
		InstrumentInvoice2Photo         string `json:"instrument_invoice_2_photo"`
		InstrumentInvoice3Photo         string `json:"instrument_invoice_3_photo"`
		InstrumentInvoice4Photo         string `json:"instrument_invoice_4_photo"`
		InstrumentInvoice5Photo         string `json:"instrument_invoice_5_photo"`
		InstrumentReportPhoto           string `json:"instrument_report_photo"`
		ExplosionDamageAmount           string `json:"explosion_damage_amount"`
		ExplosionFirestationReportPhoto string `json:"explosion_firestation_report_photo"`
		ExplostionDamagedItemsPhoto     string `json:"explostion_damaged_items_photo"`
		ExplosionInvoicePhoto           string `json:"explosion_invoice_photo"`
		ExplosionPlace1Photo            string `json:"explosion_place_1_photo"`
		ExplosionPlace2Photo            string `json:"explosion_place_2_photo"`
		MedicalDamageAmount             string `json:"medical_damage_amount"`
		MedicalHospitalInvoicePhoto     string `json:"medical_hospital_invoice_photo"`
		MedicalHospitalDocumentPhoto    string `json:"medical_hospital_document_photo"`
		MedicalReportPhoto              string `json:"medical_report_photo"`
		LackDamageAmount                string `json:"lack_damage_amount"`
		LackReportPhoto                 string `json:"lack_report_photo"`
		LackRadiologyPhoto              string `json:"lack_radiology_photo"`
		LackFirstReferencePhoto         string `json:"lack_first_reference_photo"`
		LackIDCardPhoto                 string `json:"lack_id_card_photo"`
		LackWitnessPhoto                string `json:"lack_witness_photo"`
		LackInvoicePhoto                string `json:"lack_invoice_photo"`
		DeathDamageAmount               string `json:"death_damage_amount"`
		DeathJudgeVotePhoto             string `json:"death_judge_vote_photo"`
		DeathWitnessPhoto               string `json:"death_witness_photo"`
		DeathIDCardPhoto                string `json:"death_id_card_photo"`
		DeathToxicologyReportPhoto      string `json:"death_toxicology_report_photo"`
		DeathCorpseExaminationPhoto     string `json:"death_corpse_examination_photo"`
		DeathProbatePhoto               string `json:"death_probate_photo"`
		DamageType                      string `json:"damage_type" binding:"required"`
	}
	var f fstep6
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": formDataValidationError})
		return
	}

	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("token not set").Error()})
		return
	}

	var m models.Request
	if err := db.First(&m, "token=?", reqToken).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	m.DamageType = f.DamageType
	// m.DamageAmount = f.DamageAmount
	m.FiringDamageAmount = cast.ToInt64(f.FiringDamageAmount)
	m.FiringPlace1Photo = f.FiringPlace1Photo
	m.FiringPlace2Photo = f.FiringPlace2Photo
	m.FiringStationReportPhoto = f.FiringStationReportPhoto
	m.FiringPoliceReportPhoto = f.FiringPoliceReportPhoto
	m.FiringCourtReportPhoto = f.FiringCourtReportPhoto
	m.FiringInvoice1Photo = f.FiringInvoice1Photo
	m.FiringInvoice2Photo = f.FiringInvoice2Photo
	m.FiringInvoice3Photo = f.FiringInvoice3Photo
	m.FiringInvoice4Photo = f.FiringInvoice4Photo
	m.FiringInvoice5Photo = f.FiringInvoice5Photo
	m.InstrumentDamageAmount = cast.ToInt64(f.InstrumentDamageAmount)
	m.InstrumentInvoicePhoto = f.InstrumentInvoicePhoto
	m.InstrumentInvoice2Photo = f.InstrumentInvoice2Photo
	m.InstrumentInvoice3Photo = f.InstrumentInvoice3Photo
	m.InstrumentInvoice4Photo = f.InstrumentInvoice4Photo
	m.InstrumentInvoice5Photo = f.InstrumentInvoice5Photo
	m.InstrumentReportPhoto = f.InstrumentReportPhoto
	m.ExplosionDamageAmount = cast.ToInt64(f.ExplosionDamageAmount)
	m.ExplosionFirestationReportPhoto = f.ExplosionFirestationReportPhoto
	m.ExplostionDamagedItemsPhoto = f.ExplostionDamagedItemsPhoto
	m.ExplosionInvoicePhoto = f.ExplosionInvoicePhoto
	m.ExplosionPlace1Photo = f.ExplosionPlace1Photo
	m.ExplosionPlace2Photo = f.ExplosionPlace2Photo
	m.MedicalDamageAmount = cast.ToInt64(f.MedicalDamageAmount)
	m.MedicalHospitalInvoicePhoto = f.MedicalHospitalInvoicePhoto
	m.MedicalHospitalDocumentPhoto = f.MedicalHospitalDocumentPhoto
	m.MedicalReportPhoto = f.MedicalReportPhoto
	m.LackDamageAmount = cast.ToInt64(f.LackDamageAmount)
	m.LackReportPhoto = f.LackReportPhoto
	m.LackRadiologyPhoto = f.LackRadiologyPhoto
	m.LackFirstReferencePhoto = f.LackFirstReferencePhoto
	m.LackIDCardPhoto = f.LackIDCardPhoto
	m.LackWitnessPhoto = f.LackWitnessPhoto
	m.LackInvoicePhoto = f.LackInvoicePhoto
	m.DeathDamageAmount = cast.ToInt64(f.DeathDamageAmount)
	m.DeathJudgeVotePhoto = f.DeathJudgeVotePhoto
	m.DeathWitnessPhoto = f.DeathWitnessPhoto
	m.DeathIDCardPhoto = f.DeathIDCardPhoto
	m.DeathToxicologyReportPhoto = f.DeathToxicologyReportPhoto
	m.DeathCorpseExaminationPhoto = f.DeathCorpseExaminationPhoto
	m.DeathProbatePhoto = f.DeathProbatePhoto

	sumAmount := m.FiringDamageAmount +
		m.InstrumentDamageAmount +
		m.ExplosionDamageAmount +
		m.MedicalDamageAmount +
		m.LackDamageAmount +
		m.DeathDamageAmount

	if sumAmount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": "شما مبلغی وارد نکردید."})
		return
	}

	m.Status = models.RequestStatusInProgress
	m.ReferenceCode = RandStringRunes(8)
	db.Save(&m)

	sendSmsTracking(m.ReferenceCode, m.MobileNumber)

	if client, err := connections.GetElasticsearch(); err == nil {
		es.StoreRequestItem(client, m)
	}

	c.JSON(200, gin.H{"status": 200, "reference_code": m.ReferenceCode})
}
