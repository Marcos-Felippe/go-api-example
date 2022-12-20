package usecase

import (
	"github.com/projetosgo/exemploapi/application/database/repository"
)

type DeleteUserUseCase struct {
	UserRepository *repository.UserRepository
}

func NewDeleteUserUseCase(userRepository *repository.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		UserRepository: userRepository,
	}
}

func (c *DeleteUserUseCase) Execute(id string) error {

	err := c.UserRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
