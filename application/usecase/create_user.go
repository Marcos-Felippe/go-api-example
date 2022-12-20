package usecase

import (
	"github.com/projetosgo/exemploapi/application/database/repository"
	"github.com/projetosgo/exemploapi/entity"
)

type UserInputDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserUseCase struct {
	UserRepository *repository.UserRepository
}

func NewCreateUserUseCase(userRepository *repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (c *CreateUserUseCase) Execute(input UserInputDTO) (*UserOutputDTO, error) {
	user, err := entity.NewUser(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = user.GenerateID()
	if err != nil {
		return nil, err
	}

	err = c.UserRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
