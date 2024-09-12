package model

import (
	"time"
	"to-do-app/iternal/item/domain/model"
)

type User struct {
    ID   uint   `gorm:"primaryKey" json:"id"`
    Name string `json:"name"`
    Email string `gorm:"uniqueIndex;size=255" json:"email"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
	Password string `gorm:"size:255" json:"-"`
    Items []model.Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
}