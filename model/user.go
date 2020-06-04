package model

import (
	"time"
)

type User struct {
	Id int64 `json:"id"`
	OAuth string `json:"oauth"`
	Name string `json:"name"`
	OAuthId string `json:"oauthid"`
	IsAdmin bool `json:"isadmin"`
	LastLogin time.Time `json:"lastlogin"`
}