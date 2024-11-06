package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var intance_db_global *gorm.DB

func GetConnection() *gorm.DB {
	return intance_db_global
}

func OpenConnection(config *Config) (*gorm.DB, error) {
	if intance_db_global != nil {
		return intance_db_global, nil
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Host,
		config.Database.Port,
		config.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	intance_db_global = db
	return db, nil
}
