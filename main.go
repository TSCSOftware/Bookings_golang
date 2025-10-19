// Medical Appointment Booking System
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
	"log"

	"bookings/database"
	"bookings/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	database.InitDB()
	defer database.CloseDB()

	// Create database tables if they don't exist
	if err := database.CreateTables(); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // In production, specify your frontend URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// API Routes
	api := r.Group("/api")
	{
		// Clinic routes
		clinics := api.Group("/clinics")
		{
			clinics.GET("", handlers.GetClinics)
			clinics.GET("/:id", handlers.GetClinic)
			clinics.POST("", handlers.CreateClinic)
			clinics.PUT("/:id", handlers.UpdateClinic)
			clinics.DELETE("/:id", handlers.DeleteClinic)
		}

		// Patient routes
		patients := api.Group("/patients")
		{
			patients.GET("", handlers.GetPatients)
			patients.GET("/:id", handlers.GetPatient)
			patients.POST("", handlers.CreatePatient)
			patients.PUT("/:id", handlers.UpdatePatient)
			patients.DELETE("/:id", handlers.DeletePatient)
		}

		// Employee routes
		employees := api.Group("/employees")
		{
			employees.GET("", handlers.GetEmployees)
			employees.GET("/:id", handlers.GetEmployee)
			employees.POST("", handlers.CreateEmployee)
			employees.PUT("/:id", handlers.UpdateEmployee)
			employees.DELETE("/:id", handlers.DeleteEmployee)
		}

		// Service routes
		services := api.Group("/services")
		{
			services.GET("", handlers.GetServices)
			services.GET("/:id", handlers.GetService)
			services.POST("", handlers.CreateService)
			services.PUT("/:id", handlers.UpdateService)
			services.DELETE("/:id", handlers.DeleteService)
		}

		// Appointment routes
		appointments := api.Group("/appointments")
		{
			appointments.GET("", handlers.GetAppointments)
			appointments.GET("/:id", handlers.GetAppointment)
			appointments.POST("", handlers.CreateAppointment)
			appointments.PUT("/:id", handlers.UpdateAppointment)
			appointments.DELETE("/:id", handlers.DeleteAppointment)
		}

		// Waiting list routes
		waitingList := api.Group("/waiting-list")
		{
			waitingList.GET("", handlers.GetWaitingList)
			waitingList.GET("/:id", handlers.GetWaitingListItem)
			waitingList.POST("", handlers.CreateWaitingListItem)
			waitingList.PUT("/:id", handlers.UpdateWaitingListItem)
			waitingList.DELETE("/:id", handlers.DeleteWaitingListItem)
		}
	}

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Medical Appointment Booking API is running",
		})
	})

	log.Println("Server starting on port 8080...")
	log.Fatal(r.Run(":8080"))
}
