package fiber

import (
	"to-do-app/config"
	"to-do-app/iternal/middleware"
	userHttp "to-do-app/iternal/user/delivery/http"
	userPersistence "to-do-app/iternal/user/infrastructure/persistence"
	userUseCase "to-do-app/iternal/user/usecase"

	itemHttp "to-do-app/iternal/item/delivery/http"
	itemPersistence "to-do-app/iternal/item/infrastructure/persistence"
	itemUseCase "to-do-app/iternal/item/usecase"

	authHttp "to-do-app/iternal/auth/delivery/http"
	authPersistence "to-do-app/iternal/auth/infrastructure/persistence"
	authUseCase "to-do-app/iternal/auth/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var middlewares middleware.AuthMiddleware

func SetupRoutes(api fiber.Router, db *gorm.DB) {
    userRoutes(api, db)
    authRoutes(api, db)
    itemRoutes(api, db)
}

func authRoutes(api fiber.Router, db *gorm.DB) {
    userRepo := userPersistence.NewGromUserRepository(db)
    tokenRepo := authPersistence.NewGromTokenRepository(db)
    authUseCase := authUseCase.NewAuthUseCase(tokenRepo, userRepo, config.Config.JWTSecret)
    authHandler := authHttp.NewUserHandler(authUseCase)
    auth := api.Group("/auth")
    {
        auth.Post("/login", authHandler.Login)
        auth.Post("/register", authHandler.Register)
        auth.Get("/me", authHandler.Me)
        auth.Get("/refresh", authHandler.Refresh)
    }
}

func userRoutes(api fiber.Router, db *gorm.DB) {
    userRepo := userPersistence.NewGromUserRepository(db)
    userUseCase := userUseCase.NewUserUseCaseImpl(userRepo)
    userHandler := userHttp.NewUserHandler(userUseCase)
    user := api.Group("/user", middlewares.Auth(db))
    {
        user.Post("/", userHandler.CreateUser)
        user.Get("/:id", userHandler.GetUser)
        user.Patch("/:id", userHandler.UpdateUser)
        user.Delete("/:id", userHandler.DeleteUser)
    }
}

func itemRoutes(api fiber.Router, db *gorm.DB) {
    itemRepo := itemPersistence.NewGromUserRepository(db)
    itemUseCase := itemUseCase.NewItemUseCase(itemRepo)
    itemHandler := itemHttp.NewItemHandler(itemUseCase)
    item := api.Group("/item", middlewares.Auth(db))
    {
        item.Post("/", itemHandler.CreateItem)
        item.Get("/",itemHandler.GetUserItems)
        item.Patch("/:id",itemHandler.UpdateItem)
        item.Delete("/:id",itemHandler.DeleteItem)
    }
}
