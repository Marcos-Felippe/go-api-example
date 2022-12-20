package entity

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name string, email string) (*User, error) {
	user := &User{
		Name:  name,
		Email: email,
	}
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("invalid name")
	}
	if u.Email == "" {
		return errors.New("invalid email")
	}
	return nil
}

func (u *User) GenerateID() error {
	u.ID = uuid.New().String()
	err := u.Validate()
	if err != nil {
		return err
	}
	return nil
}
