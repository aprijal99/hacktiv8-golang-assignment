package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null;varchar(255)" valid:"required~Title can not be empty"`
	Caption   string    `json:"caption" gorm:"varchar(255)"`
	PhotoUrl  string    `json:"photo_url" gorm:"not null;varchar(255)" valid:"required~Photo Url can not be empty"`
	UserId    uint      `json:"user_id"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidation := govalidator.ValidateStruct(photo)
	if errValidation != nil {
		err = errValidation
		return
	}

	err = nil
	return
}
