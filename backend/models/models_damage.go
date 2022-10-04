package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Request status
const (
	RequestStatusNotCompleted        = "NOT_COMPLETED"       // form not completed
	RequestStatusRejected            = "REJECTED"            // rejected by tavanir
	RequestStatusAccepted            = "TAVANIR_ACCEPTED"    // accepted by tavanir
	RequestStatusInProgress          = "IN_PROGRESS"         // jaari - option 1
	RequestStatusClosed              = "CLOSED"              // makhtoomeh - option 2
	RequestStatusSuspended           = "SUSPENDED"           // moavvagh - option 3
	RequestStatusIncomplete          = "INCOMPLETE"          // darkhaste naghes - option 4
	RequestStatusIncompleteChange    = "INCOMPLETE_CHANGE"   // darkhaste naghes change - option 4
	RequestStatusReadyToPay          = "READY_TO_PAY"        // amadeye pardakht - option 5
	RequestStatusPayed               = "PAYED"               // pardakht shode - option 6
	RequestStatusCanceledByUser      = "CANCELED_BY_USER"    // enserafe zi nafe option 7
	RequestStatusInactive            = "INACTIVE"            // in active
	RequestStatusAmountRejected      = "AMOUNT_REJECTED"     // amount rejected
	RequestStatusNeedToVisitInPerson = "Need_Visit_InPerson" // Need to visit in person
	RequestCompliant                 = "REQUEST_COMPLIANT"
)

// expert status
const (
	RequestExpertStatusDefault  = "DEFAULT"
	RequestExpertStatusAccepted = "ACCEPTED"
	RequestExpertStatusRejected = "REJECTED"
)

// Request damage types
const (
	RequestDamageTypeDeath      = "death_damage"
	RequestDamageTypeLack       = "lack_damage"
	RequestDamageTypeMedical    = "medical_damage"
	RequestDamageTypeExplosion  = "explosion_damage"
	RequestDamageTypeInstrument = "instrument_damage"
	RequestDamageTypeFiring     = "firing_damage"
)

