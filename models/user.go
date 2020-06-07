package models

import (
	"time"
)

// User definition
type User struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	OAuth     string    `json:"oauth"`
	Name      string    `json:"name"`
	OAuthID   string    `json:"oauthid"`
	IsAdmin   bool      `json:"isadmin"`
	LastLogin time.Time `json:"lastlogin"`
}
