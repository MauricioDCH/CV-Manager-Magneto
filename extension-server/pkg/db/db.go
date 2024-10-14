package db

import (
	"extension-server/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {
	var connStr string
	if cfg.PrivateIP != "" {
		// Conexión mediante IP privada
		connStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.PrivateIP, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	} else {
		// Conexión mediante el Cloud SQL Proxy
		connStr = fmt.Sprintf("host=/cloudsql/%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.InstanceConnectionName, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	}

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Probar la conexión
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting SQL DB: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Conexión exitosa a base de datos")
	return db, nil
}
