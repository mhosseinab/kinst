package cmd

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/cobra"
	"es"
	"tools"
)

// storeSearchCmd represents the storeSearch command
var storeSearchCmd = &cobra.Command{
	Use:   "storeSearch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storeSearch called")
		NS := tools.GetEnv("DB_DNS", "kowthar_user:fg4wf5VSERGSWGSAFAAREER@tcp(127.0.0.1:3306)/kowthar_ins?charset=utf8&parseTime=True&loc=Local")

		var err error
		db, err := gorm.Open("mysql", NS)
		if err != nil {
			log.Println(err)
		}
		defer db.Close()
		es.StoreAllRecords(db)
	},
}

func init() {
	rootCmd.AddCommand(storeSearchCmd)
}
