package api

import (
	"io"
	"os"

	"models"

	"github.com/gin-gonic/gin"
)

func upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.JSON(400, err)
		return
	}
	filename := header.Filename

	s := models.NewStorage(filename)

	err = os.MkdirAll("/home/kowthar/www/media/storage/", 0777)
	if err != nil {
		c.JSON(400, err)
		return
	}

	out, err := os.Create(s.Path)
	if err != nil {
		c.JSON(400, err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(400, err)
		return
	}

	if err := db.Create(&s).Error; err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(200, s)
}
