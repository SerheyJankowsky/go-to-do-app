package model

type Item struct {
    ID   uint   `gorm:"primaryKey"`
    Title string
	Description string
    UserId uint	
}