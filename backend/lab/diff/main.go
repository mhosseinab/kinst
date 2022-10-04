package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/r3labs/diff"
	"models"
	"tools"
	"gopkg.in/d4l3k/messagediff.v1"
)

func main() {
	a := models.Request{
		DamageType:  models.RequestDamageTypeExplosion,
		UpdatedAt:   time.Now().Add(-time.Hour * 1),
		Description: "salam",
	}
	b := models.Request{
		DamageType:  models.RequestDamageTypeDeath,
		UpdatedAt:   time.Now(),
		Description: "salam1",
	}
	dif, equal := messagediff.PrettyDiff(a, b)

	if !equal {
		log.Println(dif)
	}

	changelog, err := diff.Diff(a, b)
	log.Println("l", changelog, err, changelog == nil)
	spew.Dump(changelog)
	jm, _ := json.Marshal(changelog)
	log.Println(string(jm))

	NS := tools.GetEnv("DB_DNS", "kowthar_user:fg4wf5VSERGSWGSAFAAREER@tcp(127.0.0.1:3306)/kowthar_ins?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open("mysql", NS)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	cl := models.RequestChangelog{
		Changelogs: string(jm),
		CreatedAt:  time.Now(),
		RequestID:  1,
	}
	cl.UserID = 0
	db.Create(&cl)
}
