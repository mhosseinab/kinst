package admin

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func postResource(c *gin.Context, val interface{}) (interface{}, error) {
	if err := c.ShouldBindJSON(&val); err != nil {
		return nil, err
	}

	err := db.Create(reflect.ValueOf(val).Interface()).Error
	return val, err
}

func deleteResource(c *gin.Context, val, id interface{}) error {
	obj, err := getResource(c, val, id)
	if err != nil {
		return err
	}

	return db.Delete(obj).Error
}

func putResource(c *gin.Context, val, id interface{}) (interface{}, error) {
	obj, err := getResource(c, val, id)
	if err != nil {
		return nil, err
	}

	if err := c.ShouldBindJSON(&obj); err != nil {
		return nil, err
	}

	err = db.Save(reflect.ValueOf(val).Interface()).Error
	return val, err
}

func getResource(c *gin.Context, val interface{}, id interface{}) (interface{}, error) {
	err := db.Find(reflect.ValueOf(val).Interface(), "id=?", id).Error
	return val, err
}

func getResourceList(c *gin.Context, val interface{}) (interface{}, int, error) {
	reqVals := c.Request.URL.Query()
	d := db.Model(val)

	if id := reqVals.Get("id"); id != "" {
		d = d.Where("id=?", id)
	}

	if q := reqVals.Get("q"); q != "" {
		vo, vok := val.(resourceObject)
		if vok {
			for _, field := range getResourceSearchField(vo) {
				qf := fmt.Sprintf("%s like ?", field)
				d = d.Where(qf, "%"+q+"%")
			}
		} else {
			log.Println("unhandled get getResourceSearchField", reflect.ValueOf(val).String())
		}
	}

	var count int
	d = d.Count(&count)

	if sortField := reqVals.Get("_sort"); sortField != "" {
		if order := reqVals.Get("_order"); order != "" {
			ord := fmt.Sprintf("%s %s", sortField, order)
			d = d.Order(ord)
		}
	}

	if start := reqVals.Get("_start"); start != "" {
		d = d.Offset(start)
		if end := reqVals.Get("_end"); end != "" {
			d = d.Limit(cast.ToInt(end) - cast.ToInt(start))
		}
	}

	err := d.Find(val).Error

	return val, count, err
}
