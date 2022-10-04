package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func writeJSONorError(c *gin.Context, object interface{}, err error) {
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, object)
}

func listHandler(c *gin.Context, entity interface{}) {
	r, count, err := getResourceList(c, entity)
	c.Header("X-Total-Count", cast.ToString(count))
	writeJSONorError(c, r, err)
}

func itemHandler(c *gin.Context, entity, id interface{}) {
	r, err := getResource(c, entity, id)
	writeJSONorError(c, r, err)
}

func postHandler(c *gin.Context, entity interface{}) {
	r, err := postResource(c, entity)
	writeJSONorError(c, r, err)
}

func putHandler(c *gin.Context, entity, id interface{}) {
	r, err := putResource(c, entity, id)
	writeJSONorError(c, r, err)
}

func deleteHandler(c *gin.Context, entity, id interface{}) {
	err := deleteResource(c, entity, id)
	writeJSONorError(c, "deleted", err)
}
