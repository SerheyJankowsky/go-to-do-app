package usecase

import (
	itemModel "to-do-app/iternal/item/domain/model"
	userModel "to-do-app/iternal/user/domain/model"
	"to-do-app/iternal/user/domain/repository"
)

type UserUseCaseImpl struct{
	repo repository.UserRepository
}

func NewUserUseCaseImpl(repo repository.UserRepository) *UserUseCaseImpl{
	return &UserUseCaseImpl{
		repo: repo,
	}
}

func (u *UserUseCaseImpl) CreateUser(name string,password string) (*userModel.User,error){
	user := &userModel.User{
		Name: name,
		Password: password,
		Items: []itemModel.Item{},
	}
	if err := u.repo.Create(user);err != nil{
		return nil,err
	}
	return user,nil
}

func (u *UserUseCaseImpl) GetUser(id uint)(*userModel.User,error){
	return u.repo.FindByID(id)
}

func (u *UserUseCaseImpl) UpdateUser(id uint,name string)(*userModel.User,error){
	user,err := u.repo.FindByID(id)
	if err !=nil{
		return nil,err
	}
	user.Name = name
	if err:=u.repo.Update(user);err!=nil{
		return nil,err
	}
	return user,nil
}

func (u *UserUseCaseImpl) DeleteUser (id uint) error{
	return u.repo.Delete(id)
}