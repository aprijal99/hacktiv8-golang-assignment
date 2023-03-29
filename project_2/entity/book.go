package entity

import "time"

type Book struct {
	Id        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null;unique;type:varchar(200)"`
	Author    string `gorm:"not null;unique;type:varchar(200)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
