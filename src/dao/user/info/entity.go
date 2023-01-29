package info

import (
	"time"
)

type UserInfo struct {
	Uuid      string `gorm:"primarykey;type:uuid"`
	Name      string
	Birthday  time.Time
	Gender    string `gorm:"size:1"`
	PhotoURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Form struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday,omitempty"`
	Gender   string `json:"gender,omitempty"`
	PhotoURL string `json:"photo_url,omitempty"`
}
