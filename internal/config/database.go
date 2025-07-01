package config

import (
	"fmt"
	"log"
	"os"

	"go-ecommerce/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Lấy cấu hình từ biến môi trường (Docker hoặc .env)
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "123")
	dbname := getEnv("DB_NAME", "ecommerce")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
		os.Exit(1)
	}

	DB = db
	fmt.Println("✅ Connected to PostgreSQL!")
}

func AutoMigrate() {
	err := DB.AutoMigrate(
		&model.Product{},
		&model.User{},
		&model.Order{},
		&model.OrderItem{},
	)
	if err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}
	fmt.Println("✅ Database migrated!")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
func GetDB() *gorm.DB {
	return DB
}