package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/renevall/kiseru/model"
)

//InitDB initializes the database
func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&model.Employer{},
		&model.Area{},
		&model.Position{},
		&model.Employee{},
		&model.IncomeRate{},
		&model.Payment{},
	)

	return db
}
