package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyName_WhereCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{}
	assert.Error(t, user.Validate(), "invalid name")
}

func TestGivenAnEmptyEmail_WhereCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{Name: "mf"}
	assert.Error(t, user.Validate(), "invalid email")
}

func TestGivenAValidParams_WhenCallNewUser_ThenShouldReceiverCreateUserWithAllParams(t *testing.T) {
	user := User{
		Name:  "mf",
		Email: "mf@email",
	}
	assert.Equal(t, "mf", user.Name)
	assert.Equal(t, "mf@email", user.Email)
	assert.Nil(t, user.Validate())
}

func TestGivenAValidParams_WhenCallNewUserFunc_ThenShouldReceiverCreateUserWithAllParams(t *testing.T) {
	user, err := NewUser(
		"mf",
		"mf@email",
	)
	assert.Nil(t, err)
	assert.Equal(t, "mf", user.Name)
	assert.Equal(t, "mf@email", user.Email)
}

func TestGivenAValidParams_WhenICallGenerateID_ThenIShouldSetID(t *testing.T) {
	user, err := NewUser(
		"mf",
		"mf@email",
	)
	assert.Nil(t, err)
	assert.Nil(t, user.GenerateID())
}
