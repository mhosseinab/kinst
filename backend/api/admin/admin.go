package admin

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

// SetDB sets db on admin api
func SetDB(d *gorm.DB) {
	db = d
}

type resourceObject interface {
	GetSearchFields() []string
}

func getResourceSearchField(r resourceObject) []string {
	return r.GetSearchFields()
}

// RestFullResources
func RestFullResources(c *gin.Context) {
	reqP := strings.Split(c.Param("req"), "/")

	var isList bool
	if len(reqP) == 2 {
		isList = true
	}
	entity, err := getEntityFromReq(reqP[1], isList)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	switch c.Request.Method {
	case http.MethodGet:
		if len(reqP) == 2 {
			listHandler(c, entity)
		} else if len(reqP) == 3 {
			itemHandler(c, entity, reqP[2])
		}
	case http.MethodPost:
		postHandler(c, entity)
	case http.MethodPut:
		putHandler(c, entity, reqP[2])
	case http.MethodDelete:
		deleteHandler(c, entity, reqP[2])
	}

}
