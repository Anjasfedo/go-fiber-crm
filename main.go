package main

import (
	"fmt"

	"github.com/Anjasfedo/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setRoutes(app *fiber.App) {
	app.Get(GetLeads)

	app.Get(GetLeadById)

	app.Post(CreateLeadById)

	app.post(DeleteLeadById)
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
