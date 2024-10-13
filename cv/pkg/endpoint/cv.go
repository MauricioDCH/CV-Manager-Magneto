package endpoint

import (
    "CV_MANAGER/models"
    "CV_MANAGER/pkg/service"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type UpdateCVRequest struct {
    Name       string `json:"name"`
    LastName   string `json:"last_name"`
    Email      string `json:"email"`
    Phone      string `json:"phone"`
    Experience string `json:"experience"`
    Skills     string `json:"skills"`
    Languages  string `json:"languages"`
    Education  string `json:"education"`
}

func MakeUpdateCVHandler(svc service.CVService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        cvID, err := strconv.Atoi(vars["id"])
        if err != nil {
            http.Error(w, "ID inválido", http.StatusBadRequest)
            return
        }

        var req UpdateCVRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Cuerpo de solicitud inválido", http.StatusBadRequest)
            return
        }

        updatedCV := models.CV{
            Name:       req.Name,
            LastName:   req.LastName,
            Email:      req.Email,
            Phone:      req.Phone,
            Experience: req.Experience,
            Skills:     req.Skills,
            Languages:  req.Languages,
            Education:  req.Education,
        }

        cv, err := svc.UpdateCV(uint(cvID), updatedCV)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(cv)
    }
}
