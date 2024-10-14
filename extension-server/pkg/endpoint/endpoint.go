package endpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"extension-server/pkg/service"
)

// Estructura que representa un campo de entrada
type InputField struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"` // Agregar el campo Value
}

// Estructura que representa el JSON completo recibido
type RequestData struct {
	Inputs []InputField `json:"inputs"`
	Email  string       `json:"email"`
}

// Estructura que representa la respuesta con los inputs modificados
type ResponseData struct {
	Inputs []InputField `json:"inputs"`
}

// Estructura que representa la respuesta de la IA
type IAResponse struct {
	Inputs []InputField `json:"inputs"`
}

// Función que maneja las solicitudes POST y devuelve la respuesta generada por la IA
func HandlePostRequest(svc service.Service) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		// Leer el cuerpo de la solicitud
		bodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		// Decodificar el JSON recibido
		var requestData service.RequestData // Usar RequestData del paquete service
		err = json.Unmarshal(bodyBytes, &requestData)
		if err != nil {
			http.Error(responseWriter, "Error al decodificar el JSON", http.StatusBadRequest)
			return
		}

		// Validar que el correo electrónico no esté vacío
		if requestData.Email == "" {
			http.Error(responseWriter, "El campo 'email' no puede estar vacío", http.StatusBadRequest)
			return
		}

		// Llamar al servicio para obtener los datos del CV usando el email
		cvsData, err := svc.GetCvsByEmail(requestData.Email)
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Error en el servicio: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// Crear un mapa con los valores del CV para facilitar la asignación
		cvsValues := map[string]string{
			"name":       cvsData.Name,
			"last_name":  cvsData.LastName,
			"email":      cvsData.Email,
			"phone":      cvsData.Phone,
			"experience": cvsData.Experience,
			"skills":     cvsData.Skills,
			"languages":  cvsData.Languages,
			"education":  cvsData.Education,
		}

		// Recorrer los inputs y añadir el valor correspondiente desde cvsData
		for i, input := range requestData.Inputs {
			if val, ok := cvsValues[input.Name]; ok {
				requestData.Inputs[i].Value = val
			}
		}

		// Generar la respuesta de IA utilizando los datos del CV
		iaResponse, err := svc.GeminiQuery(cvsData, requestData) // Llama a GeminiQuery
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Error al consultar la IA: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// Decodificar la respuesta de la IA
		var iaResp service.IAResponse                     // Usar IAResponse del paquete service
		err = json.Unmarshal([]byte(iaResponse), &iaResp) // Deserializar la respuesta de la IA
		if err != nil {
			http.Error(responseWriter, "Error al procesar la respuesta de la IA", http.StatusInternalServerError)
			return
		}

		// Generar la respuesta final en JSON usando la respuesta de la IA
		response, err := svc.GenerateResponse(iaResp) // Cambiado para usar IAResponse
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Error al generar la respuesta: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		// Configurar el tipo de contenido como JSON
		responseWriter.Header().Set("Content-Type", "application/json")

		// Escribir la respuesta generada por la IA
		_, err = responseWriter.Write(response)
		if err != nil {
			http.Error(responseWriter, "Error al escribir la respuesta JSON", http.StatusInternalServerError)
			return
		}
	}
}
