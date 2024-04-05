package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model // use Embedded Struct from gorm. See https://gorm.io/docs/models.html 
	Name  string `json:"name"`
	Vin   string `json:"vin"`
	Brand string `json:"brand"`
}

