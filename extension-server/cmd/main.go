package main

import (
	"extension-server/config"
	"extension-server/models"
	"extension-server/pkg/db"

	"extension-server/pkg/service"
	"extension-server/pkg/transport"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error cargando la configuración: %v", err)
	}

	dbConn, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	if err := dbConn.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error ejecutando migraciones: %v", err)
	}
	log.Println("Migraciones ejecutadas con éxito")

	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatalf("Error obteniendo la instancia de la base de datos SQL: %v", err)
	}
	defer sqlDB.Close()
	log.Println("Conexión a la base de datos cerrada exitosa")

	svc, err := service.NewService(sqlDB)
	if err != nil {
		log.Fatalf("Error creando el servicio: %v", err)
	}

	r := transport.NewRouter(svc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Servidor escuchando en el puerto %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
