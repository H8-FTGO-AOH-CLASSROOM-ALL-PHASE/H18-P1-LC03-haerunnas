package main

import (
	"database/sql"
	"fmt"

	"github.com/hrnnsx/go-toolkit/stdio"
)

func registration(db *sql.DB) {

	name := stdio.Prompt("Enter patient name: ").AsString()
	birthDate := stdio.Prompt("Enter patient birth date (YYYY-MM-DD): ").AsString()
	medicalHistory := stdio.Prompt("Enter patient medical history: ").AsString()

	query := `
	INSERT INTO Patients(Name, DateOfBirth, MedicalHistory) VALUES (?, ?, ?)
	`
	_, err := db.Exec(query, name, birthDate, medicalHistory)
	if err != nil {
		fmt.Println("Failed to register patient: ", err)
		return
	}

	fmt.Println("Patient registered successfully")
}

func addTreatment(db *sql.DB) {
	patientId := stdio.Prompt("Enter patient ID: ").AsString()
	doctorId := stdio.Prompt("Enter doctor ID: ").AsString()

	diagnosis := stdio.Prompt("Enter Diagnosis: ").AsString()
	medication := stdio.Prompt("Enter Treatment: ").AsString()

	if patientId == "" || doctorId == "" || diagnosis == "" || medication == "" {
		fmt.Println("Error: Semua harus data diisi.")
		return
	}

	var exist bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM Patients WHERE PatientId = ? )", patientId).Scan(&exist)
	if err != nil || !exist {
		fmt.Printf("Error: Gak ada patient dengan id %s!", patientId)
		return
	}
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM Doctors WHERE DoctorId = ? )", patientId).Scan(&exist)
	if err != nil || !exist {
		fmt.Printf("Error: Gak ada patient dengan id %s!", patientId)
		return
	}

	query := `INSERT INTO Treatments (PatientId, DoctorId, Diagnosis, Medication) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(query, patientId, doctorId, diagnosis, medication)
	if err != nil {
		fmt.Println("Failed to register patient: ", err)
		return
	}

	fmt.Println("Treatment added successfully")
}

func patientReport(db *sql.DB) {
	query := `
	SELECT Patientid, Name, DateOfBirth, MedicalHistory FROM Patients
	`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error executing query: ", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var patientId int
		var Name string
		var DateOfBirth string
		var MedicalHistory string

		err = rows.Scan(&patientId, &Name, &DateOfBirth, &MedicalHistory)
		if err != nil {
			fmt.Printf("Error cuy: %v\n", err)
			return
		}

		fmt.Printf("\nID: %d\n", patientId)
		fmt.Printf("Name: %s\n", Name)
		fmt.Printf("Date of Birth: %s\n", DateOfBirth)
		fmt.Printf("Medical History: %s\n", MedicalHistory)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error reading rows: ", err)
	}
}

func doctorActivity(db *sql.DB) {
	query := `
		SELECT
			d.Name,
			COUNT(DISTINCT a.AppointmentId) AS TotalAppointments,
			COUNT(DISTINCT t.TreatmentId) AS TotalTreatments
		FROM Doctors d
		LEFT JOIN Appointments a ON d.DoctorId = a.DoctorId
		LEFT JOIN Treatments t ON d.DoctorId = t.DoctorId
		GROUP BY d.DoctorId, d.Name
	`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error executing query: ", err)
		return
	}
	defer rows.Close()

	fmt.Print("\n-- Doctor Activity --\n")
	for rows.Next() {
		var doctorName string
		var appointment int
		var treatment int

		err = rows.Scan(&doctorName, &appointment, &treatment)
		if err != nil {
			fmt.Printf("Error cuy: %v\n", err)
			return
		}

		fmt.Printf("\nDoctor: %s\n", doctorName)
		fmt.Printf("Appointments: %d\n", appointment)
		fmt.Printf("Treatments: %d\n", treatment)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error reading rows: ", err)
	}
}
func billOverview(db *sql.DB) {
	query := `
		SELECT 
			COALESCE(SUM(Amount), 0.00) AS TotalRevenue,
			0.00 AS OutstandingBills,
			COALESCE(SUM(Amount), 0.00) AS DailyIncome
		FROM Billing
	`

	var totalRevenue float64
	var outstandingBills float64
	var dailyIncome float64

	err := db.QueryRow(query).Scan(&totalRevenue, &outstandingBills, &dailyIncome)
	if err != nil {
		fmt.Println("Error executing query: ", err)
		return
	}

	fmt.Println("\n--- Billing Overview ---")
	fmt.Printf("Total Revenue: %.2f\n", totalRevenue)
	fmt.Printf("Outstanding Bills: %.2f\n", outstandingBills)
	fmt.Printf("Daily Income: %.2f\n", dailyIncome)
}

func appointmentStats(db *sql.DB) {
	query := `
		SELECT 
			COUNT(CASE WHEN t.TreatmentId IS NOT NULL THEN 1 END) AS Completed,
			COUNT(CASE WHEN t.TreatmentId IS NULL AND a.Date >= CURDATE() THEN 1 END) AS Scheduled,
			COUNT(CASE WHEN t.TreatmentId IS NULL AND a.Date < CURDATE() THEN 1 END) AS Missed
		FROM Appointments a
		LEFT JOIN Treatments t ON a.PatientId = t.PatientId AND a.DoctorId = t.DoctorId
	`

	var completed int
	var scheduled int
	var missed int

	err := db.QueryRow(query).Scan(&completed, &scheduled, &missed)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}

	fmt.Println("\n--- Appointment Stats ---")
	fmt.Printf("Scheduled: %d\n", scheduled)
	fmt.Printf("Completed: %d\n", completed)
	fmt.Printf("Missed: %d\n", missed)
}
