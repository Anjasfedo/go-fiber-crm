package lead

import (
	"github.com/Anjasfedo/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads() {

}

func GetLeadById() {

}

func CreateLead() {

}

func DeleteLeadById() {
	
}
