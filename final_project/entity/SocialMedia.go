package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// SocialMedia represents the entity for a social media
type SocialMedia struct {
	Id             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null;varchar(20)" valid:"required~Name can not be empty"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"not null;varchar(100)" valid:"required~Social media URL can not be empty"`
	UserId         uint      `json:"user_id"`
	User           *User     `json:"user"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

func (socialMedia *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValidation := govalidator.ValidateStruct(socialMedia)
	if errValidation != nil {
		err = errValidation
		return
	}

	err = nil
	return
}
