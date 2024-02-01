package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	expectedName, expectedEmail := "John Doe", "j@j.com"
	user, err := NewUser(expectedName, expectedEmail, "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, expectedName, user.Name)
	assert.NotNil(t, expectedEmail, user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	password, invalidPassword := "1234", "123456"

	user, err := NewUser("John Doe", "j@j.com", password)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(password))
	assert.False(t, user.ValidatePassword(invalidPassword))
	assert.NotEqual(t, password, user.Password)
}
