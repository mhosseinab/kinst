package api

import (
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"models"
	"tools"
)

func step1(c *gin.Context) {
	if v := tools.GetEnv("DISABLE_STEPS", "0"); v == "1" {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}
	type fstep1 struct {
		BillIdentifier string `form:"bill_identifier" json:"bill_identifier" binding:"required,lte=13,gte=6"`
		CompanyID      string `json:"company_id" binding:"required"`
	}
	var f fstep1
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}

	if f.CompanyID == "062" || f.CompanyID == "061" {
		c.JSON(400, gin.H{"error_msg": "مهلت ثبت خسارت به پایان رسیده"})
		return
	}

	m := &models.Request{
		BillIdentifier: f.BillIdentifier,
		CompanyID:      f.CompanyID,
		Token:          guuid.New().String(),
	}

	if err := db.Create(m).Error; err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{"status": 200, "token": m.Token})
}
