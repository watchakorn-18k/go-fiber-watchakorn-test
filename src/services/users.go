package services

import (
	"go-fiber-proton/domain/entities"
	"go-fiber-proton/domain/repositories"
)

type usersService struct {
	UsersRepository repositories.IUsersRepository
}

type IUsersService interface {
	GetAllUser() (*[]entities.UserDataFormat, error)
	InsertNewAccount(data *entities.NewUserBody) bool
	UpdateUser(data *entities.NewUserBody) bool
	DeleteUser(name string) bool
}

func NewUsersService(repo0 repositories.IUsersRepository) IUsersService {
	return &usersService{
		UsersRepository: repo0,
	}
}

func (sv usersService) GetAllUser() (*[]entities.UserDataFormat, error) {
	userData, err := sv.UsersRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return userData, nil

}

func (sv usersService) InsertNewAccount(data *entities.NewUserBody) bool {
	if sv.UsersRepository.FindUser(data.Name) {
		return false
	}
	status := sv.UsersRepository.InsertNewUser(data)
	return status
}

func (sv usersService) UpdateUser(data *entities.NewUserBody) bool {
	if !sv.UsersRepository.FindUser(data.Name) {
		return false
	}
	status := sv.UsersRepository.UpdateUser(data)
	return status
}

func (sv usersService) DeleteUser(name string) bool {
	if !sv.UsersRepository.FindUser(name) {
		return false
	}
	status := sv.UsersRepository.DeleteUser(name)
	return status
}
