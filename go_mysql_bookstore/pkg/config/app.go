package config

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect establishes a connection to the MySQL database
func Connect() {
	dsn := "shubham:Shubham@123/bookstore?charset=utf8&parseTime=True&loc=Local"
	
	var err error
	
	// Try connecting with retry logic
	for i := 0; i < 3; i++ {
		db, err = gorm.Open("mysql", dsn)
		if err == nil {
			break
		}
		
		log.Printf("Failed to connect to database (attempt %d/3): %v", i+1, err)
		time.Sleep(time.Second * 2)
	}
	
	if err != nil {
		log.Fatalf("Failed to connect to database after 3 attempts: %v", err)
	}
	
	// Configure the connection pool
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	
	// Enable GORM's log mode in development (disable in production)
	db.LogMode(true)
	
	log.Println("Connected to MySQL database successfully")
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}

// Close closes the database connection
func Close() {
	if db != nil {
		db.Close()
		log.Println("Database connection closed")
	}
}