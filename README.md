# Medical Appointment Booking System

A comprehensive Go-based REST API for managing medical appointments with PostgreSQL integration and a Dart client for mobile/web applications.

## Features

- **Clinic Management** - Manage multiple medical facilities with contact information
- **Patient Records** - Complete patient management with medical record numbers, insurance, and emergency contacts
- **Staff Management** - Medical staff (doctors, nurses) with specialties and license tracking
- **Service Catalog** - Medical services with pricing, duration, and specialty requirements
- **Appointment Scheduling** - Book, confirm, and manage appointments with status tracking
- **Waiting Lists** - Manage patient queues for popular services with urgency levels
- **Payment Tracking** - Track appointment payments and statuses
- **REST API** - Full REST API with JSON responses
- **Dart Client** - Ready-to-use Dart HTTP client for Flutter/web applications
- **UTC Time Handling** - All timestamps stored in UTC with proper timezone support
- **Comprehensive Testing** - Database and API testing suite

## Prerequisites

- Go 1.24.5 or later
- PostgreSQL database
- Dart SDK (for using the Dart client)

## Setup

1. **Install Go**: Download and install Go from [golang.org](https://golang.org/dl/)

2. **Install PostgreSQL**: Download and install PostgreSQL from [postgresql.org](https://www.postgresql.org/download/)

3. **Create a database**: Create a PostgreSQL database for the application:
   ```sql
   CREATE DATABASE bookings_db;
   CREATE USER bookings_user WITH PASSWORD 'your_secure_password';
   GRANT ALL PRIVILEGES ON DATABASE bookings_db TO bookings_user;
   ```

4. **Set database connection**: Set the `DATABASE_URL` environment variable with your database credentials:
   ```bash
   export DATABASE_URL="postgres://bookings_user:your_secure_password@localhost:5432/bookings_db?sslmode=disable"
   ```

   **Security Note**: Never commit database credentials to version control. Use environment variables or secure credential management systems in production.

## Environment Variables

The application requires the following environment variables:

- `DATABASE_URL`: PostgreSQL connection string (required)
  - Format: `postgres://username:password@host:port/database?sslmode=disable`

Example:
```bash
export DATABASE_URL="postgres://bookings_user:my_secure_password@localhost:5432/bookings_db?sslmode=disable"
```

## Database Schema

The application creates the following tables with PostgreSQL enums:

### Core Tables
- **clinics** - Medical facilities with contact information
- **patients** - Patient records with medical and insurance details
- **employees** - Medical staff with specialties and license information
- **services** - Medical services with pricing and duration
- **appointments** - Scheduled appointments with status tracking
- **waiting_list** - Patient waiting lists with urgency levels

### Supporting Tables
- **employee_services** - Junction table linking staff to services they provide
- **work_templates** - Weekly work schedules for employees
- **day_overrides** - Holiday and special schedule changes
- **time_off** - Staff vacation and leave tracking
- **slot_holds** - Temporary appointment reservations

### Enums
- **appointment_status**: SCHEDULED, CONFIRMED, IN_PROGRESS, COMPLETED, CANCELLED, NO_SHOW
- **appointment_type**: INITIAL_CONSULTATION, FOLLOW_UP, PROCEDURE, EMERGENCY
- **payment_status**: PENDING, PAID, REFUNDED
- **urgency_level**: LOW, MEDIUM, HIGH, URGENT
- **waiting_list_status**: ACTIVE, CONTACTED, SCHEDULED, EXPIRED

### Key Features
- All timestamps use TIMESTAMPTZ (UTC with timezone)
- Foreign key relationships with CASCADE deletes where appropriate
- Indexes on commonly queried fields (patient_id, employee_id, datetime, status)
- Nullable fields for optional data (insurance, emergency contacts, etc.)

## Running the Application

1. **Set environment variables**:
   ```bash
   export DATABASE_URL="postgres://bookings_user:your_secure_password@localhost:5432/bookings_db?sslmode=disable"
   ```

2. Navigate to the project directory:
   ```bash
   cd c:\dev\DATABASE\Bookings_golang
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The application will:
- Connect to PostgreSQL database using the DATABASE_URL
- Create all necessary tables, enums, and indexes
- Start the HTTP server on port 8080
- Enable CORS for cross-origin requests

4. **Test the API**:
   ```bash
   curl http://localhost:8080/health
   ```

## Testing

1. **Set environment variables**:
   ```bash
   export DATABASE_URL="postgres://bookings_user:your_secure_password@localhost:5432/bookings_db?sslmode=disable"
   ```

2. Run the comprehensive test suite:
   ```bash
   go build -o test_db.exe test_db.go
   .\test_db.exe
   ```

This will test all database operations and API endpoints.

## Dart Client

A complete Dart HTTP client is included for easy integration with Flutter or web applications.

### Setup Dart Client
```bash
cd c:\dev\DATABASE\Bookings_golang
dart pub get
```

### Usage Example
```dart
import 'api_client.dart';

void main() async {
  final apiClient = ApiClient();

  // Get all clinics
  List<Map<String, dynamic>> clinics = await apiClient.getClinics();

  // Create a new patient
  Map<String, dynamic> newPatient = {
    'first_name': 'John',
    'last_name': 'Doe',
    'email': 'john.doe@email.com',
    'phone': '+1234567890',
    'date_of_birth': '1990-05-15',
    'medical_record_number': 'MRN001'
  };
  Map<String, dynamic> patient = await apiClient.createPatient(newPatient);
}
```

See `api_client.dart` for complete documentation and examples.

## Dependencies

- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver for Go
- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework for Go
- [http](https://pub.dev/packages/http) - HTTP client for Dart

## API Endpoints

The application provides REST API endpoints for all major operations:

### Health Check
- `GET /health` - Check if the API is running

### Clinics
- `GET /api/clinics` - Get all clinics
- `GET /api/clinics/:id` - Get clinic by ID
- `POST /api/clinics` - Create a new clinic
- `PUT /api/clinics/:id` - Update clinic
- `DELETE /api/clinics/:id` - Delete clinic

### Patients
- `GET /api/patients` - Get all patients
- `GET /api/patients/:id` - Get patient by ID
- `POST /api/patients` - Create a new patient
- `PUT /api/patients/:id` - Update patient
- `DELETE /api/patients/:id` - Delete patient

### Employees
- `GET /api/employees` - Get all employees
- `GET /api/employees/:id` - Get employee by ID
- `POST /api/employees` - Create a new employee
- `PUT /api/employees/:id` - Update employee
- `DELETE /api/employees/:id` - Delete employee

### Services
- `GET /api/services` - Get all services
- `GET /api/services/:id` - Get service by ID
- `POST /api/services` - Create a new service
- `PUT /api/services/:id` - Update service
- `DELETE /api/services/:id` - Delete service

### Appointments
- `GET /api/appointments` - Get all appointments
- `GET /api/appointments/:id` - Get appointment by ID
- `POST /api/appointments` - Create a new appointment
- `PUT /api/appointments/:id` - Update appointment
- `DELETE /api/appointments/:id` - Delete appointment

### Waiting List
- `GET /api/waiting-list` - Get all waiting list items
- `GET /api/waiting-list/:id` - Get waiting list item by ID
- `POST /api/waiting-list` - Create a new waiting list item
- `PUT /api/waiting-list/:id` - Update waiting list item
- `DELETE /api/waiting-list/:id` - Delete waiting list item

## Sample API Requests

### Create a Clinic
```bash
curl -X POST http://localhost:8080/api/clinics \
  -H "Content-Type: application/json" \
  -d '{
    "name": "City Medical Center",
    "address": "123 Main St, Colombo",
    "phone": "+94-11-1234567",
    "email": "info@citymedical.lk",
    "active": true
  }'
```

### Create a Patient
```bash
curl -X POST http://localhost:8080/api/patients \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@email.com",
    "phone": "+94-77-1234567",
    "date_of_birth": "1985-05-15",
    "medical_record_number": "MRN001",
    "insurance_provider": "ABC Insurance",
    "insurance_id": "INS123456",
    "emergency_contact_name": "Jane Doe",
    "emergency_contact_phone": "+94-77-7654321",
    "active": true
  }'
