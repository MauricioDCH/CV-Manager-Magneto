package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"extension-server/pkg/gemini" // Importar el paquete gemini

	"github.com/google/generative-ai-go/genai"
)

// Interface del servicio que define los métodos
type Service interface {
	GetCvsByEmail(email string) (CvsData, error)
	GeminiQuery(cvsData CvsData, requestData RequestData) (string, error)
	GenerateResponse(iaResponse IAResponse) ([]byte, error)
}

// CvsData representa la estructura de datos que quieres devolver
type CvsData struct {
	UserID     int    `json:"user_id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Experience string `json:"experience"`
	Skills     string `json:"skills"`
	Languages  string `json:"languages"`
	Education  string `json:"education"`
}

// Estructura que representa un campo de entrada
type InputField struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"` // Agregar el campo Value
}

// Estructura para la petición
type RequestData struct {
	Inputs []InputField `json:"inputs"`
	Email  string       `json:"email"`
}

// Estructura para la respuesta
type ResponseField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ResponseToExtension struct {
	ResponseFromServer []ResponseField `json:"response_from_server"`
}

// Estructura para la respuesta de la IA
type IAResponse struct {
	Inputs []InputField `json:"inputs"`
}

// Implementación del servicio
type service struct {
	db *sql.DB
}

// NewService crea una nueva instancia del servicio.
// Recibe una conexión de base de datos como parámetro.
func NewService(db *sql.DB) (Service, error) {
	if db == nil {
		return nil, fmt.Errorf("la conexión a la base de datos no puede ser nula")
	}

	return &service{db: db}, nil
}

// Método para obtener los datos del CV por email
func (s *service) GetCvsByEmail(email string) (CvsData, error) {
	var data CvsData

	query := `
    SELECT u.id AS user_id, c.name, c.last_name, c.email, c.phone, c.experience, c.skills, c.languages, c.education
    FROM users u
    JOIN cvs c ON u.id = c.user_id
    WHERE u.correo = $1;
	`

	err := s.db.QueryRow(query, email).Scan(
		&data.UserID,
		&data.Name,
		&data.LastName,
		&data.Email,
		&data.Phone,
		&data.Experience,
		&data.Skills,
		&data.Languages,
		&data.Education,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return data, fmt.Errorf("no se encontraron registros para el email: %s", email)
		}
		return data, err
	}

	return data, nil
}

// returnResponse procesa la respuesta de Gemini y devuelve el texto completo.
func returnResponse(resp *genai.GenerateContentResponse) string {
	var responseText strings.Builder

	// Recorrer los candidatos y sus contenidos para construir la respuesta
	for _, candidate := range resp.Candidates {
		if candidate.Content != nil { // Asegurarnos de que candidate.Content no sea nil
			for _, part := range candidate.Content.Parts {
				// Comprobamos si `part` es de tipo `genai.Text`
				if textPart, ok := part.(genai.Text); ok {
					// Convertimos el tipo genai.Text a string y lo agregamos al texto de respuesta
					responseText.WriteString(string(textPart))
				} else {
					log.Printf("Tipo inesperado encontrado en la respuesta: %T\n", part)
				}
			}
		} else {
			log.Println("El candidato no contiene contenido válido.")
		}
	}

	// Obtener la respuesta completa como string
	fullResponse := responseText.String()

	// Limpiar la respuesta para eliminar etiquetas de código y caracteres no deseados
	fullResponse = strings.ReplaceAll(fullResponse, "```json\n", "")
	fullResponse = strings.ReplaceAll(fullResponse, "```", "")
	fullResponse = strings.TrimSpace(fullResponse) // Eliminar espacios en blanco al inicio y final
	return fullResponse
}

