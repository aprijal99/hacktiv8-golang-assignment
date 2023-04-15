package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the entity for a comment
type Comment struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	User      *User     `json:"user"`
	PhotoId   uint      `json:"photo_id" valid:"required~Photo ID can not be empty"`
	Photo     *Photo    `json:"photo"`
	Message   string    `json:"message" gorm:"not null;varchar(255)" valid:"required~Message can not be empty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidation := govalidator.ValidateStruct(comment)
	if errValidation != nil {
		err = errValidation
		return
	}

	err = nil
	return
}
