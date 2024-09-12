package http

import (
	"strconv"
	"to-do-app/iternal/user/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler{
	return &UserHandler{
		useCase: useCase,
	}
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error{
	var req struct{
		Name string `json:"name"`
		Password string `json:"password"`
	}
	if err:=c.BodyParser(&req);err!=nil{
		return err
	}
	user,err := u.useCase.CreateUser(req.Name,req.Password)

	if err!=nil{
		return err
	}

	return c.JSON(user)
}

func (u *UserHandler) GetUser(c *fiber.Ctx) error{
	id,_:=strconv.ParseInt(c.Params("id"),10,32)
	user,err:= u.useCase.GetUser(uint(id))
	if err!=nil{
		return err
	}
	return c.JSON(user)
}

func (u *UserHandler) UpdateUser(c *fiber.Ctx) error{
	id,_:= strconv.ParseInt(c.Params("id"),10,32)
	var req struct{
		Name string `json:"name"`
	}
	if err := c.BodyParser(&req);err != nil{
		return err
	}
	user,err := u.useCase.UpdateUser(uint(id),req.Name)
	if err !=nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserHandler) DeleteUser(c *fiber.Ctx) error{
	id,_ := strconv.ParseInt(c.Params("id"),10,32)
	if err:=u.useCase.DeleteUser(uint(id));err!=nil{
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}