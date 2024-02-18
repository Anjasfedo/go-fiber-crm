package main

import (
	"fmt"

	"github.com/Anjasfedo/go-fiber-crm/database"
	"github.com/Anjasfedo/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)              // Defines a route to access all lead data using the GET method.
	app.Get("/api/v1/lead/:id", lead.GetLeadById)       // Defines a route to access lead data by ID using the GET method.
	app.Post("/api/v1/lead/", lead.CreateLead)          // Defines a route to create new lead data using the POST method.
	app.Delete("/api/v1/lead/:id", lead.DeleteLeadById) // Defines a route to delete lead data by ID using the DELETE method.
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "/database/leads.db")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database") // Prints a message indicating successful connection to the database.

	database.DBConn.AutoMigrate(&lead.Lead{}) // Automatically migrates the Lead model to the database schema.

	fmt.Println("Database migrated") // Prints a message indicating successful migration of the database.
}

func main() {
	app := fiber.New() // Initializes a new Fiber application.

	initDatabase() // Initializes the database connection and migrates the database schema.

	setRoutes(app) // Sets up routes for the Fiber application.

	app.Listen(3000) // Starts the Fiber application to listen on port 3000 for incoming requests.

	defer database.DBConn.Close() // Defer closing the database connection until the main function exits.
}
