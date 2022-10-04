package tavanir

import (
	"time"

	"github.com/spf13/cast"

	"github.com/r3labs/diff"
)

type APIError struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type Message struct {
	Id        int64  `gorm:"primary_key"`
	MessageID string `json:"messageId" gorm:"type:varchar(64); index:idx_tvn_message_id"`
	CaseID    string `json:"caseId" gorm:"type:varchar(64); index:idx_tvn_message_case_id"`
	RefID     string `json:"refId"`
	Message   string `json:"message"`
	DateTime  string `json:"dateTime"`
	Seen      string `json:"seen"`
}

type MessageSlice []Message

func (d *Message) TableName() string {
	return "tavanir_message"
}

type Document struct {
	Id             int64  `gorm:"primary_key"`
	CaseId         int64  `json:"case_id"`
	DocumentTypeID string `json:"documentTypeId"`
	FileName       string `json:"fileName"`
	FileType       string `json:"fileType"`
	Content        string `gorm:"type:longtext" json:"content"`
}
type DocumentSlice []Document

func (d *Document) TableName() string {
	return "tavanir_documents"
}

type BankInfo struct {
	AccountOwnerName   string `json:"accountOwnerName"`
	AccountShebaNumber string `json:"accountShebaNumber"`
	AccountBankName    string `json:"accountBankName"`
	Description        string `json:"description"`
}

const (
	CaseStatusNew string = "NEW"
)

type CaseElastic struct {
	Case
	Amount int64 `json:"amount"`
}

func (c Case) ToElastic() CaseElastic {
	ce := CaseElastic{
		Case:   c,
		Amount: cast.ToInt64(c.Amount),
	}
	return ce
}

type Case struct {
	Id                 int64  `gorm:"primary_key"`
	TavanirID          string `json:"id" gorm:"column:tavanir_id;type:varchar(64);uniqueIndex"`
	TrackingID         string `json:"trackingId"`
	BillID             string `json:"billId"`
	CompanyID          string `json:"companyId"`
	UserType           string `json:"userType"`
	NationalID         string `json:"nationalId"`
	UserName           string `json:"userName"`
	State              string `json:"state"`
	StateName          string `json:"stateName"`
	City               string `json:"city"`
	CityName           string `json:"cityName"`
	Address            string `json:"address"`
	PostalCode         string `json:"postalCode"`
	LocationTypeID     string `json:"locationTypeId"`
	LocationStatusID   string `json:"locationStatusId"`
	Amount             string `json:"amount"`
	CompensationTypeID string `json:"compensationTypeId"`
	Descr              string `json:"descr"`
	EventDate          string `json:"eventDate"`
	EventTime          string `json:"eventTime"`
	MobileNo           string `json:"mobileNo"`

	Documents []Document `json:"documents"`

	MissingDocuments     string    `json:"missing_documents"`
	NotCoveredReason     int       `json:"not_covered_reason"`
	AcceptedAmount       int64     `json:"accepted_amount"`
	ExpertAcceptedAmount int64     `json:"expert_accepted_amount"`
	ExpertDescription    string    `gorm:"type:text" json:"expert_description"`
	ExpertUpdatedAt      time.Time `json:"expert_updated_at"`
	LackDataDescription  string    `gorm:"type:text" json:"lack_data_description"`
	ExpertStatus         string    `gorm:"default:'DEFAULT'" json:"expert_status"`
	Sheba                string    `json:"sheba"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`

	// additional fields
	Status      string `gorm:"default:'NEW';index:status" json:"status"`
	IsDuplicate bool   `gorm:"-"`
}

func (c *Case) TableName() string {
	return "tavanir_damage"
}

type CaseSlice []Case

