package api

import (
	"fmt"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
	"models"
	"tools"
)

func step3(c *gin.Context) {
	if v := tools.GetEnv("DISABLE_STEPS", "0"); v == "1" {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}
	type fstep3 struct {
		Firstname      string `form:"firstname" json:"firstname" binding:"required"`
		Surname        string `form:"surname" json:"surname"`
		NationalCode   string `form:"national_code" json:"national_code"`
		MobileNumber   string `form:"mobile_number" json:"mobile_number" binding:"required"`
		Province       string `form:"province" json:"province" binding:"required"`
		City           string `form:"city" json:"city" binding:"required"`
		Address        string `form:"address" json:"address" binding:"required"`
		SubscriberType string `json:"subscriber_type" binding:"required"`
		Sheba          string `json:"sheba"`
		PostalAddress  string `form:"postal_address" json:"postal_address" binding:"required"`
		EconomicCode   string `json:"economic_code"`
	}
	var f fstep3
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}

	if f.SubscriberType == "1" {
		if f.NationalCode == "" || f.Surname == "" {
			c.JSON(400, gin.H{"error_msg": formDataValidationError})
			return
		}
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

	m.Firstname = f.Firstname
	m.Surname = f.Surname
	m.NationalCode = f.NationalCode
	m.MobileNumber = f.MobileNumber
	m.Province = f.Province
	m.Sheba = f.Sheba
	m.City = f.City
	m.SubscriberType = 1
	if cast.ToInt(f.SubscriberType) != 0 {
		m.SubscriberType = cast.ToInt(f.SubscriberType)
	}
	m.Address = f.Address
	m.PostalAddress = f.PostalAddress
	m.EconomicCode = f.EconomicCode

	db.Save(&m)

	c.JSON(200, gin.H{"status": 200})
}
