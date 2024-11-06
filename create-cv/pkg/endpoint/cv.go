package endpoint

import (
	"CV_MANAGER/pkg/service"
	"encoding/json"
	"net/http"
)

type CreateCVRequest struct {
	Title      string `json:"title"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Experience string `json:"experience"`
	Skills     string `json:"skills"`
	Languages  string `json:"languages"`
	Education  string `json:"education"`
	UserID     uint   `json:"user_id"`
}

type CreateCVResponse struct {
	ID    uint   `json:"id"`
	Error string `json:"error,omitempty"`
}

func MakeCreateCVHandler(svc service.CVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateCVRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		cv, err := svc.CreateCV(req.Title, req.Name, req.LastName, req.Email, req.Phone, req.Experience, req.Skills, req.Languages, req.Education, req.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := CreateCVResponse{ID: cv.ID}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
		}
	}
}
