package models

import (
	"time"

	"github.com/r3labs/diff"
)

// RequestChangelog struct
type RequestChangelog struct {
	ID            uint           `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	RequestID     uint           `gorm:"index:request_id" json:"request_id"`
	Changelogs    string         `gorm:"type:text" json:"-"`
	ChangelogData diff.Changelog `gorm:"-" json:"changelog"`
	UserID        uint           `gorm:"index:user_id" json:"user_id"`
}

type RequestChangelogSlice []RequestChangelog
