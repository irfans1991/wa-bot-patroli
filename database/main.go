package database

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Users struct {
// 	Id                     int       `gorm:"primaryKey"`
// 	Name                   string    `gorm:"size:255;not null"`
// 	Username               string    `gorm:"size:255;not null"`
// 	Email                  string    `gorm:"size:255;not null;unique"`
// 	Password               string    `gorm:"not null"`
// 	Created_at, Updated_at time.Time `gorm:"not null"`
// }

// func main() {
// 	// Load .env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	// Get environment variables
// 	dbUser := os.Getenv("DB_USER")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbName := os.Getenv("DB_NAME")

// 	// Connection string
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbName, dbPort)

// 	// Open connection to the database
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Error connecting to the database: %v", err)
// 	}

// 	fmt.Println("Successfully connected to the database")

// 	// Example of using GORM
// 	var users []Users
// 	db.Find(&users)
// 	// db.Raw("SELECT * FROM users").Scan(&result)

// 	for _, user := range users {
// 		fmt.Println(user)
// 	}
// 	fmt.Println("Result of 'SELECT 1':", users)
// }
