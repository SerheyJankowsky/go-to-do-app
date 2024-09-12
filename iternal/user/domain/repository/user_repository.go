package repository

import "to-do-app/iternal/user/domain/model"



type UserRepository interface {
    Create(user *model.User) error
    FindByID(id uint) (*model.User, error)
    FindByEmail(email string)(*model.User, error)
    Update(user *model.User) error
    Delete(id uint) error
}