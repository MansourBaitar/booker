package http

import (
	"app/internal/model"
	"app/internal/validate"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func (r *Router) HandleGetHolidays(res http.ResponseWriter, req *http.Request) {
	holidays, err := r.holidayRepo.FindAll()
	if err != nil {
		// TODO: write error response
		r.writeServerError(res)
		return
	}

	b, err := json.Marshal(holidays)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func (r *Router) HandleGetOneHoliday(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	holidayId, err := validate.IsUint(id)
	if err != nil {
		r.logger.Printf("Invalid holiday ID found in url")
		// TODO: write error response
		r.writeServerError(res)
		return
	}

	holiday, err := r.holidayRepo.FindById(holidayId)
	if err != nil {
		// TODO: write error response
		r.writeServerError(res)
		return
	}
	b, err := json.Marshal(holiday)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func (r *Router) HandleCreateHoliday(res http.ResponseWriter, req *http.Request) {

	fromDate, _ := time.Parse("02-01-2006", req.FormValue("fromDate"))
	toDate, _ := time.Parse("02-01-2006", req.FormValue("toDate"))
	creation, _ := time.Parse("02-01-2006", req.FormValue("creation"))

	holiday, err := r.holidayRepo.Save(model.Holiday{
		Type:     req.FormValue("type"),
		UserId:   req.FormValue("userId"),
		Name:     req.FormValue("name"),
		Status:   req.FormValue("status"),
		FromDate: fromDate,
		ToDate:   toDate,
		Creation: creation,
	})
	b, err := json.Marshal(holiday)
	if err != nil {
		// TODO: return decent HTTP response
		r.logger.Printf("unable to marshal holiday")
		r.writeServerError(res)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(b)
}

func (r *Router) HandleDeleteHoliday(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	holidayId, err := validate.IsUint(id)
	if err != nil {
		r.logger.Printf("Invalid holiday ID found in url")
		// TODO: write error response
		r.writeServerError(res)
		return
	}
	success, err := r.holidayRepo.DeleteHoliday(holidayId)
	if err != nil {
		r.logger.Printf("something went wrong when deleting holiday")
		// TODO: write error response
		r.writeServerError(res)
		return
	}
	b, err := json.Marshal(success)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}
