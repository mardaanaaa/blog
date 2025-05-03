package db

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"rest-project/internal/models"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	// migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	// "github.com/golang-migrate/migrate/v4"
)

var DB *gorm.DB

// InitDB - инициализация базы данных и миграций
func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment")
	}

	// Считываем переменные окружения
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	sslmode := "disable"

	// Строка подключения
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)
	fmt.Println(dbUrl)

	// Подключение к БД через GORM
	gormDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database using GORM:", err)
	}
	DB = gormDB

	// Автоматическая миграция моделей
	err = DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// Путь к миграциям (если будешь использовать migrate)
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working dir: %v", err)
	}
	log.Println("Current working dir:", cwd)
	migrationsPath := filepath.Join(cwd, "internal", "db", "migrations")
	log.Println("Migrations path:", fmt.Sprintf("file://%s", migrationsPath))

	// Ниже можно раскомментировать, если будешь использовать golang-migrate
	/*
		sqlDB, err := sql.Open("postgres", dbUrl)
		if err != nil {
			log.Fatal(err)
		}
		driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
		if err != nil {
			log.Fatal(err)
		}
		m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", migrationsPath), "postgres", driver)
		if err != nil {
			log.Fatal(err)
		}
		if err := m.Up(); err != nil && err.Error() != "no change" {
			log.Fatal(err)
		}
	*/
}

// GetDB - возвращает подключение к БД
func GetDB() *gorm.DB {
	return DB
}
