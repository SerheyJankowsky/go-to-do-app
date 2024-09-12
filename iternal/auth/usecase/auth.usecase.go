package usecase

import (
	"to-do-app/iternal/auth/delivery/http/dto"
	"to-do-app/iternal/auth/domain/model"
	userModel "to-do-app/iternal/user/domain/model"
)

type AuthUseCase interface{
	 Login(*dto.LoginDto)(*userModel.User,*model.Token,error)
    Register(*dto.RegisterDto)(*userModel.User,*model.Token,error)
	 Me(string)(*userModel.User,error)
	 Refresh(*userModel.User,string)(string,error)
}