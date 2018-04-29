package model

import (
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//Employee holds the data of a employee
type Employee struct {
	gorm.Model
	Firstname  string       `json:"first_name" gorm:"column:first_name"`
	Lastname   string       `json:"last_name" gorm:"column:last_name"`
	UUID       uuid.UUID    `json:"uuid" gorm:"column:uuid"`
	StartDate  time.Time    `json:"start_date" gorm:"column:start_date"`
	EndDate    *time.Time   `json:"end_date" gorm:"column:end_date"`
	Salary     *decimal.Big `json:"salary" gorm:"column:salary"`
	HourlyRate *decimal.Big `json:"hourly_rate" gorm:"column:hourly_rate"`
	Email      string       `json:"email" gorm:"column:email"`
	AreaID     uint         `json:"area_id" gorm:"column:area_id"`
	EmployerID uint         `json:"employer_id" gorm:"column:employer_id"`
}
