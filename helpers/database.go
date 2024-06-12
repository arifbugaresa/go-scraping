package helpers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

var DBConnection *gorm.DB

func InitDatabase() {
	//todo use env file (for test only)
	dsn := "root:password@tcp(localhost:3306)/dans_golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&User{})

	DBConnection = db

	if err := createSampleUsers(); err != nil {
		log.Fatalf("failed to create sample users: %v", err)
	}
}

func createSampleUsers() error {
	// Check if sample user exists
	var tempUser User
	if err := DBConnection.First(&tempUser, "username = ?", "user").Error; err == nil {
		return nil
	}

	// Create sample user
	return DBConnection.Create(User{Username: "user", Password: "password"}).Error
}
