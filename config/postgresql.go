package config

import (
	"gopportunities/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgresql")
	dsn := "host=localhost user=postgres password=postgres dbname=gopportunities port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("PostgreSQL opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("PostgreSQL automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
