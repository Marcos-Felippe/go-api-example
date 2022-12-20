package usecase

import (
	"github.com/projetosgo/exemploapi/application/database/repository"
	"github.com/projetosgo/exemploapi/entity"
)

type UpdateUserUseCase struct {
	UserRepository *repository.UserRepository
}

func NewUpdateUserUseCase(userRepository *repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UserRepository: userRepository,
	}
}

func (c *UpdateUserUseCase) Execute(id string, input UserInputDTO) (*UserOutputDTO, error) {
	user, err := entity.NewUser(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	user.ID = id

	user, err = c.UserRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
