package shared

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	Uuid      uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
