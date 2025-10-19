// Medical Appointment Booking System - Test Suite
// Copyright (C) 2025
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"log"
	"time"

	"bookings/database"
	"bookings/models"
)

func stringPtr(s string) *string {
	return &s
}

func float64Ptr(f float64) *float64 {
	return &f
}

func testDB() {
	fmt.Println("=== Starting Database and API Tests ===")

	// Initialize database connection
	database.InitDB()
	defer database.CloseDB()

	fmt.Println("✅ Database connection initialized")

	// Create tables
	if err := database.CreateTables(); err != nil {
		log.Fatalf("❌ Failed to create tables: %v", err)
	}
	fmt.Println("✅ Database tables created successfully")

	// Test Clinic CRUD
	testClinicCRUD()

	// Test Patient CRUD
	testPatientCRUD()

	// Test Employee CRUD
	testEmployeeCRUD()

	// Test Service CRUD
	testServiceCRUD()

	// Test Appointment CRUD
	testAppointmentCRUD()

	// Test Waiting List CRUD
	testWaitingListCRUD()

	fmt.Println("=== All Tests Completed Successfully! ===")
}

func testClinicCRUD() {
	fmt.Println("\n--- Testing Clinic CRUD ---")

	// Create clinic
	clinic := &models.Clinic{
		Name:    "Test Clinic",
		Address: "123 Test Street",
		Phone:   "+1234567890",
		Email:   "test@clinic.com",
		Active:  true,
	}

	if err := database.CreateClinic(clinic); err != nil {
		log.Printf("❌ Failed to create clinic: %v", err)
		return
	}
	fmt.Printf("✅ Created clinic with ID: %d\n", clinic.ID)

	// Get clinic
	retrievedClinic, err := database.GetClinic(clinic.ID)
	if err != nil {
		log.Printf("❌ Failed to get clinic: %v", err)
		return
	}
	fmt.Printf("✅ Retrieved clinic: %s\n", retrievedClinic.Name)

	// Update clinic
	clinic.Address = "456 Updated Street"
	if err := database.UpdateClinic(clinic.ID, clinic); err != nil {
		log.Printf("❌ Failed to update clinic: %v", err)
		return
	}
	fmt.Println("✅ Updated clinic successfully")

	// Get all clinics
	clinics, err := database.GetClinics()
	if err != nil {
		log.Printf("❌ Failed to get clinics: %v", err)
		return
	}
	fmt.Printf("✅ Found %d clinics\n", len(clinics))

	// Delete clinic
	if err := database.DeleteClinic(clinic.ID); err != nil {
		log.Printf("❌ Failed to delete clinic: %v", err)
		return
	}
	fmt.Println("✅ Deleted clinic successfully")
}

func testPatientCRUD() {
	fmt.Println("\n--- Testing Patient CRUD ---")

	// Create patient
	dateOfBirth := "1990-01-01"
	insuranceProvider := "Test Insurance"
	insuranceID := "INS123456"
	emergencyName := "Jane Doe"
	emergencyPhone := "+0987654321"
	patient := &models.Patient{
		FirstName:             "John",
		LastName:              "Doe",
		Email:                 "john.doe@example.com",
		Phone:                 "+1234567890",
		DateOfBirth:           &dateOfBirth,
		MedicalRecordNumber:   "MRN123456",
		InsuranceProvider:     &insuranceProvider,
		InsuranceID:           &insuranceID,
		EmergencyContactName:  &emergencyName,
		EmergencyContactPhone: &emergencyPhone,
		Active:                true,
	}

	if err := database.CreatePatient(patient); err != nil {
		log.Printf("❌ Failed to create patient: %v", err)
		return
	}
	fmt.Printf("✅ Created patient with ID: %d\n", patient.ID)

	// Get patient
	retrievedPatient, err := database.GetPatient(patient.ID)
	if err != nil {
		log.Printf("❌ Failed to get patient: %v", err)
		return
	}
	fmt.Printf("✅ Retrieved patient: %s %s\n", retrievedPatient.FirstName, retrievedPatient.LastName)

	// Update patient
	patient.Phone = "+1111111111"
	if err := database.UpdatePatient(patient.ID, patient); err != nil {
		log.Printf("❌ Failed to update patient: %v", err)
		return
	}
	fmt.Println("✅ Updated patient successfully")

	// Get all patients
	patients, err := database.GetPatients()
	if err != nil {
		log.Printf("❌ Failed to get patients: %v", err)
		return
	}
	fmt.Printf("✅ Found %d patients\n", len(patients))

	// Delete patient
	if err := database.DeletePatient(patient.ID); err != nil {
		log.Printf("❌ Failed to delete patient: %v", err)
		return
	}
	fmt.Println("✅ Deleted patient successfully")
}

