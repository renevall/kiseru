package model

import (
	"github.com/ericlagergren/decimal"
	"github.com/jinzhu/gorm"
)

//RateTable holds the rate which is going to be used to
//calculate salaries income tax
type IncomeRate struct {
	gorm.Model
	Alias      string       `json:"alias" gorm:"column:alias"`
	Min        *decimal.Big `json:"min" gorm:"column:min;type:decimal(10,4)"`
	Max        *decimal.Big `json:"max" gorm:"column:max;type:decimal(10,4)"`
	Base       *decimal.Big `json:"base" gorm:"column:base;type:decimal(10,4)"`
	Percentage *decimal.Big `json:"percentage" gorm:"column:percentage;type:decimal(10,4)"`
	Excess     *decimal.Big `json:"excess" gorm:"column:excess;type:decimal(10,4)"`
}