```

### Create an Employee (Doctor)
```bash
curl -X POST http://localhost:8080/api/employees \
  -H "Content-Type: application/json" \
  -d '{
    "clinic_id": 1,
    "first_name": "Dr. Sarah",
    "last_name": "Williams",
    "email": "sarah.williams@clinic.com",
    "phone": "+94-77-9876543",
    "license_number": "MD123456",
    "specialty": "cardiology",
    "timezone": "Asia/Colombo",
    "active": true
  }'
```

### Create a Service
```bash
curl -X POST http://localhost:8080/api/services \
  -H "Content-Type: application/json" \
  -d '{
    "name": "General Consultation",
    "description": "Comprehensive health check and consultation",
    "duration_minutes": 30,
    "price": 150.00,
    "specialty_required": "general_medicine",
    "active": true
  }'
```

### Create an Appointment
```bash
curl -X POST http://localhost:8080/api/appointments \
  -H "Content-Type: application/json" \
  -d '{
    "patient_id": 1,
    "employee_id": 1,
    "service_id": 1,
    "clinic_id": 1,
    "start_datetime": "2025-10-25T10:00:00Z",
    "end_datetime": "2025-10-25T10:30:00Z",
    "status": "SCHEDULED",
    "appointment_type": "INITIAL_CONSULTATION",
    "notes": "Patient reports chest pain",
    "payment_status": "PENDING",
    "payment_amount": 150.00
  }'
