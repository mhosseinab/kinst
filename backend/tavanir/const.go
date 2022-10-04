package tavanir

import (
	"strings"

	"models"
)

const (
	baseURL string = "https://bime.tavanir.org.ir/api/Insurance/v1/Rest/"
)

// --------------------------------------
/////////////
// Methods //
/////////////
const (
	MethodGetNextMessages           string = "getNextMessages"
	MethodGetCase                   string = "getCase"
	MethodGetAssignedInfo           string = "getAssignedInfo"
	MethodGetDocumentUpdate         string = "getDocumentUpdate"
	MethodSetPhysicalRequired       string = "setPhysicalRequired"
	MethodSetNotCovered             string = "setNotCovered"
	MethodSetCovered                string = "setCovered"
	MethodGetTavanirCoverCompliant  string = "getTavanirCoverCompliant"
	MethodSetDocumentRequired       string = "setDocumentRequired"
	MethodSetAmountEstimate         string = "setAmountEstimate"
	MethodGetUpdateBankInfo         string = "getUpdateBankInfo"
	MethodGetAmountCompliant        string = "getAmountCompliant"
	MethodGetTavanirAmountCompliant string = "getTavanirAmountCompliant"
	MethodAddPayment                string = "addPayment"
	MethodCheckout                  string = "checkout"
	MethodCheckoutWithoutPayment    string = "checkoutWithoutPayment"
	MethodPaymentInfoUpdateRequest  string = "paymentInfoUpdateRequest"
	MethodGetCheckoutCompliant      string = "getCheckoutCompliant"
	MethodGetCaseStatus             string = "getCaseStatus"
	FetchFailedMessage              string = "FetchFailedMessage"
)

// --------------------------------------
/////////////

const (
	StatusMessageAssigned               string = "assigned"
	StatusMessageTavanirCoverCompliant  string = "tavanirCoverCompliant"
	StatusMessageDocumentUpdate         string = "documentUpdate"
	StatusMessageAmountCompliant        string = "amountCompliant"
	StatusMessageTavanirAmountCompliant string = "tavanirAmountCompliant"
	StatusMessageUpdateBankInfo         string = "updateBankInfo"
	StatusMessageCheckoutCompliant      string = "checkoutCompliant"

	justNotSeenTypeAll    int8 = 0
	justNotSeenTypeUnread int8 = 1

	recordsCountType100   uint16 = 100
	recordsCountType1000  uint16 = 1000
	recordsCountType10000 uint16 = 10000
)

var tavanirDamageTypes = map[string]string{
	"1": "فوت",
	"2": "نقص عضو",
	"3": "هزينه پزشکي",
	"4": "انفجار",
	"5": "لوازم تجهيزات",
	"6": "آتش سوزي",
}

var tavanirStatusText = map[string]string{
	//"p0" : //شروع
	//"p1" : //در انتظار بررسی اولیه توسط شرکت برق
	//"p2" : //
	//"p3" : //در انتظار ثبت نتیجه بررسی شرکت برق
	"p4": models.RequestStatusInProgress, //در انتظار بررسی در شرکت بیمه
	"p5": models.RequestStatusIncomplete, //در انتظار تصحیح نقص مدارک
	"p6": models.RequestStatusInProgress, //در انتظار تعیین خسارت
	"p7": models.RequestStatusInProgress, //در انتظار مشاهده خسارت
	"p8": models.RequestStatusReadyToPay, //در انتظار ثبت اطلاعات بانکی
	//"p9" : //در انتظار بررسی اعتراض توسط دستگاه ناظر
	//"p10" : //در انتظار تایید/اعتراض
	"p11": models.RequestStatusReadyToPay, //در انتظار بررسی و واریز توسط بیمه
	"p12": models.RequestStatusPayed,      //در انتظار تایید واریز توسط مشترک
	//"p13" : //در انتظار بررسی اعتراض به پرداخت
	//"p14" : //در انتظار بررسی پرونده در دستگاه ناظر
	"p99": models.RequestStatusPayed, //پایان
}

var TavanirStatus = map[string]string{
	"NEW":               "جدید",
	"CLOSED":            "مختومه",
	"IN_PROGRESS":       "جاری",
	"INCOMPLETE":        "درخواست ناقص",
	"INCOMPLETE_CHANGE": "تکمیل درخواست ناقص",
	"PAYED":             "پرداخت شده",
	"READY_TO_PAY":      "آماده پرداخت",
	"SUSPENDED":         "معوق",
}

var TavanirDamageTypes = map[string]string{
	"1": "فوت",
	"2": "نقص عضو",
	"3": "هزينه پزشکي",
	"4": "انفجار",
	"5": "لوازم تجهيزات",
	"6": "آتش سوزي",
}

func FixArabicCharacters(str string) string {
	ye1 := "ي"
	ye2 := "ی"
	ke1 := "ك"
	ke2 := "ک"
	return strings.ReplaceAll(strings.ReplaceAll(str, ye1, ye2), ke1, ke2)
}
