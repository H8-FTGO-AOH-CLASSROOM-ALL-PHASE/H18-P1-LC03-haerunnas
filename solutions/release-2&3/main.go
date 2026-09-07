package main

import (
	"fmt"
	"healthtrack/database"
	"log"

	"github.com/hrnnsx/go-toolkit/stdio"
)

type Menu struct {
	id          int
	description string
}

var listMenu = []Menu{
	{id: 1, description: "Register a patient"},
	{id: 2, description: "Add treatment"},
	{id: 3, description: "Patient reports"},
	{id: 4, description: "Doctor activity"},
	{id: 5, description: "Billing overview"},
	{id: 6, description: "Appointment stats"},
	{id: 7, description: "Exit"},
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	defer db.Close()

	for {
		fmt.Print("\n--- Patient Management Interface ---\n\n")
		for _, item := range listMenu {
			fmt.Printf("%d. %s\n", item.id, item.description)
		}

		input := stdio.Prompt("Masukkan pilihanmu: ").AsString()

		switch input {
		case "1":
			registration(db)
		case "2":
			addTreatment(db)
		case "3":
			patientReport(db)
		case "4":
			doctorActivity(db)
		case "5":
			billOverview(db)
		case "6":
			appointmentStats(db)
		case "7":
			return
		}
	}
}
