package helper

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// import "github.com/asaskevich/govalidator"

func GetValidationErrorMessage(err error) interface{} {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return "Username or Email has been taken"
	}

	var errs []string
	for _, e := range err.(govalidator.Errors).Errors() {
		errs = append(errs, e.Error())
	}

	return errs
}
