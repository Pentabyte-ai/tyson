package database

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

// ConnectDB connects to the MySQL database and returns the GORM DB object.
func ConnectDB() *gorm.DB {
    // Load environment variables from .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Retrieve environment variables
    dbUser := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")

    // Check if any of the environment variables are empty
    if dbUser == "" || password == "" || dbName == "" || host == "" || port == "" {
        log.Fatalf("Database configuration variables are missing. Check your .env file.")
    }

    // Construct DSN (Data Source Name)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, password, host, port, dbName)

    // Connect to the database using GORM v2
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // Enable detailed logging
    })
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Get the underlying database object to set connection pool parameters
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Error getting DB from GORM: %v", err)
    }

    // Set connection pool settings
    sqlDB.SetMaxOpenConns(25)           // Maximum number of open connections
    sqlDB.SetMaxIdleConns(25)           // Maximum number of idle connections
    sqlDB.SetConnMaxLifetime(5 * time.Minute) // Connection max lifetime

    // Verify the connection with a ping
    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

    log.Println("Database connection established successfully.")
    return db
}
