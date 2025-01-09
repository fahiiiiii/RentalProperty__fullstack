package conf

import (
    "fmt"
    "log"
    // "os"
    "property-listing/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

// func InitDB() {
//     dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
//         os.Getenv("DB_HOST"),
//         os.Getenv("DB_USER"),
//         os.Getenv("DB_PASSWORD"),
//         os.Getenv("DB_NAME"),
//         os.Getenv("DB_PORT"),
//     )
func InitDB() {
    // Hard-coded database connection parameters
    DB_HOST := "localhost" // Change this to your DB host if needed
    DB_USER  := "fahimah"
    DB_PASSWORD := "fahimah123"
    DB_NAME := "rental_db"
    DB_PORT := "5432"

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        DB_HOST,
        DB_USER ,
        DB_PASSWORD,
        DB_NAME,
        DB_PORT,
    )


    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto-migrate the Location table
    err = db.AutoMigrate(&models.Location{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    DB = db
}