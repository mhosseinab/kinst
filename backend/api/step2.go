package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"models"
	"tools"
)

func step2(c *gin.Context) {
	if v := tools.GetEnv("DISABLE_STEPS", "0"); v == "1" {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}
	type fstep2 struct {
		Size string `form:"size" json:"size" binding:"required,oneof=1 2 3"`
	}
	var f fstep2
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

	m.LocationUsage = f.Size
	db.Save(&m)

	c.JSON(200, gin.H{"status": 200})
}
