package entity

import (
	"mygram/helper"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the entity for a user
type User struct {
	Id          uint          `json:"id" gorm:"primaryKey"`
	Username    string        `json:"username" gorm:"not null;type:varchar(20);uniqueIndex" form:"username" valid:"required~Username can not be empty"`
	Email       string        `json:"email" gorm:"not null;type:varchar(50);uniqueIndex" form:"email" valid:"required~Email can not be empty,email~Email format must be 'example@email.com'"`
	Password    string        `json:"password" gorm:"not null;type:varchar(100)" form:"password" valid:"required~Password can not be empty,minstringlength(6)~Password must have a minimum length of 6 characters"`
	Age         int           `json:"age" gorm:"not null;type:int" form:"age" valid:"required~Age can not be empty,range(8|100)~The age must be 8 years older"`
	Photos      []Photo       `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedia []SocialMedia `json:"social_media" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Commets     []Comment     `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidation := govalidator.ValidateStruct(user)
	if errValidation != nil {
		err = errValidation
		return
	}

	user.Password = helper.HashPassword(user.Password)
	err = nil
	return
}
