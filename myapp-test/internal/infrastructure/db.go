package infrastructure

import (
	"fmt"
	"os"

	
	"gorm.io/driver/postgres"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	
)


func NewDB() (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch os.Getenv("DB_TYPE") {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		dialector = mysql.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(os.Getenv("DB_NAME") + ".db")
	default: // postgres
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"))
		dialector = postgres.Open(dsn)
	}

	return gorm.Open(dialector, &gorm.Config{})
}

