package model

import (
	"github.com/jinzhu/gorm"
)

//Area represents a division where the user works in
type Area struct {
	gorm.Model
	Name       string `json:"name" gorm:"column:name"`
	EmployerID uint   `json:"employer_id" gorm:"column:employer_id"`
}
