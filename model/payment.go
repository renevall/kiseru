package model

import (
	"github.com/ericlagergren/decimal"

	"github.com/jinzhu/gorm"
)

//Payment is used to record each payment the employer has issued
type Payment struct {
	gorm.Model
	EmployeeID  uint64      `json:"employee_id" gorm:"column:employee_id"`
	EmployeerID uint64      `json:"employeer_id" gorm:"column:employeer_id"`
	GrossAmount decimal.Big `json:"gross_ammount" gorm:"column:gross_ammount;type:decimal(10,4)"`
	NetAmount   decimal.Big `json:"net_ammount" gorm:"column:net_ammount;type:decimal(10,4)"`
}
