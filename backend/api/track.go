package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"models"
)

type checkResult struct {
	BillIndetifier string `json:"bill_indetifier" form:"bill_indetifier" binding:"required"`
	TrackingCode   string `json:"tracking_code" form:"tracking_code" binding:"required"`
}

func trackResult(c *gin.Context) {
	var f checkResult
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(400, gin.H{"error_msg": formDataValidationError})
		return
	}

	var m models.Request
	if err := db.Model(&models.Request{}).First(&m, "bill_identifier=? and (reference_code=? or national_code=?)", f.BillIndetifier, f.TrackingCode, f.TrackingCode).Error; err != nil {
		c.JSON(400, gin.H{"error_msg": "اطلاعات مورد نظر یافت نشد."})
		return
	}

	c.JSON(http.StatusOK, m)
}
