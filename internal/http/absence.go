package http

import (
	"app/internal/model"
	"app/internal/validate"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func (r *Router) HandleGetAbsences(res http.ResponseWriter, req *http.Request) {
	absences, err := r.absenceRepo.FindAll()
	if err != nil {
		// TODO: write error response
		r.writeServerError(res)
		return
	}

	b, err := json.Marshal(absences)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func (r *Router) HandleGetOneAbsence(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	absenceId, err := validate.IsUint(id)
	if err != nil {
		r.logger.Printf("Invalid absence ID found in url")
		// TODO: write error response
		r.writeServerError(res)
		return
	}

	absence, err := r.absenceRepo.FindById(absenceId)
	if err != nil {
		// TODO: write error response
		r.writeServerError(res)
		return
	}
	b, err := json.Marshal(absence)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func (r *Router) HandleCreateAbsence(res http.ResponseWriter, req *http.Request) {

	fromDate, _ := time.Parse("02-01-2006", req.FormValue("fromDate"))
	toDate, _ := time.Parse("02-01-2006", req.FormValue("toDate"))
	creation, _ := time.Parse("02-01-2006", req.FormValue("creation"))

	absence, err := r.absenceRepo.Save(model.Absence{
		Type:     req.FormValue("type"),
		UserId:   req.FormValue("userId"),
		Name:     req.FormValue("name"),
		Status:   req.FormValue("status"),
		FromDate: fromDate,
		ToDate:   toDate,
		Creation: creation,
	})
	b, err := json.Marshal(absence)
	if err != nil {
		// TODO: return decent HTTP response
		r.logger.Printf("unable to marshal absence")
		r.writeServerError(res)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(b)
}

func (r *Router) HandleDeleteAbsence(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	absenceId, err := validate.IsUint(id)
	if err != nil {
		r.logger.Printf("Invalid absence ID found in url")
		// TODO: write error response
		r.writeServerError(res)
		return
	}
	success, err := r.absenceRepo.DeleteAbsence(absenceId)
	if err != nil {
		r.logger.Printf("something went wrong when deleting absence")
		// TODO: write error response
		r.writeServerError(res)
		return
	}
	b, err := json.Marshal(success)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}
