// pkg/gemini/gemini.go
package gemini

import (
	"context"
	"log"

	"extension-server/config" // Importar configuración

	"github.com/google/generative-ai-go/genai" // Dependencias para Google Generative AI
	"google.golang.org/api/option"
)

// Función para conectarse a Gemini
func ConnectToGemini() (*genai.Client, context.Context) {
	// Cargar configuración desde el archivo de configuración
	configs, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	// Crear un contexto
	ctx := context.Background()

	// Clave API de Google Generative AI
	apiKey := configs.GeminiAPIKey
	if apiKey == "" {
		log.Fatal("La variable de entorno GEMINI_API_KEY no está configurada")
	}

	// Crear el cliente con la clave API
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal("Error al crear el cliente de Gemini:", err)
	}

	// Retornar el cliente y el contexto para su uso posterior
	return client, ctx
}
