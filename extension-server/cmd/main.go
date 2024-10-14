package main

import (
	"extension-server/config"
	"extension-server/models"
	"extension-server/pkg/db"

	//"extension-server/pkg/gemini"
	"extension-server/pkg/service"
	"extension-server/pkg/transport"
	"log"
	"net/http"
	"os"
)

func main() {
	// Cargar la configuración desde el archivo de configuración
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error cargando la configuración: %v", err)
	}

	// Intentar conectarse a la base de datos PostgreSQL
	dbConn, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	// Ejecutar migraciones para crear la tabla de usuarios
	if err := dbConn.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error ejecutando migraciones: %v", err)
	}
	log.Println("Migraciones ejecutadas con éxito")

	// Asegurarse de cerrar la conexión a la base de datos al finalizar
	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatalf("Error obteniendo la instancia de la base de datos SQL: %v", err)
	}
	defer sqlDB.Close()
	// Si la conexión a la base de datos fue exitosa, cerrarla al finalizar las migraciones.
	log.Println("Conexión a la base de datos cerrada exitosa")

	// Inicializar el servicio con la conexión a la base de datos
	svc, err := service.NewService(sqlDB)
	if err != nil {
		log.Fatalf("Error creando el servicio: %v", err)
	}

	// Crear un nuevo enrutador para manejar las solicitudes
	r := transport.NewRouter(svc)

	// Configurar el puerto del servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Valor por defecto si no se encuentra la variable de entorno
	}

	// Iniciar el servidor HTTP
	log.Printf("Servidor escuchando en el puerto %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
