// Medical Appointment Booking System - Dart API Client
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

import 'dart:convert';
import 'package:http/http.dart' as http;

/// ApiClient class to handle HTTP requests to the Medical Appointment Booking API
///
/// This client provides methods to interact with all endpoints of the medical booking system.
/// All methods return Futures and should be used with async/await.
///
/// Example usage:
/// ```dart
/// final apiClient = ApiClient();
///
/// // Get all clinics
/// List<Map<String, dynamic>> clinics = await apiClient.getClinics();
///
/// // Create a new patient
/// Map<String, dynamic> newPatient = {
///   'first_name': 'John',
///   'last_name': 'Doe',
///   'date_of_birth': '1990-05-15',
///   'phone': '+1234567890',
///   'email': 'john.doe@email.com',
///   'address': '123 Main St, City, State 12345',
///   'medical_record_number': 'MRN001'
/// };
/// Map<String, dynamic> createdPatient = await apiClient.createPatient(newPatient);
/// ```
class ApiClient {
  final String baseUrl;

  ApiClient({this.baseUrl = 'http://localhost:8080/api'});

  /// Clinics endpoints

  /// Retrieves all clinics from the system.
  ///
  /// Returns a list of clinic objects with their details.
  ///
  /// Example:
  /// ```dart
  /// List<Map<String, dynamic>> clinics = await apiClient.getClinics();
  /// for (var clinic in clinics) {
  ///   print('Clinic: ${clinic['name']} - ${clinic['address']}');
  /// }
  /// ```
  Future<List<Map<String, dynamic>>> getClinics() async {
    final response = await http.get(Uri.parse('$baseUrl/clinics'));
    if (response.statusCode == 200) {
      return List<Map<String, dynamic>>.from(json.decode(response.body));
    } else {
      throw Exception('Failed to load clinics');
    }
  }

