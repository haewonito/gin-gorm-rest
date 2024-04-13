package config
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/haewonito/gin-gorm-rest/models"

  )

var DB *gorm.DB
func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Album{}, &models.Song{})
	DB = db
}