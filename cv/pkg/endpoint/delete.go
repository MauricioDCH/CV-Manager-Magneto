// endpoint/cv/delete.go
package endpoint

import (
	"CV_MANAGER/pkg/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func MakeDeleteCVHandler(svc service.CVService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID de la hoja de vida de los parámetros de la URL
		vars := mux.Vars(r)
		cvIDStr := vars["id"]
		cvID, err := strconv.ParseUint(cvIDStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid CV ID", http.StatusBadRequest)
			return
		}

		// Intentar eliminar la hoja de vida
		err = svc.DeleteCV(uint(cvID))
		if err != nil {
			http.Error(w, "Error deleting CV", http.StatusInternalServerError)
			return
		}

		// Confirmación de eliminación exitosa
		w.WriteHeader(http.StatusNoContent)
	}
}
