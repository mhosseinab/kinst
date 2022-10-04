package api

import (
	"log"
	"time"

	"tavanir"

	"es"

	"github.com/jinzhu/gorm"
	elastic "github.com/olivere/elastic/v7"
	"tools"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var client *elastic.Client

// @title Swagger Damage CRM API
// @version 1.0
// @description This is a Damage CRM server celler server.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// Runner runs api
func Runner() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	NS := tools.GetEnv("DB_DNS", "kowthar_user:fg4wf5VSERGSWGSAFAAREER@tcp(127.0.0.1:3306)/kowthar_ins?charset=utf8&parseTime=True&loc=Local")

	var err error
	db, err = gorm.Open("mysql", NS)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	tavanir.SetDb(db)

	//db.LogMode(true)

	//ups := &models.Request{}
	//if err := db.Model(&models.Request{}).Find(&ups, models.Request{
	//	NationalCode:  "0059072296",
	//	ReferenceCode: "96263785",
	//}).Error; err == nil {
	//	// ups.DamageType = models.RequestDamageTypeFiring
	//	// db.Save(&ups)
	//	db.Delete(&ups)
	//}

	//ups := &models.Request{}
	//if err := db.Model(&models.Request{}).Find(&ups, models.Request{
	//	//NationalCode:  "2679585186",
	//	ReferenceCode: "76386339",
	//}).Error; err == nil {
	//	ups.DamageType = models.RequestDamageTypeFiring
	//	db.Save(&ups)
	//}

	//ups := &tavanir.Case{}
	//if err := db.Model(&tavanir.Case{}).Find(&ups, tavanir.Case{
	//	TrackingID: "9040617894512",
	//}).Error; err == nil {
	//	ups.CompensationTypeID = "6"
	//	db.Save(&ups)
	//}

	cleanUp()

	// db.Model(&models.Request{}).ModifyColumn("expert_description", "text")
	// db.Model(&models.Request{}).ModifyColumn("lack_data_description", "text")

	//db.Exec("drop table `tavanir_damage`")
	//db.Exec("drop table `tavanir_message`")

	//db.AutoMigrate(
	//	&models.Damage{},
	//	&models.Request{},
	//	&models.Storage{},
	//	&models.User{},
	//	&models.RequestChangelog{},
	//	&tavanir.Case{},
	//	&tavanir.Document{},
	//	&tavanir.CaseChangelog{},
	//	&tavanir.StatusUpdateQueue{},
	//	&tavanir.Message{},
	//)

	client, err = elastic.NewClient(
		elastic.SetURL(tools.GetEnv("ES_HOST", "http://127.0.0.1:9200/")),
	)

	r := setupRouter()

	go esStore()

	tavanir.SetDb(db)
	//db.LogMode(true)
	go tavanir.Runner()

	log.Println("listen and serve on http://0.0.0.0:8080")
	log.Fatal(r.Run())
}

func esStore() {
	for {
		es.StoreAllTavanirRecords(db)
		es.StoreAllRecords(db)
		time.Sleep(time.Minute * 10)
	}
}
