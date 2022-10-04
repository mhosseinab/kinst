/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"tools"

	"tavanir"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

// tavanirFetchCmd represents the tavanirFetch command
var tavanirFetchCmd = &cobra.Command{
	Use:   "tavanirFetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tavanirFetch called")
		log.SetFlags(log.LstdFlags | log.Lshortfile)

		NS := tools.GetEnv("DB_DNS", "kowthar_user:fg4wf5VSERGSWGSAFAAREER@tcp(127.0.0.1:3306)/kowthar_ins?charset=utf8&parseTime=True&loc=Local")

		db, err := gorm.Open("mysql", NS)
		if err != nil {
			log.Println(err)
		}
		defer db.Close()
		//db.LogMode(true)

		db.AutoMigrate(
			&tavanir.Case{},
			&tavanir.Document{},
			&tavanir.CaseChangelog{},
		)

		tavanir.SetDb(db)
		tavanir.Runner()
	},
}

func init() {
	rootCmd.AddCommand(tavanirFetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tavanirFetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tavanirFetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
