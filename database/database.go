// Medical Appointment Booking System - Database Package
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

package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"bookings/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// InitDB initializes the database connection
func InitDB() {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL environment variable is not set. Please set it to your PostgreSQL connection string.")
	}

	var err error
	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// Clinic CRUD operations
func GetClinics() ([]models.Clinic, error) {
	rows, err := DB.Query(context.Background(), "SELECT id, name, address, phone, email, active FROM clinics ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clinics []models.Clinic
	for rows.Next() {
		var clinic models.Clinic
		err := rows.Scan(&clinic.ID, &clinic.Name, &clinic.Address, &clinic.Phone, &clinic.Email, &clinic.Active)
		if err != nil {
			return nil, err
		}
		clinics = append(clinics, clinic)
	}
	return clinics, nil
}

func GetClinic(id int) (*models.Clinic, error) {
	var clinic models.Clinic
	err := DB.QueryRow(context.Background(),
		"SELECT id, name, address, phone, email, active FROM clinics WHERE id = $1", id).
		Scan(&clinic.ID, &clinic.Name, &clinic.Address, &clinic.Phone, &clinic.Email, &clinic.Active)
	if err != nil {
		return nil, err
	}
	return &clinic, nil
}

func CreateClinic(clinic *models.Clinic) error {
	return DB.QueryRow(context.Background(),
		"INSERT INTO clinics (name, address, phone, email, active) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		clinic.Name, clinic.Address, clinic.Phone, clinic.Email, clinic.Active).Scan(&clinic.ID)
}

func UpdateClinic(id int, clinic *models.Clinic) error {
	_, err := DB.Exec(context.Background(),
		"UPDATE clinics SET name = $1, address = $2, phone = $3, email = $4, active = $5 WHERE id = $6",
		clinic.Name, clinic.Address, clinic.Phone, clinic.Email, clinic.Active, id)
	return err
}

func DeleteClinic(id int) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM clinics WHERE id = $1", id)
	return err
}

// Patient CRUD operations
func GetPatients() ([]models.Patient, error) {
	rows, err := DB.Query(context.Background(),
		"SELECT id, first_name, last_name, email, phone, date_of_birth, medical_record_number, insurance_provider, insurance_id, emergency_contact_name, emergency_contact_phone, active, created_at FROM patients ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		err := rows.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Email, &patient.Phone,
			&patient.DateOfBirth, &patient.MedicalRecordNumber, &patient.InsuranceProvider, &patient.InsuranceID,
			&patient.EmergencyContactName, &patient.EmergencyContactPhone, &patient.Active, &patient.CreatedAt)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

func GetPatient(id int) (*models.Patient, error) {
	var patient models.Patient
	err := DB.QueryRow(context.Background(),
		"SELECT id, first_name, last_name, email, phone, date_of_birth, medical_record_number, insurance_provider, insurance_id, emergency_contact_name, emergency_contact_phone, active, created_at FROM patients WHERE id = $1", id).
		Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Email, &patient.Phone,
			&patient.DateOfBirth, &patient.MedicalRecordNumber, &patient.InsuranceProvider, &patient.InsuranceID,
			&patient.EmergencyContactName, &patient.EmergencyContactPhone, &patient.Active, &patient.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func CreatePatient(patient *models.Patient) error {
	return DB.QueryRow(context.Background(),
		"INSERT INTO patients (first_name, last_name, email, phone, date_of_birth, medical_record_number, insurance_provider, insurance_id, emergency_contact_name, emergency_contact_phone, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		patient.FirstName, patient.LastName, patient.Email, patient.Phone, patient.DateOfBirth,
		patient.MedicalRecordNumber, patient.InsuranceProvider, patient.InsuranceID,
		patient.EmergencyContactName, patient.EmergencyContactPhone, patient.Active).Scan(&patient.ID)
}

func UpdatePatient(id int, patient *models.Patient) error {
	_, err := DB.Exec(context.Background(),
		"UPDATE patients SET first_name = $1, last_name = $2, email = $3, phone = $4, date_of_birth = $5, medical_record_number = $6, insurance_provider = $7, insurance_id = $8, emergency_contact_name = $9, emergency_contact_phone = $10, active = $11 WHERE id = $12",
		patient.FirstName, patient.LastName, patient.Email, patient.Phone, patient.DateOfBirth,
		patient.MedicalRecordNumber, patient.InsuranceProvider, patient.InsuranceID,
		patient.EmergencyContactName, patient.EmergencyContactPhone, patient.Active, id)
	return err
}

func DeletePatient(id int) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM patients WHERE id = $1", id)
	return err
}

// Employee CRUD operations
func GetEmployees() ([]models.Employee, error) {
	rows, err := DB.Query(context.Background(),
		"SELECT id, clinic_id, first_name, last_name, email, phone, license_number, specialty, timezone, active, created_at FROM employees ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		err := rows.Scan(&employee.ID, &employee.ClinicID, &employee.FirstName, &employee.LastName,
			&employee.Email, &employee.Phone, &employee.LicenseNumber, &employee.Specialty,
			&employee.Timezone, &employee.Active, &employee.CreatedAt)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func GetEmployee(id int) (*models.Employee, error) {
	var employee models.Employee
	err := DB.QueryRow(context.Background(),
		"SELECT id, clinic_id, first_name, last_name, email, phone, license_number, specialty, timezone, active, created_at FROM employees WHERE id = $1", id).
		Scan(&employee.ID, &employee.ClinicID, &employee.FirstName, &employee.LastName,
			&employee.Email, &employee.Phone, &employee.LicenseNumber, &employee.Specialty,
			&employee.Timezone, &employee.Active, &employee.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func CreateEmployee(employee *models.Employee) error {
	return DB.QueryRow(context.Background(),
		"INSERT INTO employees (clinic_id, first_name, last_name, email, phone, license_number, specialty, timezone, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		employee.ClinicID, employee.FirstName, employee.LastName, employee.Email, employee.Phone,
		employee.LicenseNumber, employee.Specialty, employee.Timezone, employee.Active).Scan(&employee.ID)
}

func UpdateEmployee(id int, employee *models.Employee) error {
	_, err := DB.Exec(context.Background(),
		"UPDATE employees SET clinic_id = $1, first_name = $2, last_name = $3, email = $4, phone = $5, license_number = $6, specialty = $7, timezone = $8, active = $9 WHERE id = $10",
		employee.ClinicID, employee.FirstName, employee.LastName, employee.Email, employee.Phone,
		employee.LicenseNumber, employee.Specialty, employee.Timezone, employee.Active, id)
	return err
}

func DeleteEmployee(id int) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM employees WHERE id = $1", id)
	return err
}

// Service CRUD operations
func GetServices() ([]models.Service, error) {
	rows, err := DB.Query(context.Background(),
		"SELECT id, name, description, duration_minutes, price, specialty_required, active FROM services ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		err := rows.Scan(&service.ID, &service.Name, &service.Description, &service.DurationMinutes,
			&service.Price, &service.SpecialtyRequired, &service.Active)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func GetService(id int) (*models.Service, error) {
	var service models.Service
	err := DB.QueryRow(context.Background(),
		"SELECT id, name, description, duration_minutes, price, specialty_required, active FROM services WHERE id = $1", id).
		Scan(&service.ID, &service.Name, &service.Description, &service.DurationMinutes,
			&service.Price, &service.SpecialtyRequired, &service.Active)
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func CreateService(service *models.Service) error {
	return DB.QueryRow(context.Background(),
		"INSERT INTO services (name, description, duration_minutes, price, specialty_required, active) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		service.Name, service.Description, service.DurationMinutes, service.Price, service.SpecialtyRequired, service.Active).Scan(&service.ID)
}

func UpdateService(id int, service *models.Service) error {
	_, err := DB.Exec(context.Background(),
		"UPDATE services SET name = $1, description = $2, duration_minutes = $3, price = $4, specialty_required = $5, active = $6 WHERE id = $7",
		service.Name, service.Description, service.DurationMinutes, service.Price, service.SpecialtyRequired, service.Active, id)
	return err
}

func DeleteService(id int) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM services WHERE id = $1", id)
	return err
}

// Appointment CRUD operations
func GetAppointments() ([]models.Appointment, error) {
	rows, err := DB.Query(context.Background(),
		"SELECT id, patient_id, employee_id, service_id, clinic_id, start_datetime, end_datetime, status, appointment_type, notes, medical_notes, cancellation_reason, payment_status, payment_amount, created_at, updated_at FROM appointments ORDER BY start_datetime DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		err := rows.Scan(&appointment.ID, &appointment.PatientID, &appointment.EmployeeID, &appointment.ServiceID,
			&appointment.ClinicID, &appointment.StartDatetime, &appointment.EndDatetime, &appointment.Status,
			&appointment.AppointmentType, &appointment.Notes, &appointment.MedicalNotes, &appointment.CancellationReason,
			&appointment.PaymentStatus, &appointment.PaymentAmount, &appointment.CreatedAt, &appointment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}

func GetAppointment(id int) (*models.Appointment, error) {
	var appointment models.Appointment
	err := DB.QueryRow(context.Background(),
		"SELECT id, patient_id, employee_id, service_id, clinic_id, start_datetime, end_datetime, status, appointment_type, notes, medical_notes, cancellation_reason, payment_status, payment_amount, created_at, updated_at FROM appointments WHERE id = $1", id).
		Scan(&appointment.ID, &appointment.PatientID, &appointment.EmployeeID, &appointment.ServiceID,
			&appointment.ClinicID, &appointment.StartDatetime, &appointment.EndDatetime, &appointment.Status,
			&appointment.AppointmentType, &appointment.Notes, &appointment.MedicalNotes, &appointment.CancellationReason,
			&appointment.PaymentStatus, &appointment.PaymentAmount, &appointment.CreatedAt, &appointment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func CreateAppointment(appointment *models.Appointment) error {
	return DB.QueryRow(context.Background(),
		"INSERT INTO appointments (patient_id, employee_id, service_id, clinic_id, start_datetime, end_datetime, status, appointment_type, notes, payment_status, payment_amount) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		appointment.PatientID, appointment.EmployeeID, appointment.ServiceID, appointment.ClinicID,
		appointment.StartDatetime.UTC(), appointment.EndDatetime.UTC(), appointment.Status, appointment.AppointmentType,
		appointment.Notes, appointment.PaymentStatus, appointment.PaymentAmount).Scan(&appointment.ID)
}

func UpdateAppointment(id int, appointment *models.Appointment) error {
	_, err := DB.Exec(context.Background(),
		"UPDATE appointments SET patient_id = $1, employee_id = $2, service_id = $3, clinic_id = $4, start_datetime = $5, end_datetime = $6, status = $7, appointment_type = $8, notes = $9, medical_notes = $10, cancellation_reason = $11, payment_status = $12, payment_amount = $13, updated_at = CURRENT_TIMESTAMP WHERE id = $14",
		appointment.PatientID, appointment.EmployeeID, appointment.ServiceID, appointment.ClinicID,
		appointment.StartDatetime.UTC(), appointment.EndDatetime.UTC(), appointment.Status, appointment.AppointmentType,
		appointment.Notes, appointment.MedicalNotes, appointment.CancellationReason,
		appointment.PaymentStatus, appointment.PaymentAmount, id)
	return err
}

func DeleteAppointment(id int) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM appointments WHERE id = $1", id)
	return err
}

// Waiting List CRUD operations
func GetWaitingList() ([]models.WaitingList, error) {
	rows, err := DB.Query(context.Background(),
		"SELECT id, patient_id, service_id, preferred_employee_id, requested_date, urgency_level, notes, status, created_at FROM waiting_list ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var waitingList []models.WaitingList
	for rows.Next() {
		var item models.WaitingList
		err := rows.Scan(&item.ID, &item.PatientID, &item.ServiceID, &item.PreferredEmployeeID,
			&item.RequestedDate, &item.UrgencyLevel, &item.Notes, &item.Status, &item.CreatedAt)
		if err != nil {
			return nil, err
		}
		waitingList = append(waitingList, item)
	}
	return waitingList, nil
}

func GetWaitingListItem(id int) (*models.WaitingList, error) {
	var item models.WaitingList
	err := DB.QueryRow(context.Background(),
		"SELECT id, patient_id, service_id, preferred_employee_id, requested_date, urgency_level, notes, status, created_at FROM waiting_list WHERE id = $1", id).
		Scan(&item.ID, &item.PatientID, &item.ServiceID, &item.PreferredEmployeeID,
			&item.RequestedDate, &item.UrgencyLevel, &item.Notes, &item.Status, &item.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func CreateWaitingListItem(item *models.WaitingList) error {
	return DB.QueryRow(context.Background(),
		"INSERT INTO waiting_list (patient_id, service_id, preferred_employee_id, requested_date, urgency_level, notes, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		item.PatientID, item.ServiceID, item.PreferredEmployeeID, item.RequestedDate,
		item.UrgencyLevel, item.Notes, item.Status).Scan(&item.ID)
}

func UpdateWaitingListItem(id int, item *models.WaitingList) error {
	_, err := DB.Exec(context.Background(),
		"UPDATE waiting_list SET patient_id = $1, service_id = $2, preferred_employee_id = $3, requested_date = $4, urgency_level = $5, notes = $6, status = $7 WHERE id = $8",
		item.PatientID, item.ServiceID, item.PreferredEmployeeID, item.RequestedDate,
		item.UrgencyLevel, item.Notes, item.Status, id)
	return err
}

func DeleteWaitingListItem(id int) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM waiting_list WHERE id = $1", id)
	return err
}

// CreateTables creates all necessary database tables and indexes
func CreateTables() error {
	statements := []string{
		// Drop existing tables if they exist (in reverse order due to foreign keys)
		`DROP TABLE IF EXISTS waiting_list CASCADE`,
		`DROP TABLE IF EXISTS appointments CASCADE`,
		`DROP TABLE IF EXISTS slot_holds CASCADE`,
		`DROP TABLE IF EXISTS time_off CASCADE`,
		`DROP TABLE IF EXISTS day_overrides CASCADE`,
		`DROP TABLE IF EXISTS work_templates CASCADE`,
		`DROP TABLE IF EXISTS employee_services CASCADE`,
		`DROP TABLE IF EXISTS services CASCADE`,
		`DROP TABLE IF EXISTS employees CASCADE`,
		`DROP TABLE IF EXISTS patients CASCADE`,
		`DROP TABLE IF EXISTS clinics CASCADE`,

		// Drop existing types if they exist
		`DROP TYPE IF EXISTS appointment_status CASCADE`,
		`DROP TYPE IF EXISTS appointment_type CASCADE`,
		`DROP TYPE IF EXISTS payment_status CASCADE`,
		`DROP TYPE IF EXISTS urgency_level CASCADE`,
		`DROP TYPE IF EXISTS waiting_list_status CASCADE`,

		// Create enum types
		`CREATE TYPE appointment_status AS ENUM ('SCHEDULED', 'CONFIRMED', 'IN_PROGRESS', 'COMPLETED', 'CANCELLED', 'NO_SHOW')`,
		`CREATE TYPE appointment_type AS ENUM ('INITIAL_CONSULTATION', 'FOLLOW_UP', 'PROCEDURE', 'EMERGENCY')`,
		`CREATE TYPE payment_status AS ENUM ('PENDING', 'PAID', 'REFUNDED')`,
		`CREATE TYPE urgency_level AS ENUM ('LOW', 'MEDIUM', 'HIGH', 'URGENT')`,
		`CREATE TYPE waiting_list_status AS ENUM ('ACTIVE', 'CONTACTED', 'SCHEDULED', 'EXPIRED')`,

		// Create tables
		`CREATE TABLE IF NOT EXISTS clinics (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			address TEXT,
			phone TEXT,
			email TEXT,
			active BOOLEAN DEFAULT TRUE
		)`,
		`CREATE TABLE IF NOT EXISTS patients (
			id SERIAL PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			email TEXT UNIQUE,
			phone TEXT,
			date_of_birth TEXT,
			medical_record_number TEXT UNIQUE,
			insurance_provider TEXT,
			insurance_id TEXT,
			emergency_contact_name TEXT,
			emergency_contact_phone TEXT,
			active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS employees (
			id SERIAL PRIMARY KEY,
			clinic_id INTEGER NOT NULL REFERENCES clinics(id),
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			email TEXT UNIQUE,
			phone TEXT,
			license_number TEXT UNIQUE,
			specialty TEXT,
			timezone TEXT DEFAULT 'Asia/Colombo',
			active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS services (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL UNIQUE,
			description TEXT,
			duration_minutes INTEGER NOT NULL CHECK (duration_minutes > 0),
			price DECIMAL,
			specialty_required TEXT,
			active BOOLEAN DEFAULT TRUE
		)`,
		`CREATE TABLE IF NOT EXISTS employee_services (
			employee_id INTEGER NOT NULL REFERENCES employees(id) ON DELETE CASCADE,
			service_id INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE,
			UNIQUE (employee_id, service_id)
		)`,
		`CREATE TABLE IF NOT EXISTS work_templates (
			id SERIAL PRIMARY KEY,
			employee_id INTEGER NOT NULL REFERENCES employees(id),
			weekday INTEGER NOT NULL CHECK (weekday >= 1 AND weekday <= 7),
			start_time TIME,
			end_time TIME,
			slot_granularity_minutes INTEGER DEFAULT 15,
			is_active BOOLEAN DEFAULT TRUE
		)`,
		`CREATE TABLE IF NOT EXISTS day_overrides (
			id SERIAL PRIMARY KEY,
			employee_id INTEGER NOT NULL REFERENCES employees(id),
			date DATE NOT NULL,
			is_closed BOOLEAN DEFAULT FALSE,
			start_time TIME,
			end_time TIME,
			reason TEXT,
			UNIQUE (employee_id, date)
		)`,
		`CREATE TABLE IF NOT EXISTS time_off (
			id SERIAL PRIMARY KEY,
			employee_id INTEGER NOT NULL REFERENCES employees(id),
			start_datetime TIMESTAMPTZ NOT NULL,
			end_datetime TIMESTAMPTZ NOT NULL,
			reason TEXT,
			approved BOOLEAN DEFAULT FALSE
		)`,
		`CREATE TABLE IF NOT EXISTS slot_holds (
			id SERIAL PRIMARY KEY,
			employee_id INTEGER NOT NULL REFERENCES employees(id),
			service_id INTEGER NOT NULL REFERENCES services(id),
			start_datetime TIMESTAMPTZ NOT NULL,
			end_datetime TIMESTAMPTZ NOT NULL,
			patient_id INTEGER REFERENCES patients(id),
			hold_token TEXT NOT NULL UNIQUE,
			expires_at TIMESTAMPTZ NOT NULL,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS appointments (
			id SERIAL PRIMARY KEY,
			patient_id INTEGER NOT NULL REFERENCES patients(id),
			employee_id INTEGER NOT NULL REFERENCES employees(id),
			service_id INTEGER NOT NULL REFERENCES services(id),
			clinic_id INTEGER NOT NULL REFERENCES clinics(id),
			start_datetime TIMESTAMPTZ NOT NULL,
			end_datetime TIMESTAMPTZ NOT NULL,
			status appointment_status DEFAULT 'SCHEDULED',
			appointment_type appointment_type,
			notes TEXT,
			medical_notes TEXT,
			cancellation_reason TEXT,
			payment_status payment_status DEFAULT 'PENDING',
			payment_amount DECIMAL,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS waiting_list (
			id SERIAL PRIMARY KEY,
			patient_id INTEGER NOT NULL REFERENCES patients(id),
			service_id INTEGER NOT NULL REFERENCES services(id),
			preferred_employee_id INTEGER REFERENCES employees(id),
			requested_date TEXT,
			urgency_level urgency_level DEFAULT 'MEDIUM',
			notes TEXT,
			status waiting_list_status DEFAULT 'ACTIVE',
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		)`,

		// Create indexes
		`CREATE INDEX IF NOT EXISTS idx_appointments_patient_id ON appointments(patient_id)`,
		`CREATE INDEX IF NOT EXISTS idx_appointments_employee_id ON appointments(employee_id)`,
		`CREATE INDEX IF NOT EXISTS idx_appointments_datetime ON appointments(start_datetime)`,
		`CREATE INDEX IF NOT EXISTS idx_appointments_status ON appointments(status)`,
		`CREATE INDEX IF NOT EXISTS idx_slot_holds_datetime ON slot_holds(start_datetime, end_datetime)`,
		`CREATE INDEX IF NOT EXISTS idx_time_off_datetime ON time_off(start_datetime, end_datetime)`,
	}

	for _, stmt := range statements {
		_, err := DB.Exec(context.Background(), stmt)
		if err != nil {
			return fmt.Errorf("failed to execute statement: %s, error: %v", stmt, err)
		}
	}

	fmt.Println("All tables and indexes created successfully!")
	return nil
}
