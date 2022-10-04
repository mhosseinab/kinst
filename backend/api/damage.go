package api

import "github.com/gin-gonic/gin"

// UserCardMe godoc
// @Summary user cards
// @Description user cards
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Accept  json
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Security ApiKeyAuth
// @Tags Card
// @Success 200 {object} entities.UserCardSlice
// @Failure 401 {object} baseError "unauthorized"
// @Router /api/v1/user/card/me/ [get]
func DamagePost(c *gin.Context) {

}
