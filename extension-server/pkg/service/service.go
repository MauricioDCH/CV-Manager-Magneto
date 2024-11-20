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
	GetCvsByEmail(Idcv int) (CvsData, error)
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
	Idcv   int          `json:"idcv"`
	//Email  string       `json:"email"`
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

func (s *service) GetCvsByEmail(idcv int) (CvsData, error) {
	var data CvsData

	query := `
    SELECT u.id AS user_id, c.name, c.last_name, c.email, c.phone, c.experience, c.skills, c.languages, c.education
    FROM users u
    JOIN cvs c ON u.id = c.user_id
    WHERE c.id = $1;
	`

	err := s.db.QueryRow(query, idcv).Scan(
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
			return data, fmt.Errorf("no se encontraron registros para el idcv: %d", idcv)
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
	model := client.GenerativeModel("gemini-1.5-flash-8b")
	fmt.Printf("Consulta realizada a la IA GEMINI con los datos del id: %d\n", requestData.Idcv)

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

	ia_bad_response_example := `
	{
		"inputs": [
			{
				"name": "02frstname",
				"value": "Mauricio"
			},
			{
				"name": "04lastname",
				"value": "Correa H"
			},
			{
				"name": "04fullname",
				"value": "Mauricio David Correa H"
			},
			{
				"name": "23cellphon",
				"value": "3101234567"
			},
			{
				"name": "24emailadr",
				"value": "mauriciodch@gmail.com"
			},
			{
				"name": "30_user_id",
				"value": 38
			},
			{
				"name": "60pers_sex",
				"value": ""
			},
			{
				"name": "61pers_ssn",
				"value": null
			},
			{
				"name": "03middle_i",
				"value": null
			}
		]
    }`

	/*
		message := fmt.Sprintf(
			"Por favor, procesa la siguiente información para devolver una respuesta JSON clara y precisa:\n\n"+
				"1. **Datos del CV**:\n%s\n\n"+
				"2. **Datos de la petición**:\n%s\n\n"+
				"3. **Instrucciones específicas**:\n"+
				"- Sólo debes devolver en la respuesta JSON aquellos datos de la petición que tengan valores específicos en el CV.\n"+
				"- **Omitir** cualquier campo en la petición que no tenga un valor en los datos del CV.\n"+
				"- Si hay campos duplicados en la petición, responder sólo con el valor especificado en los datos del CV.\n\n"+
				"Ejemplo esperado en JSON:\n%s\n\n"+
				"NOTA IMPORTANTE: No inventes valores, no incluyas campos sin valor, y no devuelvas explicaciones adicionales. "+
				"La respuesta debe ser sólo el JSON en el formato especificado, sin valores nulos, o vacíos, o espacios"+
				"\n ASEGÚRATE DE QUE LA RESPUESTA SEA SÓLO EN FORMATO JSON ADMISIBLE PARA ENVIARLO DESDE EL SERVIDOR.",
			cvsDataJSON, requestDataJSON, ia_response_example)
	*/

	message := fmt.Sprintf(
		"Por favor, utiliza la siguiente información para generar una respuesta adecuada en formato JSON:\n\n"+
			"**Datos del CV:**\n%s\n\n"+

			"**Datos de la petición:**\n%s\n\n"+

			"**INSTRUCCIONES:**\n"+
			"1. Necesito que analices muy muy bien los datos recibidos en la petición, para que me entregues una respuesta correcta de acuerdo"+
			" con los datos de la hoja de vida, curriculum vitae o postulación a un trabajo.\n"+
			"2. No debe perjudicar al postulante.\n"+
			"3. Necesito que la respuesta se base en SÓLO en los datos proporcionados en los JSON que se envían y que contengan un valor en Datos del CV:**\n%s\n\n."+
			"4. Asegúrate de combinar ambos conjuntos de datos de manera coherente.\n\n"+

			"**EJEMPLO DE LA RESPUESTA ESPERADA EN FORMATO JSON:**\n%s\n\n"+

			"**VERIFICACIONES QUE SE DEBEN HACER ANTES DE GENERAR LA RESPUESTA:**\n\n"+
			"1. Se debe de verificar que si hay valores repetidos en los datos de la petición, se deben responder con el mismo valor.\n"+
			"2. Las entradas de la lista del JSON SÓLO  deben ser los datos de LA PETICIÓN sí y sólo sí tengan un valor en **LOS DATOS DEL CV**.\n"+
			"3. Volver a verificar que: si en la petición hay un campo, el cual no está en **DATOS DEL CV**, entonces ¡SE DEBEN OMITIR DE LA RESPUESTA!"+
			" o sea, que estos campos ¡NOOOOO PUEDE! estar en la respuesta, NO debe de estar en el JSON de la respuesta.\n"+
			"4. La respuesta ¡NO PUEDE! contener entradas en el JSON que tengan valores \"value\" nulos, o sea que NO DEBE DE TENER en el parámetro \"value\" del JSON un valor null.\n"+
			"5. La respuesta ¡NO PUEDE! contener un valores \"value\" que sea un string vacío, o sea que NO DEBE DE TENER en el parámetro \"value\" del JSON un valor \"\".\n"+
			"6. La respuesta ¡NO PUEDE! contener valores inventados, deben ser sólo lo que contenga **DATOS DEL CV**.\n"+
			"7. Asegúrate de que la respuesta sea clara, concisa.\n"+
			"**ESTE ES UN EJEMPLO DE UNA RESPUESTA RESPUESTA QUE ESTÁ EQUIVOCADA: **\n%s\n\n"+
			"8. ¡ASEGÚRATE DE QUE LA RESPUESTA SEA SÓLO EN FORMATO JSON Y CON LAS INDENTACIONES ADECUADAS!",
		cvsDataJSON, requestDataJSON, cvsDataJSON, ia_response_example, ia_bad_response_example)
	responseInputs, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		return "", fmt.Errorf("error en GenerateContent: %v", err)
	}

	if responseInputs == nil {
		return "", fmt.Errorf("la respuesta de la IA es nil")
	}

	fmt.Println("--------------------------------------------------------------------------------------------------------------")
	fmt.Println("PETICIÓN REALIZADA A GEMINI:")
	fmt.Printf("%s\n", message)
	fmt.Println("--------------------------------------------------------------------------------------------------------------")

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
