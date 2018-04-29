package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//Employer holds the data of a employeer
type Employer struct {
	gorm.Model
	UUID    uuid.UUID `json:"uuid" gorm:"column:uuid"`
	Name    string    `json:"name" gorm:"column:name"`
	Address string    `json:"address" gorm:"column:address"`
	Phone   string    `json:"phone" gorm:"column:phone"`
	Email   string    `json:"email" gorm:"column:email"`
}
