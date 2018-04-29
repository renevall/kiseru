package model

import (
	"github.com/jinzhu/gorm"
)

//Position holds the positions the employee has
type Position struct {
	gorm.Model
	Name string `gorm:"name"`
	Area string `gorm:"area"`
}
