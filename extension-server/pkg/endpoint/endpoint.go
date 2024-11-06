package endpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"extension-server/pkg/service"
)

type InputField struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

type RequestData struct {
	Inputs []InputField `json:"inputs"`
	Email  string       `json:"email"`
}

type ResponseData struct {
	Inputs []InputField `json:"inputs"`
}

type IAResponse struct {
	Inputs []InputField `json:"inputs"`
}

func HandlePostRequest(svc service.Service) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		bodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		var requestData service.RequestData
		err = json.Unmarshal(bodyBytes, &requestData)
		if err != nil {
			http.Error(responseWriter, "Error al decodificar el JSON", http.StatusBadRequest)
			return
		}

		if requestData.Idcv == 0 {
			http.Error(responseWriter, "El campo 'idcv' no puede ser 0", http.StatusBadRequest)
			return
		}

		cvsData, err := svc.GetCvsByEmail(requestData.Idcv)
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Error en el servicio: %s", err.Error()), http.StatusInternalServerError)
			return
		}

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

		for i, input := range requestData.Inputs {
			if val, ok := cvsValues[input.Name]; ok {
				requestData.Inputs[i].Value = val
			}
		}

		iaResponse, err := svc.GeminiQuery(cvsData, requestData)
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Error al consultar la IA: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		var iaResp service.IAResponse
		err = json.Unmarshal([]byte(iaResponse), &iaResp)
		if err != nil {
			http.Error(responseWriter, "Error al procesar la respuesta de la IA", http.StatusInternalServerError)
			return
		}

		response, err := svc.GenerateResponse(iaResp)
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Error al generar la respuesta: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		responseWriter.Header().Set("Content-Type", "application/json")

		_, err = responseWriter.Write(response)
		if err != nil {
			http.Error(responseWriter, "Error al escribir la respuesta JSON", http.StatusInternalServerError)
			return
		}
	}
}
