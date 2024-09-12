package main

import (
    "fmt"
    "log"
    "to-do-app/config"
    tokenModel "to-do-app/iternal/auth/domain/model"
    itemModel "to-do-app/iternal/item/domain/model"
    userModel "to-do-app/iternal/user/domain/model"
    fiberRouter "to-do-app/pkg/fiber"

    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Config.DBHost, config.Config.DBUser, config.Config.DBPassword, config.Config.DBName, config.Config.DBPort)
    app := fiber.New()
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Fatal to connect db")
    }
    if err := db.AutoMigrate(&userModel.User{}, &itemModel.Item{}, &tokenModel.Token{}); err != nil {

    }
    api := app.Group("api/v1")
    fiberRouter.SetupRoutes(api, db)
    log.Fatal(app.Listen(":3000"))
}
