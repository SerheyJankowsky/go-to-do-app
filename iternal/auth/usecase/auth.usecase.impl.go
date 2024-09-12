package usecase

import (
	"errors"
	"fmt"
	"time"
	"to-do-app/config"
	"to-do-app/iternal/auth/delivery/http/dto"
	"to-do-app/iternal/auth/domain/model"
	tokenRepo "to-do-app/iternal/auth/domain/repository"
	userModel "to-do-app/iternal/user/domain/model"
	userRepo "to-do-app/iternal/user/domain/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaimes struct {
	 jwt.RegisteredClaims
    UserID uint   `json:"user_id"`
    Email  string `json:"user_email"`
	//  Exp *jwt.NumericDate `josn:"exp"`
}

type AuthUseCaseImpl struct{
	tokenRepo tokenRepo.TokenRepository
	userRepo userRepo.UserRepository
	jwtSecret string
	exp time.Time
}

func NewAuthUseCase(t tokenRepo.TokenRepository,u userRepo.UserRepository,jwtSecret string)*AuthUseCaseImpl{
	return &AuthUseCaseImpl{
		tokenRepo: t,
		userRepo: u,
		jwtSecret: jwtSecret,
		exp: time.Now().Add(time.Hour * 72),
	}
}

func (u *AuthUseCaseImpl)Login(dto *dto.LoginDto)(*userModel.User,*model.Token,error){
	user,err := u.userRepo.FindByEmail(dto.Email)
	if err!=nil{
		return nil,nil,errors.New("invalid credentials")
	}
	if err := u.comparePassword([]byte(user.Password),[]byte(dto.Password));err!=nil{
		return nil,nil,errors.New("invalid credentials")
	}
	existToken,_:=u.tokenRepo.FindByUserId(uint(user.ID))
	if existToken!=nil{
		u.tokenRepo.Delete(uint(existToken.ID))
	}
	jwtString,err:=u.generateToken(user.ID,user.Email)
	if err!=nil{
		return nil,nil,errors.New("create jwt error")
	}
	token := model.Token{
		Hash: jwtString,
		ExpAt: u.exp.Unix(),
		UserId: user.ID,
	}
	if err:=u.tokenRepo.Create(&token);err!=nil{
		fmt.Println(err)
		return nil,nil,errors.New("create token error")
	}
	return user,&token,nil
}

func (u *AuthUseCaseImpl)Register(dto *dto.RegisterDto)(*userModel.User,*model.Token,error){
	hashPassword,err := u.hashPassword([]byte(dto.Password))
	if err != nil {
		return  nil,nil,err
	}
	user:= &userModel.User{
		Name: dto.Name,
		Email: dto.Email,
		Password: hashPassword,
	}
	if err:=u.userRepo.Create(user);err!=nil{
		return nil,nil,err
	}
	jwtToken,err := u.generateToken(user.ID,user.Email)
	if err!=nil{
		return nil,nil,err
	}
	 token := &model.Token{
		Hash: jwtToken,
		ExpAt: u.exp.Unix(),
		UserId: user.ID,
	}
	if err:=u.tokenRepo.Create(token);err!=nil{
		return nil,nil,err
	}
	return user,token,nil
}

func(u *AuthUseCaseImpl)Me(tokenStirng string)(*userModel.User,error){
	token,err:=DecodeToken(tokenStirng)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user,err:=u.userRepo.FindByID(uint(token.UserID))
	if err!=nil{
		return nil,err
	}
	if _,err:=u.compareToken(tokenStirng,uint(user.ID));err!=nil{
		return nil,err
	}
	return user,nil
}

func (u *AuthUseCaseImpl)Refresh(user *userModel.User,tokenStirng string)(string,error){
	if err:=ValidateToken(tokenStirng);err!=nil{
		return "",err
	}
	existToken,err :=u.compareToken(tokenStirng,uint(user.ID))
	if err!=nil{
		return "",err
	}
	if existToken!=nil{
		u.tokenRepo.Delete(uint(existToken.ID))
	}
	jwtString,err:=u.generateToken(user.ID,user.Email)
	if err!=nil{
		return "",err
	}
	token:=&model.Token{
		Hash: jwtString,
		ExpAt: u.exp.Unix(),
		UserId: user.ID,
	}
	u.tokenRepo.Create(token)
	return jwtString,nil
}

func(u *AuthUseCaseImpl)compareToken(tokenString string,id uint)(*model.Token,error){
	existToken,err:=u.tokenRepo.FindByUserId(id)
	if err!=nil{
		return nil,errors.New("previos token not found")
	}
	if existToken.Hash!=tokenString{
		return nil,errors.New("invalid token")
	}
	return existToken,nil
}

func (u *AuthUseCaseImpl)hashPassword(password []byte)(string,error){
	hashPassword,err :=bcrypt.GenerateFromPassword(password,bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword),nil
}

func (u *AuthUseCaseImpl) comparePassword(password []byte, hashPassword []byte)error{
	err:= bcrypt.CompareHashAndPassword(password,hashPassword)
	return err
}

func (u *AuthUseCaseImpl) generateToken(id uint,email string)(string,error){
	claimes:=CustomClaimes{
		// Exp:jwt.NewNumericDate(u.exp),
		UserID:id,
		Email:email,
		RegisteredClaims: jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(u.exp),
		  IssuedAt:  jwt.NewNumericDate(time.Now()),
		  NotBefore: jwt.NewNumericDate(time.Now()),
		  Issuer:    "to-do-app",
		  Subject:   "user-id, email",
    },
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claimes)
	return token.SignedString([]byte(u.jwtSecret))
}

func ValidateToken(tokenString string)(error){
	secret:=[]byte(config.Config.JWTSecret)
	token,err:= jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }
        return secret, nil
	})
	if err!=nil{
		return errors.New("parsing error")
	}
	if !token.Valid{
		return errors.New("token invalid")
	}
	return nil
}

func DecodeToken(tokenString string)(*CustomClaimes,error){
	claimes := CustomClaimes{}
	token, err := jwt.ParseWithClaims(tokenString,&claimes,func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JWTSecret),nil
	})
	if err != nil {
		return  nil,err
	}
	if claims, ok := token.Claims.(*CustomClaimes); ok && token.Valid {
    	return claims,nil
	} 
	return nil,errors.New("token error")
}