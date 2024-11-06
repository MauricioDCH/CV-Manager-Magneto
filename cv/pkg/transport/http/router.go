package http

import (
	"CV_MANAGER/pkg/endpoint"
	"CV_MANAGER/pkg/service"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(cvSvc service.CVService) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/cv/{id}", endpoint.MakeUpdateCVHandler(cvSvc)).Methods("PUT")
	r.HandleFunc("/cv/user/{user_id}", endpoint.MakeListCVsHandler(cvSvc)).Methods("GET")
	r.HandleFunc("/cv/{id}", endpoint.MakeDeleteCVHandler(cvSvc)).Methods("DELETE")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return corsHandler(r)
}
