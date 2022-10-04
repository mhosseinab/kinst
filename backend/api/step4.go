package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"models"
	"tools"
)

func step4(c *gin.Context) {
	if v := tools.GetEnv("DISABLE_STEPS", "0"); v == "1" {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}
	type fstep4 struct {
		LastBillPhoto string `form:"last_bill_photo" json:"last_bill_photo" binding:"required,url"`
		IDCardPhoto   string `form:"id_card_photo" json:"id_card_photo" binding:"required,url"`
		OtherPhoto    string `form:"other_photo" json:"other_photo"`
	}
	var f fstep4
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

	m.LastBillPhoto = f.LastBillPhoto
	m.IDCardPhoto = f.IDCardPhoto
	m.OtherPhoto = f.OtherPhoto
	db.Save(&m)

	c.JSON(200, gin.H{"status": 200})
	return
}
