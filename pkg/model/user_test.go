package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	usr := NewUser("testuser")

	assert.NotNil(t, usr)
	assert.NotEmpty(t, usr.ID)
	assert.Equal(t, usr.Login, "testuser")
	assert.NotNil(t, usr.CreatedAt)
	assert.NotNil(t, usr.UpdatedAt)
	assert.IsType(t, time.Time{}, usr.CreatedAt)
	assert.IsType(t, time.Time{}, usr.UpdatedAt)
}
