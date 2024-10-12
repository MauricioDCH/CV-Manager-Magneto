package http

import (
	"CV_MANAGER/pkg/endpoint"
	"CV_MANAGER/pkg/service"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(svc service.CVService) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/create-cv", endpoint.MakeCreateCVHandler(svc)).Methods("POST")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return corsHandler(r)
}
