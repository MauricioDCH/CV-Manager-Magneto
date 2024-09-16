package endpoint

import (
	"CV_MANAGER/pkg/service"
	"encoding/json"
	"net/http"
)

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Error string `json:"error,omitempty"`
}

func MakeLoginHandler(svc service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := svc.LoginUser(req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		resp := LoginUserResponse{
			ID:    user.ID,
			Name:  user.Nombre,
			Email: user.Correo,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
			return
		}
	}
}
