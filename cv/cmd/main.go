package main

import (
	"log"
	"net/http"
	"os"

	"CV_MANAGER/config"
	"CV_MANAGER/models"
	"CV_MANAGER/pkg/db"
	"CV_MANAGER/pkg/service"
	transportHttp "CV_MANAGER/pkg/transport/http"

	_ "github.com/lib/pq"
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

	if err := dbConn.AutoMigrate(&models.CV{}); err != nil {
		log.Fatalf("Error ejecutando migraciones: %v", err)
	}

	cvService := service.NewCVService(dbConn)

	r := transportHttp.NewRouter(cvService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}

	log.Printf("Servidor escuchando en el puerto %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
