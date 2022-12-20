package usecase

import (
	"github.com/projetosgo/exemploapi/application/database/repository"
)

type GetUserUseCase struct {
	UserRepository *repository.UserRepository
}

func NewGetUserUseCase(userRepository *repository.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{
		UserRepository: userRepository,
	}
}

func (c *GetUserUseCase) Execute(id string) (*UserOutputDTO, error) {

	user, err := c.UserRepository.GetOne(id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
