package models

import (
	"time"
)

// Color is definition
type Color struct {
	ID         int64     `json:"id" gorm:"primary_key" binding:"required"`
	Like       int64     `json:"like"`
	Color      string    `json:"color"`
	UserID     int64     `json:"userid"`
	Username   string    `json:"username"`
	Display    bool      `json:"display"`
	CreateDate time.Time `json:"createdate"`
}
