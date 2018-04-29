package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

//IncomeRate holds the rate which is going to be used to
//calculate salaries income tax
type IncomeRate struct {
	gorm.Model
	Alias      string          `json:"alias" gorm:"column:alias"`
	Min        decimal.Decimal `json:"min" gorm:"column:min;type:decimal(10,4)"`
	Max        decimal.Decimal `json:"max" gorm:"column:max;type:decimal(10,4)"`
	Base       decimal.Decimal `json:"base" gorm:"column:base;type:decimal(10,4)"`
	Percentage decimal.Decimal `json:"percentage" gorm:"column:percentage;type:decimal(10,4)"`
	Excess     decimal.Decimal `json:"excess" gorm:"column:excess;type:decimal(10,4)"`
}
