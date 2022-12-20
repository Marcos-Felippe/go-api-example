package usecase

import (
	"github.com/projetosgo/exemploapi/application/database/repository"
	"github.com/projetosgo/exemploapi/entity"
)

type GetAllUseCase struct {
	UserRepository *repository.UserRepository
}

func NewGetAllUseCase(userRepository *repository.UserRepository) *GetAllUseCase {
	return &GetAllUseCase{
		UserRepository: userRepository,
	}
}

func (c *GetAllUseCase) Execute() ([]entity.User, error) {

	users, err := c.UserRepository.GetAll()
	if err != nil {
		return users, err
	}

	return users, nil
}
