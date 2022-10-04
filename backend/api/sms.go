package api

import (
	"fmt"
	"log"

	"github.com/tiaguinho/gosoap"
)

const (
	smsService = "http://webservice.example.com/App_WebService/PortalSmsPanelService.asmx?wsdl"

	smsServiceUsername = "123456"
	smsServicePassword = "123456"
)

func sendSms(params gosoap.Params) error {
	soap, err := gosoap.SoapClient(smsService)
	if err != nil {
		log.Printf("SoapClient error: %s\n", err)
		return err
	}

	err = soap.Call("SmsPanelContactSystem_QuickSmsMsTavanir", params)
	if err != nil {
		log.Printf("Call error: %s\n", err)
		return err
	}
	return nil
}

func sendSmsTracking(referenceCode, mobile string) {
	txt := `بیمه شده محترم؛
مدارک شما با کدرهگیری %s در سامانه شرکت بیمه کوثر ثبت گردید. نتیجه پس از بررسی اسناد از طریق پیامک به شما اطلاع رسانی خواهد شد.`

	params := gosoap.Params{
		"MessageText":     fmt.Sprintf(txt, referenceCode),
		"Mobile":          mobile,
		"UserNameService": smsServiceUsername,
		"PassWordService": smsServicePassword,
	}
	sendSms(params)
}

func sendSmsChaneStatus(mobile string) {
	t := `بیمه شده محترم
در خواست خسارت شما مورد بررسی قرار گرفت جهت مشاهده آخرین وضعیت به قسمت پیگیری سایت tavanir.example.com مراجعه نمایید`

	params := gosoap.Params{
		"MessageText":     t,
		"Mobile":          mobile,
		"UserNameService": smsServiceUsername,
		"PassWordService": smsServicePassword,
	}

	sendSms(params)
}

func sendSmsTavanirChaneStatus(mobile string) {
	t := `بیمه شده محترم
در خواست خسارت شما مورد بررسی قرار گرفت جهت مشاهده آخرین وضعیت به سامانه ذیل قسمت پیگیری مراجعه نمایید.
http://bime.tavanir.org.ir`

	params := gosoap.Params{
		"MessageText":     t,
		"Mobile":          mobile,
		"UserNameService": smsServiceUsername,
		"PassWordService": smsServicePassword,
	}

	sendSms(params)
}
