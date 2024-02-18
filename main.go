package main

import (
	"fmt"

	"github.com/Anjasfedo/go-fiber-crm/database"
	"github.com/Anjasfedo/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setRoutes(app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads)

	app.Get("/api/v1/lead/:id",lead.GetLeadById)

	app.Post("/api/v1/lead/",lead.CreateLeadById)

	app.Delete("/api/v1/lead/:id",lead.DeleteLeadById)
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed connect to database")
	}

	fmt.Println("Connected to database")

	database.DBConn.AutoMigrate(&lead.lead{})

	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()

	initDatabase()

	setRoutes(app)

	app.Listen(3000)

	defer database.DBConn.Close()


}