func testEmployeeCRUD() {
	fmt.Println("\n--- Testing Employee CRUD ---")

	// First create a clinic for the employee
	clinic := &models.Clinic{
		Name:    "Employee Test Clinic",
		Address: "789 Employee St",
		Phone:   "+1234567890",
		Email:   "employee@clinic.com",
		Active:  true,
	}
	if err := database.CreateClinic(clinic); err != nil {
		log.Printf("❌ Failed to create clinic for employee: %v", err)
		return
	}

	// Create employee
	employee := &models.Employee{
		ClinicID:      clinic.ID,
		FirstName:     "Dr. Jane",
		LastName:      "Smith",
		Email:         "jane.smith@clinic.com",
		Phone:         "+1234567890",
		LicenseNumber: "LIC123456",
		Specialty:     "Cardiology",
		Timezone:      "Asia/Colombo",
		Active:        true,
	}

	if err := database.CreateEmployee(employee); err != nil {
		log.Printf("❌ Failed to create employee: %v", err)
		return
	}
	fmt.Printf("✅ Created employee with ID: %d\n", employee.ID)

	// Get employee
	retrievedEmployee, err := database.GetEmployee(employee.ID)
	if err != nil {
		log.Printf("❌ Failed to get employee: %v", err)
		return
	}
	fmt.Printf("✅ Retrieved employee: %s %s\n", retrievedEmployee.FirstName, retrievedEmployee.LastName)

	// Update employee
	employee.Phone = "+2222222222"
	if err := database.UpdateEmployee(employee.ID, employee); err != nil {
		log.Printf("❌ Failed to update employee: %v", err)
		return
	}
	fmt.Println("✅ Updated employee successfully")

	// Get all employees
	employees, err := database.GetEmployees()
	if err != nil {
		log.Printf("❌ Failed to get employees: %v", err)
		return
	}
	fmt.Printf("✅ Found %d employees\n", len(employees))

	// Delete employee
	if err := database.DeleteEmployee(employee.ID); err != nil {
		log.Printf("❌ Failed to delete employee: %v", err)
		return
	}
	fmt.Println("✅ Deleted employee successfully")

	// Clean up clinic
	database.DeleteClinic(clinic.ID)
}

func testServiceCRUD() {
	fmt.Println("\n--- Testing Service CRUD ---")

	// Create service
	service := &models.Service{
		Name:              "General Consultation",
		Description:       "General medical consultation",
		DurationMinutes:   30,
		Price:             100.00,
		SpecialtyRequired: "General Medicine",
		Active:            true,
	}

	if err := database.CreateService(service); err != nil {
		log.Printf("❌ Failed to create service: %v", err)
		return
	}
	fmt.Printf("✅ Created service with ID: %d\n", service.ID)

	// Get service
	retrievedService, err := database.GetService(service.ID)
	if err != nil {
		log.Printf("❌ Failed to get service: %v", err)
		return
	}
	fmt.Printf("✅ Retrieved service: %s\n", retrievedService.Name)

	// Update service
	service.Price = 120.00
	if err := database.UpdateService(service.ID, service); err != nil {
		log.Printf("❌ Failed to update service: %v", err)
		return
	}
	fmt.Println("✅ Updated service successfully")

	// Get all services
	services, err := database.GetServices()
	if err != nil {
		log.Printf("❌ Failed to get services: %v", err)
		return
	}
	fmt.Printf("✅ Found %d services\n", len(services))

	// Delete service
	if err := database.DeleteService(service.ID); err != nil {
		log.Printf("❌ Failed to delete service: %v", err)
		return
	}
	fmt.Println("✅ Deleted service successfully")
}

