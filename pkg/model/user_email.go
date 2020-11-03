package model

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

type Email struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID     uuid.UUID
	Email      string
	IsVerified bool
	VerifiedAt sql.NullTime
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
