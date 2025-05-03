package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

// InitDB - инициализация базы данных и миграций
func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment")
	}
	databaseName := "postgres"
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	sslmode := "disable"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)
	fmt.Println(dbUrl)

	// Открытие соединения с базой данных
	sqlDB, err := sql.Open(databaseName, dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Настройка миграций
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Получаем текущую рабочую директорию
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working dir: %v", err)
	}
	log.Println("Current working dir:", cwd)

	// Путь к миграциям
	migrationsPath := filepath.Join(cwd, "internal", "db", "migrations")
	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)
	log.Println("Migrations path:", migrationsURL)

	// Запуск миграций
	m, err := migrate.NewWithDatabaseInstance(migrationsURL, databaseName, driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	// Подключение к базе данных через GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = gormDB
}

// GetDB - получение подключения к базе данных
func GetDB() *gorm.DB {
	return DB
}
