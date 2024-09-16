package db

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) {
	// Crear la tabla users si no existe
	queryUsers := `
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `
	_, err := db.Exec(queryUsers)
	if err != nil {
		log.Fatalf("Error ejecutando la migración de usuarios: %v", err)
	}

	// Crear la tabla encrypted_keys si no existe
	queryEncryptedKeys := `
    CREATE TABLE IF NOT EXISTS encrypted_keys (
        id SERIAL PRIMARY KEY,
        key TEXT NOT NULL
    );
    `
	_, err = db.Exec(queryEncryptedKeys)
	if err != nil {
		log.Fatalf("Error ejecutando la migración de claves cifradas: %v", err)
	}

	log.Println("Migración ejecutada correctamente")
}
