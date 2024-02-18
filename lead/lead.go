package lead

import (
	"github.com/Anjasfedo/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn

	var leads []Lead

	db.Find(&leads)

	c.JSON(leads)

}

func GetLeadById(c *fiber.Ctx) {
	id := c.Params("id")

	db := database.DBConn

	var lead Lead

	db.Find(&lead, id)

	c.JSON(lead)
}

func CreateLead(c *fiber.Ctx) {
	db := database.DBConn

	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(lead)

	c.JSON(lead)
}

func DeleteLeadById(c *fiber.Ctx) {
	id := c.Params("id")

	db := database.DBConn

	var lead Lead

	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("Failed to find Lead with ID")
		return
	}

	db.Delete(&lead)

	c.Send("Lead deleted")
}
