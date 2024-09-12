package http

import (
    "errors"
    "strings"
    "to-do-app/iternal/auth/delivery/http/dto"
    "to-do-app/iternal/auth/usecase"

    "github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
    useCase usecase.AuthUseCase
}

func NewUserHandler(useCase usecase.AuthUseCase) *AuthHandler {
    return &AuthHandler{
        useCase: useCase,
    }
}

func (u *AuthHandler) Register(c *fiber.Ctx) error {
    var dto dto.RegisterDto
    if err := c.BodyParser(&dto); err != nil {
        return err
    }
    user, token, err := u.useCase.Register(&dto)
    if err != nil {
        return err
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "user":  user,
        "token": token.Hash,
    })
}

func (u *AuthHandler) Login(c *fiber.Ctx) error {
    var dto dto.LoginDto
    if err := c.BodyParser(&dto); err != nil {
        return err
    }
    user, token, err := u.useCase.Login(&dto)
    if err != nil {
        return err
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "user":  user,
        "token": token.Hash,
    })
}

func (u *AuthHandler) Me(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    token, err := GetAuthHeader(authHeader)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON("token not found")
    }
    user, err := u.useCase.Me(token)
    if err != nil {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(user)
}

func (u *AuthHandler) Refresh(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    token, err := GetAuthHeader(authHeader)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON("token not found")
    }
    user, err := u.useCase.Me(token)
    if err != nil {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    newToken, err := u.useCase.Refresh(user, token)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "token": newToken,
        "user":  user,
    })
}

func GetAuthHeader(authHeader string) (string, error) {
    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        return "", errors.New("doesnt hase auth header")
    }
    token := strings.TrimPrefix(authHeader, "Bearer ")
    return token, nil
}
