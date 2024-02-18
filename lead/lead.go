package lead

import (
	"github.com/Anjasfedo/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct {
	gorm.Model        // Embeds gorm.Model to inherit fields like ID, CreatedAt, UpdatedAt, and DeletedAt.
	Name       string `json:"name"`    // Defines the Name field of a Lead with JSON tag.
	Company    string `json:"company"` // Defines the Company field of a Lead with JSON tag.
	Email      string `json:"email"`   // Defines the Email field of a Lead with JSON tag.
	Phone      int    `json:"phone"`   // Defines the Phone field of a Lead with JSON tag.
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn

	var leads []Lead

	db.Find(&leads) // Retrieves all leads from the database.

	c.JSON(leads) // Sends the leads as a JSON response.
}

func GetLeadById(c *fiber.Ctx) {
	id := c.Params("id") // Retrieves the ID parameter from the request URL.

	db := database.DBConn

	var lead Lead

	db.Find(&lead, id) // Retrieves the lead from the database based on the provided ID.

	c.JSON(lead) // Sends the lead as a JSON response.
}

func CreateLead(c *fiber.Ctx) {
	db := database.DBConn

	lead := new(Lead) // Initializes a new Lead instance.

	if err := c.BodyParser(lead); err != nil { // Parses the request body and binds it to the lead struct.
		c.Status(503).Send(err) // Sends an error response if parsing fails.
		return
	}

	db.Create(lead) // Creates a new lead record in the database.

	c.JSON(lead) // Sends the newly created lead as a JSON response.
}

func DeleteLeadById(c *fiber.Ctx) {
	id := c.Params("id") // Retrieves the ID parameter from the request URL.

	db := database.DBConn

	var lead Lead

	db.First(&lead, id) // Retrieves the lead from the database based on the provided ID.

	if lead.Name == "" { // Checks if the lead exists.
		c.Status(500).Send("Failed to find Lead with ID") // Sends an error response if the lead is not found.
		return
	}

	db.Delete(&lead) // Deletes the lead record from the database.

	c.Send("Lead deleted") // Sends a success message indicating lead deletion.
}
