package model

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name string `json:"name"`
}