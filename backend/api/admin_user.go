package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"models"
	"golang.org/x/crypto/bcrypt"
)

func adminUserList(c *gin.Context) {
	user := getUserFromSession(c)
	if user.Role != models.UserRoleAdmin {
		c.JSON(http.StatusBadRequest, fmt.Errorf("permission error"))
		return
	}

	offset, limit, order := getRequestFilters(c)

	var count int
	var ul []models.User
	db.Model(&models.User{}).
		Count(&count).
		Offset(offset).
		Order(order).
		Limit(limit).
		Find(&ul)

	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"total_count": count,
		},
		"objects": ul,
	})
}

func adminUserItem(c *gin.Context) {
	id := c.Param("id")

	user := getUserFromSession(c)
	if user.Role != models.UserRoleAdmin {
		c.JSON(http.StatusBadRequest, fmt.Errorf("permission error"))
		return
	}

	var u models.User
	db.Model(&models.User{}).Find(&u, "id=?", id)

	c.JSON(http.StatusOK, u)
}

func adminUserCreate(c *gin.Context) {
	type createUserForm struct {
		Username    string   `json:"username" binding:"required"`
		Password    string   `json:"password" binding:"required"`
		Role        string   `json:"role" binding:"required"`
		Description string   `json:"description"`
		Province    []string `json:"province"`
	}

	var f createUserForm
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.MinCost)

	ua := models.User{
		Username:    f.Username,
		Password:    string(hash),
		Role:        f.Role,
		Description: f.Description,
		Province:    strings.Join(f.Province, ","),
	}
	if err := db.Create(&ua).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, ua)

}

func adminUserChangePassword(c *gin.Context) {
	type editUserForm struct {
		Password string `form:"password" json:"password" binding:"required,min=6,max=24"`
	}

	var f editUserForm
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_msg": fmt.Errorf("رمز عبور حداقل 6 و حداکثر 24 کاراکتر باشد").Error(),
		})
		return
	}

	user := getUserFromSession(c)

	//if user.UserID == 1 {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error_msg": fmt.Errorf("امکان تغییر این کاربر وجود ندارد").Error(),
	//	})
	//	return
	//}

	var u models.User
	if err := db.Model(&models.User{}).First(&u, "id=?", user.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if f.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.MinCost)
		u.Password = string(hash)
	}

	if err := db.Save(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

func adminUserEdit(c *gin.Context) {
	type editUserForm struct {
		Username    string   `json:"username" binding:"required"`
		Password    string   `json:"password"`
		Role        string   `json:"role" binding:"required"`
		Description string   `json:"description"`
		Province    []string `json:"province"`
		Status      string   `json:"status"`
	}

	var f editUserForm
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")
	if id == "1" {
		c.JSON(http.StatusBadRequest, fmt.Errorf("امکان تغییر این کاربر وجود ندارد"))
		return
	}
	var u models.User
	if err := db.Model(&models.User{}).First(&u, "id=?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	u.Username = f.Username
	u.Role = f.Role
	u.Status = f.Status
	u.Description = f.Description
	u.Province = strings.Join(f.Province, ",")

	if f.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.MinCost)
		u.Password = string(hash)
	}

	if err := db.Save(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}
