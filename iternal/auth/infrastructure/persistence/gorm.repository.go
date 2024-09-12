package persistence

import (
	"errors"
	"to-do-app/iternal/auth/domain/model"

	"gorm.io/gorm"
)

type GormTokenRepository struct{
	db *gorm.DB
}

func NewGromTokenRepository(db *gorm.DB)*GormTokenRepository{
	return &GormTokenRepository{
		db: db,
	}
}

func (r *GormTokenRepository)Create(token *model.Token)error{
	return r.db.Create(token).Error
}

func (r *GormTokenRepository)FindByID(id uint)(*model.Token,error){
	var token model.Token
	if err:=r.db.First(&token,id).Error;err !=nil{
		return nil,err
	}
	return &token,nil
}

func (r *GormTokenRepository)FindByUserId(id uint)(*model.Token,error){
	var token model.Token
	result := r.db.Where("user_id = ?",id).First(&token)
	if result.Error != nil{
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil
        }
		  return nil,result.Error
	}
	return &token,nil
}

func (r *GormTokenRepository)Delete(id uint)error{
	return r.db.Delete(&model.Token{},id).Error
}