package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection(cfg *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
