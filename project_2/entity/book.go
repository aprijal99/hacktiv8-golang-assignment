package entity

import "time"

type Book struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null;unique;type:varchar(200)"`
	Author    string    `json:"author" gorm:"not null;unique;type:varchar(200)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
