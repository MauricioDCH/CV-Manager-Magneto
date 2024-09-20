package db

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error ejecutando la migración: %v", err)
	}
	log.Println("Migración ejecutada correctamente")
}
