// pkg/gemini/gemini.go
package gemini

import (
	"context"
	"log"

	"extension-server/config"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ConnectToGemini() (*genai.Client, context.Context) {
	configs, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	ctx := context.Background()

	apiKey := configs.GeminiAPIKey
	if apiKey == "" {
		log.Fatal("La variable de entorno GEMINI_API_KEY no está configurada")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal("Error al crear el cliente de Gemini:", err)
	}

	return client, ctx
}
