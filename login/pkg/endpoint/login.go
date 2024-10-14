package endpoint

import (
	jwtService "CV_MANAGER/pkg/jwt"
	"CV_MANAGER/pkg/service"
	"encoding/json"
	"net/http"
	"time"
)

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
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

		// Generate JWT token for the user
		token, err := jwtService.GenerateJWT(int(user.ID))
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Option 1: Set JWT token in a cookie
		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true, // Set to true in production for HTTPS
		}
		http.SetCookie(w, cookie)

		resp := LoginUserResponse{
			ID:    user.ID,
			Name:  user.Nombre,
			Email: user.Correo,
			Token: token,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
			return
		}
	}
}
