// Medical Appointment Booking System - Models Package
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

package models

import "time"

// Clinic represents a medical clinic
type Clinic struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	Phone   string `json:"phone" db:"phone"`
	Email   string `json:"email" db:"email"`
	Active  bool   `json:"active" db:"active"`
}

// Patient represents a patient
type Patient struct {
	ID                    int       `json:"id" db:"id"`
	FirstName             string    `json:"first_name" db:"first_name"`
	LastName              string    `json:"last_name" db:"last_name"`
	Email                 string    `json:"email" db:"email"`
	Phone                 string    `json:"phone" db:"phone"`
	DateOfBirth           *string   `json:"date_of_birth" db:"date_of_birth"`
	MedicalRecordNumber   string    `json:"medical_record_number" db:"medical_record_number"`
	InsuranceProvider     *string   `json:"insurance_provider" db:"insurance_provider"`
	InsuranceID           *string   `json:"insurance_id" db:"insurance_id"`
	EmergencyContactName  *string   `json:"emergency_contact_name" db:"emergency_contact_name"`
	EmergencyContactPhone *string   `json:"emergency_contact_phone" db:"emergency_contact_phone"`
	Active                bool      `json:"active" db:"active"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
}

// Employee represents a medical employee/doctor
type Employee struct {
	ID            int       `json:"id" db:"id"`
	ClinicID      int       `json:"clinic_id" db:"clinic_id"`
	FirstName     string    `json:"first_name" db:"first_name"`
	LastName      string    `json:"last_name" db:"last_name"`
	Email         string    `json:"email" db:"email"`
	Phone         string    `json:"phone" db:"phone"`
	LicenseNumber string    `json:"license_number" db:"license_number"`
	Specialty     string    `json:"specialty" db:"specialty"`
	Timezone      string    `json:"timezone" db:"timezone"`
	Active        bool      `json:"active" db:"active"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

// Service represents a medical service
type Service struct {
	ID                int     `json:"id" db:"id"`
	Name              string  `json:"name" db:"name"`
	Description       string  `json:"description" db:"description"`
	DurationMinutes   int     `json:"duration_minutes" db:"duration_minutes"`
	Price             float64 `json:"price" db:"price"`
	SpecialtyRequired string  `json:"specialty_required" db:"specialty_required"`
	Active            bool    `json:"active" db:"active"`
}

// Appointment represents a medical appointment
type Appointment struct {
	ID                 int       `json:"id" db:"id"`
	PatientID          int       `json:"patient_id" db:"patient_id"`
	EmployeeID         int       `json:"employee_id" db:"employee_id"`
	ServiceID          int       `json:"service_id" db:"service_id"`
	ClinicID           int       `json:"clinic_id" db:"clinic_id"`
	StartDatetime      time.Time `json:"start_datetime" db:"start_datetime"`
	EndDatetime        time.Time `json:"end_datetime" db:"end_datetime"`
	Status             string    `json:"status" db:"status"`
	AppointmentType    *string   `json:"appointment_type" db:"appointment_type"`
	Notes              *string   `json:"notes" db:"notes"`
	MedicalNotes       *string   `json:"medical_notes" db:"medical_notes"`
	CancellationReason *string   `json:"cancellation_reason" db:"cancellation_reason"`
	PaymentStatus      string    `json:"payment_status" db:"payment_status"`
	PaymentAmount      *float64  `json:"payment_amount" db:"payment_amount"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

// WaitingList represents a waiting list entry
type WaitingList struct {
	ID                  int       `json:"id" db:"id"`
	PatientID           int       `json:"patient_id" db:"patient_id"`
	ServiceID           int       `json:"service_id" db:"service_id"`
	PreferredEmployeeID *int      `json:"preferred_employee_id" db:"preferred_employee_id"`
	RequestedDate       *string   `json:"requested_date" db:"requested_date"`
	UrgencyLevel        string    `json:"urgency_level" db:"urgency_level"`
	Notes               *string   `json:"notes" db:"notes"`
	Status              string    `json:"status" db:"status"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
}