func testAppointmentCRUD() {
	fmt.Println("\n--- Testing Appointment CRUD ---")

	// Create required entities first
	clinic := &models.Clinic{Name: "Appointment Clinic", Address: "123 Appt St", Phone: "+1234567890", Email: "appt@clinic.com", Active: true}
	database.CreateClinic(clinic)

	patient := &models.Patient{FirstName: "Test", LastName: "Patient", Email: "test@patient.com", Phone: "+1234567890", DateOfBirth: stringPtr("1990-01-01"), MedicalRecordNumber: "MRN999", Active: true}
	database.CreatePatient(patient)

	employee := &models.Employee{ClinicID: clinic.ID, FirstName: "Dr. Test", LastName: "Doctor", Email: "test@doctor.com", Phone: "+1234567890", LicenseNumber: "LIC999", Specialty: "General", Timezone: "Asia/Colombo", Active: true}
	database.CreateEmployee(employee)

	service := &models.Service{Name: "Test Service", Description: "Test service", DurationMinutes: 30, Price: 50.00, SpecialtyRequired: "General", Active: true}
	database.CreateService(service)

	// Create appointment
	startTime := time.Now().Add(24 * time.Hour).UTC() // Tomorrow
	endTime := startTime.Add(30 * time.Minute)

	appointment := &models.Appointment{
		PatientID:       patient.ID,
		EmployeeID:      employee.ID,
		ServiceID:       service.ID,
		ClinicID:        clinic.ID,
		StartDatetime:   startTime,
		EndDatetime:     endTime,
		Status:          "SCHEDULED",
		AppointmentType: stringPtr("INITIAL_CONSULTATION"),
		Notes:           stringPtr("Test appointment"),
		PaymentStatus:   "PENDING",
		PaymentAmount:   float64Ptr(50.00),
	}

	if err := database.CreateAppointment(appointment); err != nil {
		log.Printf("❌ Failed to create appointment: %v", err)
		return
	}
	fmt.Printf("✅ Created appointment with ID: %d\n", appointment.ID)

	// Get appointment
	retrievedAppointment, err := database.GetAppointment(appointment.ID)
	if err != nil {
		log.Printf("❌ Failed to get appointment: %v", err)
		return
	}
	fmt.Printf("✅ Retrieved appointment for patient ID: %d\n", retrievedAppointment.PatientID)

	// Update appointment
	appointment.Notes = stringPtr("Updated test appointment")
	if err := database.UpdateAppointment(appointment.ID, appointment); err != nil {
		log.Printf("❌ Failed to update appointment: %v", err)
		return
	}
	fmt.Println("✅ Updated appointment successfully")

	// Get all appointments
	appointments, err := database.GetAppointments()
	if err != nil {
		log.Printf("❌ Failed to get appointments: %v", err)
		return
	}
	fmt.Printf("✅ Found %d appointments\n", len(appointments))

	// Delete appointment
	if err := database.DeleteAppointment(appointment.ID); err != nil {
		log.Printf("❌ Failed to delete appointment: %v", err)
		return
	}
	fmt.Println("✅ Deleted appointment successfully")

	// Clean up
	database.DeleteService(service.ID)
	database.DeleteEmployee(employee.ID)
	database.DeletePatient(patient.ID)
	database.DeleteClinic(clinic.ID)
}

func testWaitingListCRUD() {
	fmt.Println("\n--- Testing Waiting List CRUD ---")

	// Create required entities
	clinic := &models.Clinic{Name: "Waiting Clinic", Address: "456 Wait St", Phone: "+1234567890", Email: "wait@clinic.com", Active: true}
	database.CreateClinic(clinic)

	patient := &models.Patient{FirstName: "Wait", LastName: "Patient", Email: "wait@patient.com", Phone: "+1234567890", DateOfBirth: stringPtr("1990-01-01"), MedicalRecordNumber: "MRN888", Active: true}
	database.CreatePatient(patient)

	service := &models.Service{Name: "Wait Service", Description: "Waiting service", DurationMinutes: 45, Price: 75.00, SpecialtyRequired: "General", Active: true}
	database.CreateService(service)

	// Create waiting list item
	waitingItem := &models.WaitingList{
		PatientID:     patient.ID,
		ServiceID:     service.ID,
		RequestedDate: stringPtr("2025-01-15"),
		UrgencyLevel:  "HIGH",
		Notes:         stringPtr("Urgent appointment needed"),
		Status:        "ACTIVE",
	}

	if err := database.CreateWaitingListItem(waitingItem); err != nil {
		log.Printf("❌ Failed to create waiting list item: %v", err)
		return
	}
	fmt.Printf("✅ Created waiting list item with ID: %d\n", waitingItem.ID)

	// Get waiting list item
	retrievedItem, err := database.GetWaitingListItem(waitingItem.ID)
	if err != nil {
		log.Printf("❌ Failed to get waiting list item: %v", err)
		return
	}
	fmt.Printf("✅ Retrieved waiting list item for patient ID: %d\n", retrievedItem.PatientID)

	// Update waiting list item
	waitingItem.Notes = stringPtr("Updated urgent notes")
	if err := database.UpdateWaitingListItem(waitingItem.ID, waitingItem); err != nil {
		log.Printf("❌ Failed to update waiting list item: %v", err)
		return
	}
	fmt.Println("✅ Updated waiting list item successfully")

	// Get all waiting list items
	waitingList, err := database.GetWaitingList()
	if err != nil {
		log.Printf("❌ Failed to get waiting list: %v", err)
		return
	}
	fmt.Printf("✅ Found %d items in waiting list\n", len(waitingList))

	// Delete waiting list item
	if err := database.DeleteWaitingListItem(waitingItem.ID); err != nil {
		log.Printf("❌ Failed to delete waiting list item: %v", err)
		return
	}
	fmt.Println("✅ Deleted waiting list item successfully")

	// Clean up
	database.DeleteService(service.ID)
	database.DeletePatient(patient.ID)
	database.DeleteClinic(clinic.ID)
}

// To run the tests, call testDB() from your main application or create a separate test binary
// You can build and run this file separately:
// go build -o test_db test_db.go
// ./test_db

// func main() {
// 	testDB()
// }
