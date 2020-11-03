package model

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

type AccessToken struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID      uuid.UUID
	Name        string
	Description string
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}
