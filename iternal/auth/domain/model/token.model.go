package model

import (
	"time"
)

type Token struct{
	ID uint `gorm:"primaryKey"`
	Hash string 
	ExpAt int64
	CreatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	UserId uint
}