```

### Add to Waiting List
```bash
curl -X POST http://localhost:8080/api/waiting-list \
  -H "Content-Type: application/json" \
  -d '{
    "patient_id": 2,
    "service_id": 1,
    "preferred_employee_id": 1,
    "requested_date": "2025-10-26T09:00:00Z",
    "urgency_level": "HIGH",
    "notes": "Patient needs urgent cardiology consultation",
    "status": "ACTIVE"
  }'
```

### Get All Appointments
```bash
curl http://localhost:8080/api/appointments
```

### Update Appointment Status
```bash
curl -X PUT http://localhost:8080/api/appointments/1 \
  -H "Content-Type: application/json" \
  -d '{
    "patient_id": 1,
    "employee_id": 1,
    "service_id": 1,
    "clinic_id": 1,
    "start_datetime": "2025-10-25T10:00:00Z",
    "end_datetime": "2025-10-25T10:30:00Z",
    "status": "CONFIRMED",
    "appointment_type": "INITIAL_CONSULTATION",
    "notes": "Patient reports chest pain",
    "payment_status": "PAID",
    "payment_amount": 150.00
  }'
```

## Project Structure

```
bookings_golang/
├── main.go                 # Application entry point and API routes
├── database/
│   └── database.go         # Database connection and CRUD operations
├── models/
│   └── models.go           # Data structures and models
├── handlers/
│   └── handlers.go         # HTTP request handlers
├── test_db.go              # Comprehensive testing suite
├── api_client.dart         # Dart HTTP client for API integration
├── pubspec.yaml           # Dart project dependencies
├── go.mod                 # Go module dependencies
└── README.md              # This file
```

## Security Considerations

- **Prepared Statements**: All database queries use parameterized queries (pgx)
- **Input Validation**: Implement proper validation for all API inputs
- **Authentication**: Add JWT or OAuth2 authentication for production use
- **Authorization**: Implement role-based access control (admin, doctor, patient)
- **Data Encryption**: Encrypt sensitive patient data at rest and in transit
- **HTTPS**: Use HTTPS for all API communications
- **Rate Limiting**: Implement rate limiting to prevent abuse
- **Audit Logging**: Log all sensitive operations for compliance
- **Regular Updates**: Keep dependencies updated and perform security audits

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `go build -o test_db.exe test_db.go && .\test_db.exe`
5. Submit a pull request

## License

This project is licensed under the **GNU Affero General Public License v3.0 (AGPL-3.0)** - see the [LICENSE](LICENSE) file for details.

**Important**: This license ensures that:
- The software remains free and open source forever
- Any modifications or derivative works must also be licensed under AGPL-3.0
- Network use (like web applications) requires source code availability
- Commercial use is allowed, but modifications must be shared

For more information about AGPL-3.0, visit [https://www.gnu.org/licenses/agpl-3.0.html](https://www.gnu.org/licenses/agpl-3.0.html)