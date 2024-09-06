package utils

import (
	"backend/api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenGormDB abre la conexión con la base de datos usando GORM y la devuelve
func OpenGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBURL()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
