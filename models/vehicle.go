package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model // use Embedded Struct from gorm. See https://gorm.io/docs/models.html 
	Name  string `json:"nome"`
	Vin   string `json:"cpf"`
	Brand string `json:"rg"`
}

var Vehicles []Vehicle