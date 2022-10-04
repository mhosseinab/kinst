package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"models"
	"tools"
)

func step5(c *gin.Context) {
	if v := tools.GetEnv("DISABLE_STEPS", "0"); v == "1" {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}
	type fstep5 struct {
		CasualityDate string `form:"casuality_date" json:"casuality_date" binding:"required"`
		CausalityTime string `form:"casuality_time" json:"casuality_time" binding:"required"`
		Description   string `form:"description" json:"description"`
		LocationType  int    `form:"location_type" json:"location_type"`
	}
	var f fstep5
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}

	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == "" {
		c.JSON(400, gin.H{"error": fmt.Errorf("token not set").Error()})
		return
	}

	var m models.Request
	if err := db.First(&m, "token=?", reqToken).Error; err != nil {
		c.JSON(400, err.Error())
		return
	}

	t1, _ := time.Parse(time.RFC3339, f.CasualityDate)
	rounded := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())

	var mold models.Request
	err := db.First(&mold, "bill_identifier=? and damage_type not in (?)", m.BillIdentifier, []string{
		models.RequestDamageTypeLack, models.RequestDamageTypeDeath}).Error
	if err != nil {
	}
	if mold.ReferenceCode != "" && mold.CasualityDate.Truncate(24*time.Hour).Equal(rounded) {
		c.JSON(400, gin.H{
			"error_msg": "با این شناسه قبض و تاریخ قبلا خسارت ثبت شده است.",
		})
		return
	}

	m.CasualityDate = rounded
	m.CasualityTime = f.CausalityTime
	m.Description = f.Description
	m.LocationType = f.LocationType
	db.Save(&m)

	c.JSON(200, gin.H{"status": 200})
}
