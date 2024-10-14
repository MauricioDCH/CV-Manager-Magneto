package transport

import (
	"extension-server/pkg/endpoint"
	"extension-server/pkg/service"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(svc service.Service) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/endpoint", endpoint.HandlePostRequest(svc)).Methods("POST")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return corsHandler(r)
}
