package models

import (
	"time"

	"gorm.io/gorm"
)

//User ...
type User struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	FirstName     string         `json:"first_name" gorm:"first_name"` //nolint:gofmt
	LastName      string         `json:"last_name" gorm:"last_name"`   //nolint:gofmt
	Email         string         `json:"email" gorm:"email"`
	Password      string         `json:"-" gorm:"password"`
	Balance       float32        `json:"balance" gorm:"balance"`
	EmailVerified bool           `json:"email_verified" gorm:"email_verified"`
	Currency      string         `json:"currency" gorm:"currency"`
	IsAdmin       bool           `json:"is_admin" gorm:"is_admin"`
	// Files         []File         `gorm:"many2many:user_files;" json:"files,omitempty"`
	// Metas         []UserMeta     `gorm:"many2many:user_meta;" json:"metas,omitempty"`
	// UserSetting   UserSetting    `gorm:"-" json:"settings,omitempty"`
}

//Role for user
type Role struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `json:"name" gorm:"name"`               //nolint:gofmt
	Slug        string `json:"slug" gorm:"slug"`               //nolint:gofmt
	Description string `json:"description" gorm:"description"` //nolint:gofmt
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
