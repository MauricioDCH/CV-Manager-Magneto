package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"extension-server/pkg/gemini"

	"github.com/google/generative-ai-go/genai"
)

type Service interface {
	GetCvsByEmail(email string) (CvsData, error)
	GeminiQuery(cvsData CvsData, requestData RequestData) (string, error)
	GenerateResponse(iaResponse IAResponse) ([]byte, error)
}

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

type InputField struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

type RequestData struct {
	Inputs []InputField `json:"inputs"`
	Email  string       `json:"email"`
}

type ResponseField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ResponseToExtension struct {
	ResponseFromServer []ResponseField `json:"response_from_server"`
}

type IAResponse struct {
	Inputs []InputField `json:"inputs"`
}

type service struct {
	db *sql.DB
}

func NewService(db *sql.DB) (Service, error) {
	if db == nil {
		return nil, fmt.Errorf("la conexión a la base de datos no puede ser nula")
	}

	return &service{db: db}, nil
}

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

func returnResponse(resp *genai.GenerateContentResponse) string {
	var responseText strings.Builder

	for _, candidate := range resp.Candidates {
		if candidate.Content != nil {
			for _, part := range candidate.Content.Parts {
				if textPart, ok := part.(genai.Text); ok {
					responseText.WriteString(string(textPart))
				} else {
					log.Printf("Tipo inesperado encontrado en la respuesta: %T\n", part)
				}
			}
		} else {
			log.Println("El candidato no contiene contenido válido.")
		}
	}

	fullResponse := responseText.String()

	fullResponse = strings.ReplaceAll(fullResponse, "```json\n", "")
	fullResponse = strings.ReplaceAll(fullResponse, "```", "")
	fullResponse = strings.TrimSpace(fullResponse)
	return fullResponse
}

func (s *service) GeminiQuery(cvsData CvsData, requestData RequestData) (string, error) {
	client, ctx := gemini.ConnectToGemini()
	model := client.GenerativeModel("gemini-1.5-flash")
	fmt.Printf("Consulta realizada a la IA GEMINI con los datos del correo: %s\n", requestData.Email)

	if ctx == nil {
		return "", fmt.Errorf("el contexto es nil")
	}

	if model == nil {
		return "", fmt.Errorf("el modelo de la IA es nil")
	}

	cvsDataJSON, _ := json.Marshal(cvsData)
	requestDataJSON, _ := json.Marshal(requestData)

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

	responseInputs, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		return "", fmt.Errorf("error en GenerateContent: %v", err)
	}

	if responseInputs == nil {
		return "", fmt.Errorf("la respuesta de la IA es nil")
	}

	fmt.Println("--------------------------------------------------------------------------------------------------------------")
	fmt.Println("RESPUESTA DE LA GEMINI:")
	iaResponse := returnResponse(responseInputs)
	fmt.Printf("%s\n", iaResponse)
	fmt.Println("--------------------------------------------------------------------------------------------------------------")

	return iaResponse, nil
}

func (s *service) GenerateResponse(iaResponse IAResponse) ([]byte, error) {
	return json.Marshal(iaResponse)
}
