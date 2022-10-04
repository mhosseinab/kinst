package tavanir

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
)

// GetNextMessages godoc
// @Description دریافت پیام های سامانه توانیر به بیمه
// @Param MessageID 	uint64	پیام هایی که از شناسه این پیام جدیدتر هستند دریافت شوند. برای دریافت پیام ها از ابتدا مقدار 1 باشد.
// @Param JustNotSeen	justNotSeentype  فقط پیام های دیده نشده
// @Param RecordsCount	recordsCountype  تعداد پیام هایی که دریافت شوند
// @Returns []Message, error
func GetNextMessages(messageID uint64, justNotSeen int8, recordsCount uint16) ([]Message, error) {
	jsonData := getNextMessageRequest{
		MessageID:    messageID,
		JustNotSeen:  justNotSeen,
		RecordsCount: recordsCount,
	}
	jsonValue, _ := json.Marshal(jsonData)

	response, err := APIRequest(MethodGetNextMessages, jsonValue)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	//fmt.Println(string(response))
	var result getNextMessageResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}

	if result.Status != 1 {
		return nil, fmt.Errorf("API error: %s", result.Error.Message)
	}
	//fmt.Println(result.Status)
	return result.Messages, nil
}

func GetCaseString(id string) (Case, error) {
	return GetCase(cast.ToUint64(id))
}

// GetCase godoc
// @Description دریافت پرونده اعلم خسارت
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Returns Case, error
func GetCase(caseID uint64) (Case, error) {
	jsonData := getCaseRequest{
		CaseID: caseID,
	}
	jsonValue, _ := json.Marshal(jsonData)
	//log.Println("json:", string(jsonValue))
	response, err := APIRequest(MethodGetCase, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return Case{}, err
	}
	//fmt.Println(string(response))
	//err = ioutil.WriteFile("testData.json", response, 0644)
	//return nil, nil
	var result getCaseResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		////log.Fatal(err)
		return Case{}, err
	}

	if result.Status != 1 {
		return Case{}, fmt.Errorf("API error: %s", result.Error.Message)
	}

	//fmt.Println(result.CaseData.ID)

	return result.CaseData, nil
}

// SetPhysicalRequired godoc
// @Description نیاز به مراجعه حضوری مشترک و ارایه مدارک
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param desc	 	string	توضیحات درباره علت نیاز به مراجعه مشترک و مدارکی که باید همراه داشته باشد
// @Returns uint64, error
func SetPhysicalRequired(caseID uint64, desc string) (uint64, error) {
	return setDescription(MethodSetPhysicalRequired, caseID, desc)
}

// GetAssignedInfo godoc
// @Description دریافت توضیحات مربوط به پرونده ارسالی به بیمه
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param refID 	uint64	شماره ارجاع
// @Returns string, error
func GetAssignedInfo(caseID, refID uint64) (string, error) {
	return getDescription(caseID, refID, MethodGetAssignedInfo)
}

