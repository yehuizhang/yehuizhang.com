package shared

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	Uuid      uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
