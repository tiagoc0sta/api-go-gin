package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model // use Embedded Struct from gorm. See https://gorm.io/docs/models.html 
	Name  string `json:"name" validate:"nonzero"`
	Vin   string `json:"vin" validate:"len=17"`  
	Brand string `json:"brand" validate:"nonzero"`
}


func ValidateVehicleData(vehicle *Vehicle) error {
	if err := validator.Validate(vehicle); err != nil {
		return err
	}
	return nil
}