// SetDocumentRequired godoc
// @Description نیاز به الحاق مدارک
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param documentTypeIds 	[]int	آرایه ای از شناسه نوع سند
// @Param desc	 	string	توضحات درباره مدارکی که باید الحاق گردند
// @Returns string, error
func SetDocumentRequired(caseID uint64, documentTypeIds []int, desc string) (uint64, error) {
	jsonData := setDocumentRequiredRequest{
		CaseID:         caseID,
		DocumentTypeID: documentTypeIds,
		Description:    desc,
	}
	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(MethodSetDocumentRequired, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return 0, err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result infoUpdateResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		////log.Fatal(err)
		return 0, err
	}

	if result.Status != 1 {
		return 0, fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return cast.ToUint64(result.Result.RefID), nil
}

// GetDocumentUpdate godoc
// @Description دریافت مستندات آپلود شده توسط مشترک
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param refID 	uint64	شماره ارجاع
// @Returns []Document, error
func GetDocumentUpdate(caseID, refID uint64) ([]Document, error) {
	jsonData := getInfoRequest{
		CaseID: caseID,
		RefID:  refID,
	}
	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(MethodGetDocumentUpdate, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return nil, err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result getDocumentUpdateResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		////log.Fatal(err)
		return nil, err
	}

	if result.Status != 1 {
		return nil, fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return result.Result.Documents, nil
}

// SetNotCovered godoc
// @Description مشمول عدم دریافت خسارت
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param documentTypeIds 	int	دلیل عدم شمول
// @Param desc	 	string	توضیحات علت عدم شمول
// @Returns uint64, error
func SetNotCovered(caseID uint64, notCoverReasonID int, desc string) (uint64, error) {
	jsonData := setNotCoveredRequest{
		CaseID:      caseID,
		Description: desc,
	}

	jsonData.ReasonID = notCoverReasonID
	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(MethodSetNotCovered, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return 0, err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result infoUpdateResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		////log.Fatal(err)
		return 0, err
	}

	if result.Status != 1 {
		return 0, fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return cast.ToUint64(result.Result.RefID), nil
}

// GetTavanirCoverCompliant godoc
// @Description دریافت توضیحات توانیر در رابطه با اعتراض مشترک به عدم شمول خسارت.
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param refID 	uint64	شماره ارجاع
// @Returns string, error
func GetTavanirCoverCompliant(caseID, refID uint64) (string, error) {
	return getDescription(caseID, refID, MethodGetTavanirCoverCompliant)
}

// SetCovered godoc
// @Description مشمول عدم دریافت خسارت
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param desc	 	string	توضیحات علت شمول
// @Returns uint64, error
func SetCovered(caseID uint64, desc string) (uint64, error) {
	jsonData := setCoveredRequest{
		CaseID:      caseID,
		Description: desc,
	}
	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(MethodSetCovered, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return 0, err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result infoUpdateResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}

	if result.Status != 1 {
		return 0, fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return cast.ToUint64(result.Result.RefID), nil
}

// SetAmountEstimate godoc
// @Description اعلم مبلغ تعیین شده برای پرداخت خسارت توسط بیمه
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param amount 	uint64	مبلغ به ریال
// @Param desc	 	string	توضیحات علت شمول
// @Returns uint64, error
func SetAmountEstimate(caseID uint64, amount uint64, desc string) (uint64, error) {
	jsonData := setAmountEstimateRequest{
		CaseID:      caseID,
		Amount:      amount,
		Description: desc,
	}
	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(MethodSetAmountEstimate, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return 0, err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result infoUpdateResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}

	if result.Status != 1 {
		return 0, fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return cast.ToUint64(result.Result.RefID), nil
}

// GetUpdateBankInfo godoc
// @Description دریافت اطلعات بانکی مشترک برای پرداخت خسارت
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param refID 	uint64	شماره ارجاع
// @Returns string, error
func GetUpdateBankInfo(caseID uint64, refID uint64) (BankInfo, error) {
	jsonData := getInfoRequest{
		CaseID: caseID,
		RefID:  refID,
	}
	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(MethodGetUpdateBankInfo, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return BankInfo{}, err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result getUpdateBankInfoResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		//log.Fatal(err)
		return BankInfo{}, err
	}

	if result.Status != 1 {
		return BankInfo{}, fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return result.Result, nil
}

// GetAmountCompliant godoc
// @Description دریافت شکایت مشترک
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param refID 	uint64	شماره ارجاع
// @Returns string, error
func GetAmountCompliant(caseID uint64, refID uint64) (string, error) {
	return getDescription(caseID, refID, MethodGetAmountCompliant)
}

// GetTavanirAmountCompliant godoc
// @Description دریافت شکایت مشترک
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param refID 	uint64	شماره ارجاع
// @Returns string, error
func GetTavanirAmountCompliant(caseID uint64, refID uint64) (string, error) {
	return getDescription(caseID, refID, MethodGetTavanirAmountCompliant)
}

// AddPayment godoc
// @Description پرداخت های چند مرحله ای پس از هر بار پرداخت بجز آخرین پرداخت این وب سرویس را فراخوانی نمایید تا پرداخت در سامانه ثبت شود. آخرین پرداخت را با checkout ثبت نمایید
// @Param caseID 				uint64	شماره پرونده اعلم خسارت
// @Param paymentAmount 		uint64	مبلغ پرداختی به ریال
// @Param description 			string	توضیحات
// @Param paymentDate 			string	YYYY-MM-DD HH24:MI:SS پرداخت تاریخ
// @Param paymentRefID 			string	شماره مرجع پرداخت
// @Param targetBankName 		string	نام بانک
// @Param targetAccountNumber 	string	شماره حساب مشترک
// @Returns uint64, error
func AddPayment(caseID uint64, description string, paymentAmount uint64, paymentDate string, paymentRefID string,
	targetBankName string, targetAccountNumber string) (uint64, error) {
	return addPayment(MethodAddPayment, caseID, description, paymentAmount, paymentDate, paymentRefID,
		targetBankName, targetAccountNumber)
}

// Checkout godoc
// @Description ثبت پرداخت خسارت و ختم پرداخت
// @Param caseID 				uint64	شماره پرونده اعلم خسارت
// @Param paymentAmount 		uint64	مبلغ پرداختی به ریال
// @Param description 			string	توضیحات
// @Param paymentDate 			string	YYYY-MM-DD HH24:MI:SS پرداخت تاریخ
// @Param paymentRefID 			string	شماره مرجع پرداخت
// @Param targetBankName 		string	نام بانک
// @Param targetAccountNumber 	string	شماره حساب مشترک
// @Returns uint64, error
func Checkout(caseID uint64, description string, paymentAmount uint64, paymentDate string, paymentRefID string,
	targetBankName string, targetAccountNumber string) (uint64, error) {
	return addPayment(MethodCheckout, caseID, description, paymentAmount, paymentDate, paymentRefID,
		targetBankName, targetAccountNumber)
}

// CheckoutWithoutPayment godoc
// @Description		ثبت ختم پرداخت بدون ثبت مبلغ
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param desc	 	string	توضیحات درباره علت نیاز به مراجعه مشترک و مدارکی که باید همراه داشته باشد
// @Returns uint64, error
func CheckoutWithoutPayment(caseID uint64, desc string) (uint64, error) {
	return setDescription(MethodCheckoutWithoutPayment, caseID, desc)
}

// PaymentInfoUpdateRequest godoc
// @Description		ایراد در اطلعات بانکی مشترک ب
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param desc	 	string	توضیحات درباره علت نیاز به مراجعه مشترک و مدارکی که باید همراه داشته باشد
// @Returns uint64, error
func PaymentInfoUpdateRequest(caseID uint64, desc string) (uint64, error) {
	return setDescription(MethodPaymentInfoUpdateRequest, caseID, desc)
}

// GetCheckoutCompliant godoc
// @Description علت شکایت مشترک از پرداخت خسارت
// @Param caseID 	uint64	شماره پرونده اعلم خسارت
// @Param refID 	uint64	شماره ارجاع
// @Returns string, error
func GetCheckoutCompliant(caseID uint64, refID uint64) (string, error) {
	return getDescription(caseID, refID, MethodGetCheckoutCompliant)
}

// AddCase godoc
// @Description دریافت مستندات آپلود شده توسط مشترک
// @Param caseAdd 	CaseAdd	شماره پرونده اعلم خسارت
// @Returns []Document, error
func AddCase(caseAdd CaseAdd) (AddCaseResult, error) {
	jsonValue, _ := json.Marshal(caseAdd)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(MethodGetDocumentUpdate, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return AddCaseResult{}, err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result addCaseResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		//log.Fatal(err)
		return AddCaseResult{}, err
	}

	if result.Status != 1 {
		return AddCaseResult{}, fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return result.Result, nil
}

func setDescription(methodName string, caseID uint64, desc string) (uint64, error) {
	jsonData := setDecriptionRequest{
		CaseID:      caseID,
		Description: desc,
	}

	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(methodName, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return 0, err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result infoUpdateResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}

	if result.Status != 1 {
		return 0, fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return cast.ToUint64(result.Result.RefID), nil
}

func addPayment(methodName string, caseID uint64, description string, paymentAmount uint64, paymentDate string, paymentRefID string,
	targetBankName string, targetAccountNumber string) (uint64, error) {
	jsonData := addPaymentRequest{
		CaseID:              caseID,
		Description:         description,
		PaymentAmount:       paymentAmount,
		PaymentDate:         paymentDate,
		PaymentRefID:        paymentRefID,
		TargetBankName:      targetBankName,
		TargetAccountNumber: targetAccountNumber,
	}
	j, _ := json.Marshal(jsonData)
	fmt.Println(string(j))
	response, err := APIRequest(methodName, j)
	if err != nil {
		////log.Fatal(err)
		return 0, err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result infoUpdateResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}

	if result.Status != 1 {
		return 0, fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return cast.ToUint64(result.Result.RefID), nil
}

func getDescription(caseID, refID uint64, methodName string) (string, error) {
	jsonData := getInfoRequest{
		CaseID: caseID,
		RefID:  refID,
	}
	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(methodName, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return "", err
	}
	// fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result getInfoResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		//log.Fatal(err)
		return "", err
	}

	if result.Status != 1 {
		return "", fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return result.Result.Description, nil
}

func getCaseStatus(caseID uint64) (string, error) {
	jsonData := getCaseRequest{
		CaseID: caseID,
	}
	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(string(jsonValue))
	response, err := APIRequest(MethodGetCaseStatus, jsonValue)
	if err != nil {
		////log.Fatal(err)
		return "", err
	}
	//fmt.Println(string(response))
	// err = ioutil.WriteFile("testData.json", response, 0644)
	// return nil, nil
	var result getCseStatusResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		//log.Fatal(err)
		return "", err
	}

	if result.Status != 1 {
		return "", fmt.Errorf("API error: %s", result.Error.Message)
	}

	// //fmt.Println(result.CaseData.ID)

	return result.Result.Status, nil
}
