package repository

import "to-do-app/iternal/auth/domain/model"

type TokenRepository interface{
	Create(t *model.Token)error
	FindByID(id uint) (*model.Token,error)
	FindByUserId(id uint)(*model.Token,error)
	Delete(id uint)error
}