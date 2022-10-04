package models

import "time"

// user roles
const (
	UserRoleAdmin    = "ADMIN"
	UserRoleBranch   = "BRANCH"
	UserRoleTavanir  = "TAVANIR"
	UserRoleReporter = "REPORTER"
)

// user status
const (
	UserStatusActive   = "ACTIVE"
	UserStatusInactive = "IN_ACTIVE"
)

// User Struct
type User struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"-"`
	LastLogin   *time.Time `json:"last_login"`
	Username    string     `gorm:"type:varchar(100);unique" json:"username"`
	Password    string     `gorm:"type:varchar(100)" json:"-"`
	Role        string     `gorm:"type:varchar(20)" json:"role"`
	Status      string     `gorm:"type:varchar(20);default:'ACTIVE'" json:"status"`
	Description string     `json:"description"`
	Province    string     `json:"province"`
}
