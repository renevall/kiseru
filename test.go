package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/renevall/kiseru/model"
	"github.com/shopspring/decimal"
)

//Test create a test setup
func Test(db *gorm.DB) {

}

func InitTestData(db *gorm.DB) {

	employer := &model.Employer{
		UUID:    uuid.New(),
		Name:    "Test Employer",
		Address: "Test Address",
		Phone:   "18002222",
		Email:   "test@testing.com",
	}

	db.Create(employer)
	if db.NewRecord(employer) {
		fmt.Println("Could not create employer")
	}

	area := &model.Area{
		Name:       "Test Area",
		EmployerID: employer.ID,
	}

	db.Create(area)
	if db.NewRecord(area) {
		fmt.Println("Could not create area")
	}

	position := &model.Position{
		Name: "Test Position",
		Area: area.Name,
	}

	db.Create(position)
	if db.NewRecord(position) {
		fmt.Println("Could not create position")
	}

	rates := generateRates()

	for _, rate := range rates {
		db.Create(&rate)
		if db.NewRecord(rate) {
			fmt.Println("Could not create new Rate")
		}
	}

	employees := generateEmployees(1, 1)
	for _, e := range employees {
		db.Create(&e)
		if db.NewRecord(e) {
			fmt.Println("Could not create new Rate")
		}
	}

	data, _ := json.MarshalIndent(employees, "", " ")
	os.Stdout.Write(data)

}

func generateRates() []model.IncomeRate {

	rates := []model.IncomeRate{
		{
			Alias:      "Estrato 1",
			Min:        decimal.NewFromFloat(0.1),       //0.1
			Max:        decimal.NewFromFloat(100000.00), //100,0000.00
			Base:       decimal.NewFromFloat(0),         //0
			Percentage: decimal.NewFromFloat(0),         //0%
			Excess:     decimal.NewFromFloat(0),         //0
		},
		{
			Alias:      "Estrato 2",
			Min:        decimal.NewFromFloat(100000.01), //100,000.01
			Max:        decimal.NewFromFloat(200000.00), //200.000.00
			Base:       decimal.NewFromFloat(0),         //0
			Percentage: decimal.NewFromFloat(0.15),      //15%
			Excess:     decimal.NewFromFloat(100000.00), //100,000.00
		},
		{
			Alias:      "Estrato 3",
			Min:        decimal.NewFromFloat(200000.01), //200,000.01
			Max:        decimal.NewFromFloat(350000.00), //350,000,00
			Base:       decimal.NewFromFloat(45000.00),  //45,000.00
			Percentage: decimal.NewFromFloat(0.20),      //20%
			Excess:     decimal.NewFromFloat(200000.00), //200,000.00
		},
		{
			Alias:      "Estrato 4",
			Min:        decimal.NewFromFloat(350000.01), //350,000.01
			Max:        decimal.NewFromFloat(500000.00), //500,0000.00
			Base:       decimal.NewFromFloat(45000.00),  //45,000.00
			Percentage: decimal.NewFromFloat(0.25),      //25%
			Excess:     decimal.NewFromFloat(350000.00), //350,000.00
		},
		{
			Alias:      "Estrato 5",
			Min:        decimal.NewFromFloat(500000.01), //500,000.01
			Max:        decimal.NewFromFloat(100000.00), //10,000,0000.00
			Base:       decimal.NewFromFloat(82500.00),  //82,500.00
			Percentage: decimal.NewFromFloat(0.30),      //30%
			Excess:     decimal.NewFromFloat(500000.00), //500,000.00
		},
	}
	return rates
}

func generateEmployees(e, a uint) []model.Employee {
	values := []float64{85000, 135000, 4000000, 225000, 630000}

	var ems []model.Employee
	month := decimal.NewFromFloat(12.0)
	for x := 1; x <= 20; x++ {
		random := rand.Intn(4)
		salary := decimal.NewFromFloat(values[random])
		hr := salary.DivRound(month, 4)
		employee := model.Employee{
			Firstname:  "Name" + strconv.Itoa(x),
			Lastname:   "Lastname" + strconv.Itoa(x),
			UUID:       uuid.New(),
			Salary:     salary,
			HourlyRate: hr,
			Email:      "Email" + strconv.Itoa(x),
			EmployerID: e,
			AreaID:     a,
		}
		ems = append(ems, employee)
	}
	return ems
}
