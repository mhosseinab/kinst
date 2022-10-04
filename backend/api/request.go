package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff"
	"connections"
	"es"
	"models"
)

// UpdateRequest by user
// http://kins-gateway.abrbit.com/api/v1/request/update/
// add token authorization header
// bosy json
func UpdateRequest(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == "" {
		c.JSON(400, gin.H{"error": fmt.Errorf("token not set").Error()})
		return
	}

	var m models.Request
	if err := db.First(&m, "token=?", reqToken).Error; err != nil {
		log.Println(err.Error())
		c.JSON(400, err.Error())
		return
	}

	beforeChange := m

	curStatus := m.Status

	switch m.Status {
	case models.RequestStatusCanceledByUser,
		models.RequestStatusPayed,
		models.RequestStatusClosed:
		c.JSON(400, gin.H{"error_msg": "امکان ویرایش وجود ندارد!"})
		return
	}

	if err := c.ShouldBind(&m); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error_msg": err.Error()})
		return
	}

	if curStatus == models.RequestStatusIncomplete {
		m.Status = models.RequestStatusIncompleteChange
	}

	if err := db.Save(&m).Error; err != nil {
		log.Println(err.Error())
	}

	changelog, _ := diff.Diff(beforeChange, m)
	if changelog != nil && m.Status != models.RequestStatusIncomplete && len(changelog) > 1 {
		changelogByte, _ := json.Marshal(changelog)
		cl := models.RequestChangelog{
			Changelogs: string(changelogByte),
			CreatedAt:  time.Now(),
			RequestID:  m.ID,
		}
		cl.UserID = 0
		db.Create(&cl)
	}

	if client, err := connections.GetElasticsearch(); err == nil {
		es.StoreRequestItem(client, m)
	}

	c.JSON(http.StatusOK, m)
}
