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

	// Crear la tabla cv (todos los campos NOT NULL)
	queryCVTable := `
    CREATE TABLE IF NOT EXISTS cvs (
        id SERIAL PRIMARY KEY,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP,
        name TEXT NOT NULL,
        last_name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        phone TEXT NOT NULL,
        experience TEXT NOT NULL,
        skills TEXT NOT NULL,
        languages TEXT NOT NULL,
        education TEXT NOT NULL,
        user_id INTEGER NOT NULL REFERENCES users(id)
    );
    `
	_, err = db.Exec(queryCVTable)
	if err != nil {
		log.Fatalf("Error ejecutando la migración del CV: %v", err)
	}

	log.Println("Migraciones ejecutadas correctamente")
}
