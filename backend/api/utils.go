package api

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("123456789")

// RandStringRunes returns random strings by runes
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getUserFromSession(c *gin.Context) *KinsClaims {
	curUser, _ := c.Get("user")
	return curUser.(*KinsClaims)
}
