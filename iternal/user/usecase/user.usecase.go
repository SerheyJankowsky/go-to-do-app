package usecase

import "to-do-app/iternal/user/domain/model"

type UserUseCase interface{
	CreateUser(name string, password string) (*model.User, error)
    GetUser(id uint) (*model.User, error)
    UpdateUser(id uint, name string) (*model.User, error)
    DeleteUser(id uint) error
}