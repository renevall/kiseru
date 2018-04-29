package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

//Payment is used to record each payment the employer has issued
type Payment struct {
	gorm.Model
	EmployeeID  uint64          `json:"employee_id" gorm:"column:employee_id"`
	EmployeerID uint64          `json:"employeer_id" gorm:"column:employeer_id"`
	GrossAmount decimal.Decimal `json:"gross_ammount" gorm:"column:gross_ammount;type:decimal(10,4)"`
	NetAmount   decimal.Decimal `json:"net_ammount" gorm:"column:net_ammount;type:decimal(10,4)"`
}