// GeminiQuery ahora es un método de la estructura `service`
func (s *service) GeminiQuery(cvsData CvsData, requestData RequestData) (string, error) {
	// Conectar a la IA Gemini
	client, ctx := gemini.ConnectToGemini()

	model := client.GenerativeModel("gemini-1.5-flash")

	fmt.Printf("Consulta realizada a la IA GEMINI con los datos del correo: %s\n", requestData.Email)

	// Verificación de valores nulos
	if ctx == nil {
		return "", fmt.Errorf("el contexto es nil")
	}
	if model == nil {
		return "", fmt.Errorf("el modelo de la IA es nil")
	}

	// Imprimir los datos del CV para verificar que se recibieron correctamente en formato JSON
	cvsDataJSON, _ := json.Marshal(cvsData)
	fmt.Printf("\nDatos del CV  --> : %s\n", cvsDataJSON)

	// Imprimir los datos de la petición para verificar que se recibieron correctamente en formato JSON
	requestDataJSON, _ := json.Marshal(requestData)
	fmt.Printf("\nDatos de la petición  --> : %s\n", requestDataJSON)

	// Simulación de respuesta de la IA
	ia_response_example := `
    {
        "inputs": [
            {
                "name": "naasdasdasdme",
                "value": "Juan"
            },
            {
                "name": "last_name",
                "value": "Pérez"
            },
            {
                "name": "email",
                "value": "juan.perez@example.com"
            },
            {
                "name": "education",
                "value": "123456789"
            },
            {
                "name": "experience",
                "value": "5 años en desarrollo de software"
            },
            {
                "name": "skills",
                "value": "Go, Python, SQL"
            },
            {
                "name": "languages",
                "value": "Español, Inglés"
            },
            {
                "name": "education",
                "value": "Ingeniería de Sistemas"
            }
        ]
    }`
	fmt.Printf("\nEjemplo de la respuesta de la IA  --> : %s\n", ia_response_example)

	// Crear el mensaje para enviar a la IA
	message := fmt.Sprintf(
		"Por favor, utiliza la siguiente información para generar una respuesta adecuada en formato JSON:\n\n"+
			"**Datos del CV:**\n%s\n\n"+
			"**Datos de la petición:**\n%s\n\n"+
			"**Instrucciones:**\n"+
			"Necesito que analices muy bien los datos recibidos en la petición para que me des una respuesta correcta de acuerdo con los datos del CV y que "+
			"no perjudiquen al postulante y antes que lo ayuden.\n"+
			"Necesito que la respuesta se base en los datos proporcionados en los JSON que se envían."+
			"Asegúrate de combinar ambos conjuntos de datos de manera coherente y de evitar duplicaciones.\n\n"+
			"**Ejemplo de la respuesta esperada en formato json:**\n%s\n\n"+
			"IMPORTANTE: LAS ENTRADAS DE LA LISTA DEL JSON SÓLO DEBEN SER LOS DATOS DE LA PETICIÓN, SI SÓLO HAY 2 EN LA PETICIÓN, SÓLO SE COMPLETAN LOS DOS.\n\n"+
			"NOTAS ADICIONAL: 1. SI EN LA PETICIÓN SE PIDE UN CAMPO QUE NO ESTÁ EN LOS DATOS DEL CV, SE DEBE RESPONDER CON UN VALOR POR INVENTADO.\n"+
			"2. LOS VALORES INVENTADOS DEBEN SER REDACTADOS EN PRIMERA PERSONA.\n\n"+
			"Asegúrate de que la respuesta sea clara, concisa y en el formato JSON especificado. NO DEBE INCULIR CONTENIDO DE EXPLICACIÓN ADICIONAL,"+
			" SÓLO DEBE SER EL JSON ESPERADO.",
		cvsDataJSON, requestDataJSON, ia_response_example)

	//fmt.Printf("\nMensaje a la IA  --> : %s\n", message)

	// Enviar la consulta a la IA
	responseInputs, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		return "", fmt.Errorf("error en GenerateContent: %v", err)
	}

	// Verificar si la respuesta es nil
	if responseInputs == nil {
		return "", fmt.Errorf("la respuesta de la IA es nil")
	}

	// Solo imprimir toda la estructura completa de la respuesta antes de procesarla
	fmt.Println("--------------------------------------------------------------------------------------------------------------")
	fmt.Println("RESPUESTA COMPLETA DE LA IA")
	iaResponse := returnResponse(responseInputs)
	fmt.Printf("%s\n", iaResponse)
	fmt.Println("--------------------------------------------------------------------------------------------------------------")

	// Devolver un mensaje temporal, solo estamos imprimiendo la respuesta por ahora
	return iaResponse, nil
}

// Método para generar la respuesta JSON en base a la petición
func (s *service) GenerateResponse(iaResponse IAResponse) ([]byte, error) {
	return json.Marshal(iaResponse)
}
