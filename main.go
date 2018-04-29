package main

func main() {
	db := InitDB()
	defer db.Close()

	InitTestData(db)
	// rates := generateRates()
	// data, _ := json.MarshalIndent(rates, "", " ")
	// os.Stdout.Write(data)

	// employees := generateEmployees(1, 2)
	// data, _ := json.MarshalIndent(employees, "", " ")
	// os.Stdout.Write(data)
}
