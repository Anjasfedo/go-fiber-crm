package lead

import "github.com/jinzhu/gorm"

type lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}
