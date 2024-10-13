package endpoint

import (
	"CV_MANAGER/pkg/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func MakeListCVsHandler(svc service.CVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el user_id de los parámetros de la URL
		vars := mux.Vars(r)
		userIDStr := vars["user_id"]
		userID, err := strconv.ParseUint(userIDStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusBadRequest)
			return
		}

		// Listar las CVs del usuario
		cvs, err := svc.ListCVsByUser(uint(userID))
		if err != nil {
			http.Error(w, "Error fetching CVs", http.StatusInternalServerError)
			return
		}

		// Responder con la lista de CVs
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cvs)
	}
}