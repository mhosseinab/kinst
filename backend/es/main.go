package es

import (
	"context"
	"log"

	"tavanir"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"models"
	"tools"

	elastic "github.com/olivere/elastic/v7"
)

// StoreAllRecords stores all records from database to elasticsearch
func StoreAllRecords(db *gorm.DB) {
	var request models.Request
	rows, err := db.Model(&models.Request{}).Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}

	client, err := elastic.NewClient(
		elastic.SetURL(tools.GetEnv("ES_HOST", "http://127.0.0.1:9200/")),
	)
	if err != nil {
		log.Println(err.Error())
		log.Println(tools.GetEnv("ES_HOST", "http://127.0.0.1:9200/"))
		return
	}

	for rows.Next() {
		db.ScanRows(rows, &request)
		err := StoreRequestItem(client, request)
		if err != nil {
			return
		}
	}
}

// StoreRequestItem stores request item in elasticsearch
func StoreRequestItem(client *elastic.Client, r models.Request) error {
	if r.Status == models.RequestStatusNotCompleted {
		return nil
	}
	_, err := client.Index().
		Index(tools.GetEnv("ES_INDEX", "request")).
		Type("item").
		Id(cast.ToString(r.ID)).
		BodyJson(r).
		Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// Tavanir elastic storage

// StoreAllRecords stores all records from database to elasticsearch
func StoreAllTavanirRecords(db *gorm.DB) {
	var tc tavanir.Case
	rows, err := db.Model(&tavanir.Case{}).Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}

	client, err := elastic.NewClient(
		elastic.SetURL(tools.GetEnv("ES_HOST", "http://127.0.0.1:9200/")),
	)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for rows.Next() {
		db.ScanRows(rows, &tc)
		err := StoreTavanirItem(client, tc)
		if err != nil {
			return
		}
	}
}

// StoreRequestItem stores request item in elasticsearch
func StoreTavanirItem(client *elastic.Client, r tavanir.Case) error {
	if r.Status == models.RequestStatusNotCompleted {
		return nil
	}
	re := r.ToElastic()
	_, err := client.Index().
		Index(tools.GetEnv("ES_TAVANIR_INDEX", "case")).
		Type("item").
		Id(cast.ToString(r.Id)).
		BodyJson(re).
		Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