type CaseChangelog struct {
	ID            uint           `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	CaseID        uint           `gorm:"index:case_id" json:"case_id"`
	ChangeLogs    string         `gorm:"type:text" json:"-"`
	ChangelogData diff.Changelog `gorm:"-" json:"changelog"`
	UserID        uint           `gorm:"index:user_id" json:"user_id"`
}

func (c *CaseChangelog) TableName() string {
	return "tavanir_case_changelog"
}

type CaseChangelogSlice []CaseChangelog

type StatusUpdateQueue struct {
	ID        int64  `gorm:"primary_key"`
	CaseID    uint64 `gorm:"index:case_id"`
	NewStatus string
	RefID     uint64
	Note      string
	Success   bool `gorm:"default=false"`
	IsDone    bool `gorm:"default=false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type StatusUpdateQueueSlice []StatusUpdateQueue

func (c *StatusUpdateQueue) TableName() string {
	return "tavanir_status_update_queue"
}

type AddCaseResult struct {
	CaseID     string `json:"caseId"`
	TrackingID string `json:"trackingId"`
}

type CaseAdd struct {
	CaseID             uint64     `json:"caseId"`
	BillID             uint64     `json:"billId"`
	UserType           uint8      `json:"userType"`
	LocationTypeID     uint8      `json:"locationTypeId"`
	NationalID         string     `json:"nationalCode"`
	UserName           string     `json:"userName"`
	MobileNo           string     `json:"mobileNo"`
	StateID            uint       `json:"stateId"`
	CityID             uint       `json:"cityId"`
	PostalCode         uint64     `json:"postalCode"`
	Address            string     `json:"address"`
	EventDate          string     `json:"eventDateTime"`
	Amount             uint64     `json:"amount"`
	CompensationTypeID uint64     `json:"compensationTypeId"`
	Documents          []Document `json:"documents"`
}

// --------------------------------------
///////////////
// Responses //
///////////////
type getCaseResponse struct {
	Status   int      `json:"status"`
	CaseData Case     `json:"result"`
	Error    APIError `json:"error"`
}

type getNextMessageResponse struct {
	Status   int       `json:"status"`
	Messages []Message `json:"result"`
	Error    APIError  `json:"error"`
}

type getInfoResponse struct {
	Status int `json:"status"`
	Result struct {
		Description string `json:"description"`
	} `json:"result"`
	Error APIError `json:"error"`
}

type getCseStatusResponse struct {
	Status int `json:"status"`
	Result struct {
		Status string `json:"status"`
		Title  string `json:"statusTitle"`
	} `json:"result"`
	Error APIError `json:"error"`
}

type infoUpdateResponse struct {
	Status int `json:"status"`
	Result struct {
		CaseID string `json:"caseId"`
		RefID  string `json:"refId"`
	} `json:"result"`
	Error APIError `json:"error"`
}

type getUpdateBankInfoResponse struct {
	Status int      `json:"status"`
	Result BankInfo `json:"result"`
	Error  APIError `json:"error"`
}

type getDocumentUpdateResponse struct {
	Status int `json:"status"`
	Result struct {
		Description string     `json:"description"`
		Documents   []Document `json:"documents"`
	} `json:"result"`
	Error APIError `json:"error"`
}

type addCaseResponse struct {
	Status int           `json:"status"`
	Result AddCaseResult `json:"result"`
	Error  APIError      `json:"error"`
}

// --------------------------------------
//////////////
// Requests //
//////////////
type getNextMessageRequest struct {
	MessageID    uint64 `json:"messageId"`
	JustNotSeen  int8   `json:"justNotSeen"`
	RecordsCount uint16 `json:"recordsCount"`
}

type getCaseRequest struct {
	CaseID uint64 `json:"caseId"`
}

type getInfoRequest struct {
	CaseID uint64 `json:"caseId"`
	RefID  uint64 `json:"refId"`
}

type setDecriptionRequest struct {
	CaseID      uint64 `json:"caseId"`
	Description string `json:"description"`
}

type setDocumentRequiredRequest struct {
	CaseID         uint64 `json:"caseId"`
	DocumentTypeID []int  `json:"documentTypeId"`
	Description    string `json:"description"`
}

type setAmountEstimateRequest struct {
	CaseID      uint64 `json:"caseId"`
	Amount      uint64 `json:"amount"`
	Description string `json:"description"`
}

type setNotCoveredRequest struct {
	CaseID      uint64 `json:"caseId"`
	ReasonID    int    `json:"notCoverReasonId"`
	Description string `json:"description"`
}

type setCoveredRequest struct {
	CaseID      uint64 `json:"caseId"`
	Description string `json:"description"`
}

type addPaymentRequest struct {
	CaseID              uint64 `json:"caseId"`
	Description         string `json:"description"`
	PaymentAmount       uint64 `json:"paymentAmount"`
	PaymentRefID        string `json:"paymentRefId"`
	PaymentDate         string `json:"paymentDate"`
	TargetBankName      string `json:"targetBankName"`
	TargetAccountNumber string `json:"targetAccountNumber"`
}
