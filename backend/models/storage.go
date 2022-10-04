package models

import (
	"crypto/sha1"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/segmentio/ksuid"
)

// Storage struct represents file uploader
type Storage struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Path           string    `json:"-"`
	URL            string    `gorm:"url" json:"url"`
	UploadFileName string    `json:"upload_file_name"`
}

func getEnv(key, def string) string {
	if v, ex := os.LookupEnv(key); ex {
		return v
	}
	return def
}

// NewStorage returns new storage :D
func NewStorage(filename string) Storage {
	mediaRoot := getEnv("MEDIA_ROOT", "media")
	cdnPrefix := getEnv("CDN_PREFIX", "http://127.0.0.1:8080/")
	prefix := ksuid.New().String()

	filename = getSha1(filename) + filepath.Ext(filename)
	fprefix := prefix + "_" + filename
	s := Storage{
		UploadFileName: filename,
		Path:           path.Join(mediaRoot, "/storage/", fprefix),
		URL:            cdnPrefix + "media/storage/" + fprefix,
	}

	return s
}

func fixFilename(s string) string {
	s = url.QueryEscape(s)
	return s
}

func getSha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
