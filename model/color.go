package model

import (
	"time"
)

type Color struct {
	Id int64 `json:"id"`
	Like int64 `json:"like"`
	Color string `json:"color"`
	Userid int64 `json:"userid"`
	Username string `json:"username"`
	Display bool `json:"display"`
	CreateDate time.Time `json:"createdate"`
}