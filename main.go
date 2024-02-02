package main

import (
	"github.com/gofiber/fiber"
)

func setRoutes(app *fiber.App) {
	app.Get(GetLeads)

	app.Get(GetLeadById)

	app.Post(CreateLeadById)

	app.post(DeleteLeadById)
}

func initDatabase() {
	
}

func main() {
	app := fiber.New()

	setRoutes(app)

	app.Listen(3000)
}
