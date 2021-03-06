package model

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User describes account attributes for a registered user.
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Login     string
	Email     string
	Password  []byte
	AvatarURL url.URL
	Name      string
	Bio       string
	Location  string
	Company   string
	Birthday  sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime

	Emails       []Email
	AccessTokens []AccessToken
}

// SetPassword overwrites the current password with hash value generated by
// the given value.
func (u *User) SetPassword(v string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = hash
	return nil
}

// ComparePassword checks if the given password compares to the current value.
func (u *User) ComparePassword(v string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(v))
}