// Request struct
type Request struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Token            string    `gorm:"index:token;" json:"token"`
	BillIdentifier   string    `gorm:"type:varchar(50);index:bill_indetifier" json:"bill_identifier"` // shenaseye ghabz
	JointsIdentifier string    `gorm:"type:varchar(50)" json:"joints_identifier"`                     // shenaseye ghabze moshaat
	CompanyID        string    `gom:"type:varchar(3);index:company_id" json:"company_id"`
	LocationUsage    string    `gom:"type:varchar(3)" json:"location_usage"`
	Firstname        string    `json:"firstname"`
	Surname          string    `json:"surname"`
	NationalCode     string    `json:"national_code"`
	MobileNumber     string    `json:"mobile_number"`
	Province         string    `gorm:"index:province" json:"province"`
	City             string    `gorm:"index:city" json:"city"`
	Address          string    `json:"address"`
	PostalAddress    string    `json:"postal_address"`
	ReferenceCode    string    `gorm:"index:reference_code" json:"reference_code"`
	CasualityDate    time.Time `json:"casuality_date"`
	CasualityTime    string    `json:"casuality_time"`
	Description      string    `gorm:"type:text" json:"description"`
	Status           string    `gorm:"default:'NOT_COMPLETED';index:status" json:"status"`
	LastBillPhoto    string    `json:"last_bill_photo"`
	LocationType     int       `json:"location_type"`
	IDCardPhoto      string    `json:"id_card_photo"`
	OtherPhoto       string    `json:"other_photo"`

	DamageType           string    `gorm:"type:varchar(50);index:damage_type" json:"damage_type"`
	AcceptedAmount       int64     `json:"accepted_amount"`
	ExpertAcceptedAmount int64     `json:"expert_accepted_amount"`
	ExpertDescription    string    `gorm:"type:text" json:"expert_description"`
	ExpertUpdatedAt      time.Time `json:"expert_updated_at"`
	LackDataDescription  string    `gorm:"type:text" json:"lack_data_description"`
	ExpertStatus         string    `gorm:"default:'DEFAULT'" json:"expert_status"`
	SubscriberType       int       `gorm:"default:1" json:"subscriber_type"`
	EconomicCode         string    `json:"economic_code"`
	Sheba                string    `json:"sheba"`
	// DamageAmount                    int64  `gorm:"default:0" json:"damage_amount"`
	SumDamageAmount                 int64  `json:"sum_damage_amount"`
	FiringDamageAmount              int64  `json:"firing_damage_amount"`
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
	InstrumentDamageAmount          int64  `json:"instrument_damage_amount"`
	InstrumentInvoicePhoto          string `json:"instrument_invoice_photo"`
	InstrumentInvoice2Photo         string `json:"instrument_invoice_2_photo"`
	InstrumentInvoice3Photo         string `json:"instrument_invoice_3_photo"`
	InstrumentInvoice4Photo         string `json:"instrument_invoice_4_photo"`
	InstrumentInvoice5Photo         string `json:"instrument_invoice_5_photo"`
	InstrumentInvoice6Photo         string `json:"instrument_invoice_6_photo"`
	InstrumentInvoice7Photo         string `json:"instrument_invoice_7_photo"`
	InstrumentInvoice8Photo         string `json:"instrument_invoice_8_photo"`
	InstrumentReportPhoto           string `json:"instrument_report_photo"`
	ExplosionDamageAmount           int64  `json:"explosion_damage_amount"`
	ExplosionFirestationReportPhoto string `json:"explosion_firestation_report_photo"`
	ExplostionDamagedItemsPhoto     string `json:"explostion_damaged_items_photo"`
	ExplosionInvoicePhoto           string `json:"explosion_invoice_photo"`
	ExplosionPlace1Photo            string `json:"explosion_place_1_photo"`
	ExplosionPlace2Photo            string `json:"explosion_place_2_photo"`
	MedicalDamageAmount             int64  `json:"medical_damage_amount"`
	MedicalHospitalInvoicePhoto     string `json:"medical_hospital_invoice_photo"`
	MedicalHospitalDocumentPhoto    string `json:"medical_hospital_document_photo"`
	MedicalReportPhoto              string `json:"medical_report_photo"`
	LackDamageAmount                int64  `json:"lack_damage_amount"`
	LackReportPhoto                 string `json:"lack_report_photo"`
	LackRadiologyPhoto              string `json:"lack_radiology_photo"`
	LackFirstReferencePhoto         string `json:"lack_first_reference_photo"`
	LackIDCardPhoto                 string `json:"lack_id_card_photo"`
	LackWitnessPhoto                string `json:"lack_witness_photo"`
	LackInvoicePhoto                string `json:"lack_invoice_photo"`
	DeathDamageAmount               int64  `json:"death_damage_amount"`
	DeathJudgeVotePhoto             string `json:"death_judge_vote_photo"`
	DeathWitnessPhoto               string `json:"death_witness_photo"`
	DeathIDCardPhoto                string `json:"death_id_card_photo"`
	DeathToxicologyReportPhoto      string `json:"death_toxicology_report_photo"`
	DeathCorpseExaminationPhoto     string `json:"death_corpse_examination_photo"`
	DeathProbatePhoto               string `json:"death_probate_photo"`
}

func (r *Request) GetDamageTypeMessage() string {
	switch r.DamageType {
	case RequestDamageTypeDeath:
		return "فوت"
	case RequestDamageTypeExplosion:
		return "انفجار"
	case RequestDamageTypeFiring:
		return "آتش سوزی"
	case RequestDamageTypeInstrument:
		return "تجهیزات"
	case RequestDamageTypeLack:
		return "نقص عضو"
	case RequestDamageTypeMedical:
		return "پزشکی"

	}

	return r.DamageType
}

func (r *Request) GetStatusMessage() string {
	switch r.Status {
	case RequestStatusInProgress:
		return "جاری"
	case RequestStatusClosed:
		return "مختومه"
	case RequestStatusSuspended:
		return "معوق"
	case RequestStatusIncomplete:
		return "درخواست ناقص"
	case RequestStatusIncompleteChange:
		return "تکمیل درخواست ناقص"
	case RequestStatusReadyToPay:
		return "آماده پرداخت"
	case RequestStatusPayed:
		return "پرداخت شده"
	case RequestStatusCanceledByUser:
		return "انصراف ذی نفع"
	}
	return ""
}

// RequestSlice slice of request
type RequestSlice []Request

// BeforeCreate hook
func (r *Request) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Status", RequestStatusNotCompleted)
	return nil
}

// BeforeSave hook
func (r *Request) BeforeSave() (err error) {
	r.SumDamageAmount =
		r.FiringDamageAmount +
			r.InstrumentDamageAmount +
			r.ExplosionDamageAmount +
			r.MedicalDamageAmount +
			r.LackDamageAmount +
			r.DeathDamageAmount

	if r.ExpertStatus == RequestExpertStatusAccepted {
		r.LackDataDescription = ""
	}

	return
}

// Damage struct
type Damage struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
