/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tiaguinho/gosoap"
)

// smsCmd represents the sms command
var smsCmd = &cobra.Command{
	Use:   "sms",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sms called")
		sendSMS()
	},
}

func init() {
	rootCmd.AddCommand(smsCmd)
}

type sms struct {
	Text string `xlm:"text"`
}

func sendSMS() {

	referenceCode := "12345678"
	phone := "989032154149"
	soap, err := gosoap.SoapClient("http://example.com:8083/App_WebService/PortalSmsPanelService.asmx?wsdl")
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	t := fmt.Sprintf(`بیمه شده عزیز
	کد رهگیری شما %s در سامانه بیمه کوثر ثبت شد وتا ۴۸ ساعت آینده همکاران ما با شما تماس خواهند گرفت.`, referenceCode)

	params := gosoap.Params{
		"MessageText":     t,
		"Mobile":          phone,
		"UserNameService": "123456",
		"PassWordService": "123456",
	}

	err = soap.Call("SmsPanelContactSystem_QuickSmsMsTavanir", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}

}
