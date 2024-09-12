package middleware

import (
	"errors"
	authHttp "to-do-app/iternal/auth/delivery/http"
	"to-do-app/iternal/auth/infrastructure/persistence"
	authUseCase "to-do-app/iternal/auth/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthMiddleware struct{
}

func NewMiddleware()*AuthMiddleware{
	return &AuthMiddleware{}
}

func(m *AuthMiddleware)Auth(db *gorm.DB)fiber.Handler{
	repo:=persistence.NewGromTokenRepository(db)
	return func(c *fiber.Ctx) error {
		authHeader:=c.Get("Authorization")
		token,err:=authHttp.GetAuthHeader(authHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON("token not found")
		}
		if err:=authUseCase.ValidateToken(token);err!=nil{
			return  c.Status(fiber.StatusUnauthorized).JSON("invalid token")
		}
		decodeToken,err:=authUseCase.DecodeToken(token)
		if err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":err.Error(),
			})
		}
		if err:= CompareToken(repo,token,uint(decodeToken.UserID));err!=nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":err.Error(),
			})
		}
		{
			c.Locals("user_id",decodeToken.UserID)
			c.Locals("user_email",decodeToken.Email)
		}
		return c.Next()
	}
}

func CompareToken(tokenRepo *persistence.GormTokenRepository,tokenString string,id uint)error{
	existToken,err:= tokenRepo.FindByUserId(id)
	if err!=nil{
		return errors.New("previos token not found")
	}
	if existToken.Hash!=tokenString{
		return errors.New("invalid token")
	}
	return nil
}