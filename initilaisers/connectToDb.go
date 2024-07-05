package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		panic("failed to connect database")
	}
}