  /// Retrieves a specific clinic by its ID.
  ///
  /// [id] - The unique identifier of the clinic.
  ///
  /// Returns the clinic object with full details.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> clinic = await apiClient.getClinic(1);
  /// print('Clinic Name: ${clinic['name']}');
  /// print('Phone: ${clinic['phone']}');
  /// ```
  Future<Map<String, dynamic>> getClinic(int id) async {
    final response = await http.get(Uri.parse('$baseUrl/clinics/$id'));
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to load clinic');
    }
  }

  /// Creates a new clinic in the system.
  ///
  /// [clinic] - A map containing clinic information.
  ///
  /// Required fields: name, address, phone, email
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> newClinic = {
  ///   'name': 'Downtown Medical Center',
  ///   'address': '456 Health St, Downtown, NY 10001',
  ///   'phone': '+1-212-555-0123',
  ///   'email': 'info@downtownmedical.com'
  /// };
  /// Map<String, dynamic> createdClinic = await apiClient.createClinic(newClinic);
  /// print('Created clinic with ID: ${createdClinic['id']}');
  /// ```
  Future<Map<String, dynamic>> createClinic(Map<String, dynamic> clinic) async {
    final response = await http.post(
      Uri.parse('$baseUrl/clinics'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(clinic),
    );
    if (response.statusCode == 201) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to create clinic');
    }
  }

  /// Updates an existing clinic's information.
  ///
  /// [id] - The unique identifier of the clinic to update.
  /// [clinic] - A map containing updated clinic information.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> updatedClinic = {
  ///   'name': 'Downtown Medical Center - Updated',
  ///   'address': '456 Health St, Downtown, NY 10001',
  ///   'phone': '+1-212-555-0124', // Changed phone number
  ///   'email': 'info@downtownmedical.com'
  /// };
  /// Map<String, dynamic> result = await apiClient.updateClinic(1, updatedClinic);
  /// ```
  Future<Map<String, dynamic>> updateClinic(
      int id, Map<String, dynamic> clinic) async {
    final response = await http.put(
      Uri.parse('$baseUrl/clinics/$id'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(clinic),
    );
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to update clinic');
    }
  }

  /// Deletes a clinic from the system.
  ///
  /// [id] - The unique identifier of the clinic to delete.
  ///
  /// Example:
  /// ```dart
  /// await apiClient.deleteClinic(1);
  /// print('Clinic deleted successfully');
  /// ```
  Future<void> deleteClinic(int id) async {
    final response = await http.delete(Uri.parse('$baseUrl/clinics/$id'));
    if (response.statusCode != 204) {
      throw Exception('Failed to delete clinic');
    }
  }

  /// Patients endpoints

  /// Retrieves all patients from the system.
  ///
  /// Returns a list of patient objects with their medical information.
  ///
  /// Example:
  /// ```dart
  /// List<Map<String, dynamic>> patients = await apiClient.getPatients();
  /// for (var patient in patients) {
  ///   print('${patient['first_name']} ${patient['last_name']} - ${patient['medical_record_number']}');
  /// }
  /// ```
  Future<List<Map<String, dynamic>>> getPatients() async {
    final response = await http.get(Uri.parse('$baseUrl/patients'));
    if (response.statusCode == 200) {
      return List<Map<String, dynamic>>.from(json.decode(response.body));
    } else {
      throw Exception('Failed to load patients');
    }
  }

  /// Retrieves a specific patient by their ID.
  ///
  /// [id] - The unique identifier of the patient.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> patient = await apiClient.getPatient(1);
  /// print('Patient: ${patient['first_name']} ${patient['last_name']}');
  /// print('DOB: ${patient['date_of_birth']}');
  /// print('Medical Record: ${patient['medical_record_number']}');
  /// ```
  Future<Map<String, dynamic>> getPatient(int id) async {
    final response = await http.get(Uri.parse('$baseUrl/patients/$id'));
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to load patient');
    }
  }

  /// Creates a new patient record in the system.
  ///
  /// [patient] - A map containing patient information.
  ///
  /// Required fields: first_name, last_name, date_of_birth, phone, email, address, medical_record_number
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> newPatient = {
  ///   'first_name': 'Jane',
  ///   'last_name': 'Smith',
  ///   'date_of_birth': '1985-03-20',
  ///   'phone': '+1-555-123-4567',
  ///   'email': 'jane.smith@email.com',
  ///   'address': '789 Oak Ave, Springfield, IL 62701',
  ///   'medical_record_number': 'MRN002'
  /// };
  /// Map<String, dynamic> createdPatient = await apiClient.createPatient(newPatient);
  /// print('Created patient with ID: ${createdPatient['id']}');
  /// ```
  Future<Map<String, dynamic>> createPatient(
      Map<String, dynamic> patient) async {
    final response = await http.post(
      Uri.parse('$baseUrl/patients'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(patient),
    );
    if (response.statusCode == 201) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to create patient');
    }
  }

  /// Updates an existing patient's information.
  ///
  /// [id] - The unique identifier of the patient to update.
  /// [patient] - A map containing updated patient information.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> updatedPatient = {
  ///   'first_name': 'Jane',
  ///   'last_name': 'Johnson', // Changed last name
  ///   'date_of_birth': '1985-03-20',
  ///   'phone': '+1-555-123-4567',
  ///   'email': 'jane.johnson@email.com', // Updated email
  ///   'address': '789 Oak Ave, Springfield, IL 62701',
  ///   'medical_record_number': 'MRN002'
  /// };
  /// Map<String, dynamic> result = await apiClient.updatePatient(1, updatedPatient);
  /// ```
  Future<Map<String, dynamic>> updatePatient(
      int id, Map<String, dynamic> patient) async {
    final response = await http.put(
      Uri.parse('$baseUrl/patients/$id'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(patient),
    );
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to update patient');
    }
  }

  /// Deletes a patient record from the system.
  ///
  /// [id] - The unique identifier of the patient to delete.
  ///
  /// Example:
  /// ```dart
  /// await apiClient.deletePatient(1);
  /// print('Patient deleted successfully');
  /// ```
  Future<void> deletePatient(int id) async {
    final response = await http.delete(Uri.parse('$baseUrl/patients/$id'));
    if (response.statusCode != 204) {
      throw Exception('Failed to delete patient');
    }
  }

  /// Employees endpoints

  /// Retrieves all employees from the system.
  ///
  /// Returns a list of employee objects with their roles and specializations.
  ///
  /// Example:
  /// ```dart
  /// List<Map<String, dynamic>> employees = await apiClient.getEmployees();
  /// for (var employee in employees) {
  ///   print('Dr. ${employee['first_name']} ${employee['last_name']} - ${employee['specialization']}');
  /// }
  /// ```
  Future<List<Map<String, dynamic>>> getEmployees() async {
    final response = await http.get(Uri.parse('$baseUrl/employees'));
    if (response.statusCode == 200) {
      return List<Map<String, dynamic>>.from(json.decode(response.body));
    } else {
      throw Exception('Failed to load employees');
    }
  }

  /// Retrieves a specific employee by their ID.
  ///
  /// [id] - The unique identifier of the employee.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> employee = await apiClient.getEmployee(1);
  /// print('Employee: ${employee['first_name']} ${employee['last_name']}');
  /// print('Role: ${employee['role']}');
  /// print('Specialization: ${employee['specialization']}');
  /// ```
  Future<Map<String, dynamic>> getEmployee(int id) async {
    final response = await http.get(Uri.parse('$baseUrl/employees/$id'));
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to load employee');
    }
  }

  /// Creates a new employee record in the system.
  ///
  /// [employee] - A map containing employee information.
  ///
  /// Required fields: first_name, last_name, role, specialization, phone, email
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> newEmployee = {
  ///   'first_name': 'Dr. Sarah',
  ///   'last_name': 'Williams',
  ///   'role': 'doctor',
  ///   'specialization': 'cardiology',
  ///   'phone': '+1-555-987-6543',
  ///   'email': 'sarah.williams@clinic.com'
  /// };
  /// Map<String, dynamic> createdEmployee = await apiClient.createEmployee(newEmployee);
  /// print('Created employee with ID: ${createdEmployee['id']}');
  /// ```
  Future<Map<String, dynamic>> createEmployee(
      Map<String, dynamic> employee) async {
    final response = await http.post(
      Uri.parse('$baseUrl/employees'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(employee),
    );
    if (response.statusCode == 201) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to create employee');
    }
  }

  /// Updates an existing employee's information.
  ///
  /// [id] - The unique identifier of the employee to update.
  /// [employee] - A map containing updated employee information.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> updatedEmployee = {
  ///   'first_name': 'Dr. Sarah',
  ///   'last_name': 'Williams',
  ///   'role': 'doctor',
  ///   'specialization': 'interventional cardiology', // Updated specialization
  ///   'phone': '+1-555-987-6543',
  ///   'email': 'sarah.williams@clinic.com'
  /// };
  /// Map<String, dynamic> result = await apiClient.updateEmployee(1, updatedEmployee);
  /// ```
  Future<Map<String, dynamic>> updateEmployee(
      int id, Map<String, dynamic> employee) async {
    final response = await http.put(
      Uri.parse('$baseUrl/employees/$id'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(employee),
    );
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to update employee');
    }
  }

  /// Deletes an employee record from the system.
  ///
  /// [id] - The unique identifier of the employee to delete.
  ///
  /// Example:
  /// ```dart
  /// await apiClient.deleteEmployee(1);
  /// print('Employee deleted successfully');
  /// ```
  Future<void> deleteEmployee(int id) async {
    final response = await http.delete(Uri.parse('$baseUrl/employees/$id'));
    if (response.statusCode != 204) {
      throw Exception('Failed to delete employee');
    }
  }

  /// Services endpoints

  /// Retrieves all medical services offered by the clinic.
  ///
  /// Returns a list of service objects with pricing and duration information.
  ///
  /// Example:
  /// ```dart
  /// List<Map<String, dynamic>> services = await apiClient.getServices();
  /// for (var service in services) {
  ///   print('${service['name']}: \$${service['price']} (${service['duration_minutes']} min)');
  /// }
  /// ```
  Future<List<Map<String, dynamic>>> getServices() async {
    final response = await http.get(Uri.parse('$baseUrl/services'));
    if (response.statusCode == 200) {
      return List<Map<String, dynamic>>.from(json.decode(response.body));
    } else {
      throw Exception('Failed to load services');
    }
  }

  /// Retrieves a specific service by its ID.
  ///
  /// [id] - The unique identifier of the service.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> service = await apiClient.getService(1);
  /// print('Service: ${service['name']}');
  /// print('Description: ${service['description']}');
  /// print('Duration: ${service['duration_minutes']} minutes');
  /// print('Price: \$${service['price']}');
  /// ```
  Future<Map<String, dynamic>> getService(int id) async {
    final response = await http.get(Uri.parse('$baseUrl/services/$id'));
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to load service');
    }
  }

  /// Creates a new medical service in the system.
  ///
  /// [service] - A map containing service information.
  ///
  /// Required fields: name, description, duration_minutes, price
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> newService = {
  ///   'name': 'General Consultation',
  ///   'description': 'Comprehensive health check and consultation with physician',
  ///   'duration_minutes': 30,
  ///   'price': 150.00
  /// };
  /// Map<String, dynamic> createdService = await apiClient.createService(newService);
  /// print('Created service with ID: ${createdService['id']}');
  /// ```
  Future<Map<String, dynamic>> createService(
      Map<String, dynamic> service) async {
    final response = await http.post(
      Uri.parse('$baseUrl/services'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(service),
    );
    if (response.statusCode == 201) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to create service');
    }
  }

  /// Updates an existing service's information.
  ///
  /// [id] - The unique identifier of the service to update.
  /// [service] - A map containing updated service information.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> updatedService = {
  ///   'name': 'General Consultation',
  ///   'description': 'Comprehensive health check and consultation with physician',
  ///   'duration_minutes': 45, // Extended duration
  ///   'price': 175.00 // Price increase
  /// };
  /// Map<String, dynamic> result = await apiClient.updateService(1, updatedService);
  /// ```
  Future<Map<String, dynamic>> updateService(
      int id, Map<String, dynamic> service) async {
    final response = await http.put(
      Uri.parse('$baseUrl/services/$id'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(service),
    );
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to update service');
    }
  }

  /// Deletes a service from the system.
  ///
  /// [id] - The unique identifier of the service to delete.
  ///
  /// Example:
  /// ```dart
  /// await apiClient.deleteService(1);
  /// print('Service deleted successfully');
  /// ```
  Future<void> deleteService(int id) async {
    final response = await http.delete(Uri.parse('$baseUrl/services/$id'));
    if (response.statusCode != 204) {
      throw Exception('Failed to delete service');
    }
  }

  /// Appointments endpoints

  /// Retrieves all appointments from the system.
  ///
  /// Returns a list of appointment objects with patient, employee, and service details.
  ///
  /// Example:
  /// ```dart
  /// List<Map<String, dynamic>> appointments = await apiClient.getAppointments();
  /// for (var appointment in appointments) {
  ///   print('Appointment on ${appointment['appointment_date']} - Status: ${appointment['status']}');
  /// }
  /// ```
  Future<List<Map<String, dynamic>>> getAppointments() async {
    final response = await http.get(Uri.parse('$baseUrl/appointments'));
    if (response.statusCode == 200) {
      return List<Map<String, dynamic>>.from(json.decode(response.body));
    } else {
      throw Exception('Failed to load appointments');
    }
  }

  /// Retrieves a specific appointment by its ID.
  ///
  /// [id] - The unique identifier of the appointment.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> appointment = await apiClient.getAppointment(1);
  /// print('Appointment Date: ${appointment['appointment_date']}');
  /// print('Status: ${appointment['status']}');
  /// print('Patient ID: ${appointment['patient_id']}');
  /// print('Doctor ID: ${appointment['employee_id']}');
  /// print('Service ID: ${appointment['service_id']}');
  /// print('Notes: ${appointment['notes']}');
  /// ```
  Future<Map<String, dynamic>> getAppointment(int id) async {
    final response = await http.get(Uri.parse('$baseUrl/appointments/$id'));
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to load appointment');
    }
  }

  /// Creates a new appointment in the system.
  ///
  /// [appointment] - A map containing appointment information.
  ///
  /// Required fields: patient_id, employee_id, service_id, clinic_id, appointment_date, status
  /// Optional fields: notes, payment_status
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> newAppointment = {
  ///   'patient_id': 1,        // ID of the patient
  ///   'employee_id': 2,       // ID of the doctor
  ///   'service_id': 1,        // ID of the service
  ///   'clinic_id': 1,         // ID of the clinic
  ///   'appointment_date': '2025-10-25T10:00:00Z', // ISO 8601 UTC format
  ///   'status': 'scheduled',  // scheduled, confirmed, completed, cancelled
  ///   'notes': 'Follow-up consultation for hypertension',
  ///   'payment_status': 'pending' // pending, paid, refunded
  /// };
  /// Map<String, dynamic> createdAppointment = await apiClient.createAppointment(newAppointment);
  /// print('Created appointment with ID: ${createdAppointment['id']}');
  /// ```
  Future<Map<String, dynamic>> createAppointment(
      Map<String, dynamic> appointment) async {
    final response = await http.post(
      Uri.parse('$baseUrl/appointments'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(appointment),
    );
    if (response.statusCode == 201) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to create appointment');
    }
  }

  /// Updates an existing appointment's information.
  ///
  /// [id] - The unique identifier of the appointment to update.
  /// [appointment] - A map containing updated appointment information.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> updatedAppointment = {
  ///   'patient_id': 1,
  ///   'employee_id': 2,
  ///   'service_id': 1,
  ///   'clinic_id': 1,
  ///   'appointment_date': '2025-10-25T11:00:00Z', // Changed time
  ///   'status': 'confirmed',  // Updated status
  ///   'notes': 'Follow-up consultation for hypertension - Confirmed',
  ///   'payment_status': 'paid' // Payment completed
  /// };
  /// Map<String, dynamic> result = await apiClient.updateAppointment(1, updatedAppointment);
  /// ```
  Future<Map<String, dynamic>> updateAppointment(
      int id, Map<String, dynamic> appointment) async {
    final response = await http.put(
      Uri.parse('$baseUrl/appointments/$id'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(appointment),
    );
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to update appointment');
    }
  }

  /// Deletes an appointment from the system.
  ///
  /// [id] - The unique identifier of the appointment to delete.
  ///
  /// Example:
  /// ```dart
  /// await apiClient.deleteAppointment(1);
  /// print('Appointment cancelled and deleted successfully');
  /// ```
  Future<void> deleteAppointment(int id) async {
    final response = await http.delete(Uri.parse('$baseUrl/appointments/$id'));
    if (response.statusCode != 204) {
      throw Exception('Failed to delete appointment');
    }
  }

  /// Waiting List endpoints

  /// Retrieves all items from the waiting list.
  ///
  /// Returns a list of waiting list items with patient requests for services.
  ///
  /// Example:
  /// ```dart
  /// List<Map<String, dynamic>> waitingList = await apiClient.getWaitingList();
  /// for (var item in waitingList) {
  ///   print('Patient ${item['patient_id']} waiting for service ${item['service_id']}');
  ///   print('Priority: ${item['priority']} - Requested: ${item['requested_date']}');
  /// }
  /// ```
  Future<List<Map<String, dynamic>>> getWaitingList() async {
    final response = await http.get(Uri.parse('$baseUrl/waiting-list'));
    if (response.statusCode == 200) {
      return List<Map<String, dynamic>>.from(json.decode(response.body));
    } else {
      throw Exception('Failed to load waiting list');
    }
  }

  /// Retrieves a specific waiting list item by its ID.
  ///
  /// [id] - The unique identifier of the waiting list item.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> item = await apiClient.getWaitingListItem(1);
  /// print('Patient ID: ${item['patient_id']}');
  /// print('Service ID: ${item['service_id']}');
  /// print('Priority: ${item['priority']}');
  /// print('Requested Date: ${item['requested_date']}');
  /// print('Notes: ${item['notes']}');
  /// ```
  Future<Map<String, dynamic>> getWaitingListItem(int id) async {
    final response = await http.get(Uri.parse('$baseUrl/waiting-list/$id'));
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to load waiting list item');
    }
  }

  /// Adds a patient to the waiting list for a service.
  ///
  /// [item] - A map containing waiting list item information.
  ///
  /// Required fields: patient_id, service_id, clinic_id, requested_date, priority
  /// Optional fields: notes
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> newWaitingItem = {
  ///   'patient_id': 3,        // ID of the patient
  ///   'service_id': 2,        // ID of the requested service
  ///   'clinic_id': 1,         // ID of the clinic
  ///   'requested_date': '2025-10-20T09:00:00Z', // When they want the appointment
  ///   'priority': 'high',     // high, medium, low
  ///   'notes': 'Patient prefers morning appointments, has mobility issues'
  /// };
  /// Map<String, dynamic> createdItem = await apiClient.createWaitingListItem(newWaitingItem);
  /// print('Added to waiting list with ID: ${createdItem['id']}');
  /// ```
  Future<Map<String, dynamic>> createWaitingListItem(
      Map<String, dynamic> item) async {
    final response = await http.post(
      Uri.parse('$baseUrl/waiting-list'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(item),
    );
    if (response.statusCode == 201) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to create waiting list item');
    }
  }

  /// Updates a waiting list item's information.
  ///
  /// [id] - The unique identifier of the waiting list item to update.
  /// [item] - A map containing updated waiting list item information.
  ///
  /// Example:
  /// ```dart
  /// Map<String, dynamic> updatedItem = {
  ///   'patient_id': 3,
  ///   'service_id': 2,
  ///   'clinic_id': 1,
  ///   'requested_date': '2025-10-22T14:00:00Z', // Changed preferred time
  ///   'priority': 'urgent',    // Increased priority
  ///   'notes': 'Patient prefers afternoon appointments now, urgent medical need'
  /// };
  /// Map<String, dynamic> result = await apiClient.updateWaitingListItem(1, updatedItem);
  /// ```
  Future<Map<String, dynamic>> updateWaitingListItem(
      int id, Map<String, dynamic> item) async {
    final response = await http.put(
      Uri.parse('$baseUrl/waiting-list/$id'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode(item),
    );
    if (response.statusCode == 200) {
      return json.decode(response.body);
    } else {
      throw Exception('Failed to update waiting list item');
    }
  }

  /// Removes an item from the waiting list.
  ///
  /// [id] - The unique identifier of the waiting list item to delete.
  ///
  /// Example:
  /// ```dart
  /// await apiClient.deleteWaitingListItem(1);
  /// print('Removed from waiting list successfully');
  /// ```
  Future<void> deleteWaitingListItem(int id) async {
    final response = await http.delete(Uri.parse('$baseUrl/waiting-list/$id'));
    if (response.statusCode != 204) {
      throw Exception('Failed to delete waiting list item');
    }
  }
}
