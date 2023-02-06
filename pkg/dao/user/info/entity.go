package info

import (
	"time"
)

type UserInfo struct {
	Id        string `gorm:"primarykey;type:uuid"`
	Name      string
	Birthday  time.Time
	Gender    string `gorm:"size:1"`
	PhotoURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Form struct {
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	PhotoURL string    `json:"photo_url"`
}
