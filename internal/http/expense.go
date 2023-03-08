package http

import (
	"app/internal/model"
	"app/internal/validate"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (r *Router) HandleGetExpenses(res http.ResponseWriter, req *http.Request) {
	expenses, err := r.expenseRepo.FindAll()
	if err != nil {
		// TODO: write error response
		r.writeServerError(res)
		return
	}

	b, _ := json.Marshal(expenses)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func (r *Router) HandleGetOneExpense(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	expenseId, err := validate.IsUint(id)
	if err != nil {
		r.logger.Printf("Invalid location ID found in url")
		// TODO: write error response
		r.writeServerError(res)
		return
	}

	expense, err := r.expenseRepo.FindById(expenseId)
	if err != nil {
		// TODO: write error response
		r.writeServerError(res)
		return
	}
	b, _ := json.Marshal(expense)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func (r *Router) HandleCreateExpense(res http.ResponseWriter, req *http.Request) {
	amount, err := strconv.ParseFloat(req.FormValue("amount"), 32)
	if err != nil {
		r.logger.Printf("unable to parse amount from expense")
		return
	}
	date, _ := time.Parse("02-01-2006", req.FormValue("date"))
	creation, _ := time.Parse("02-01-2006", req.FormValue("creation"))

	expense, err := r.expenseRepo.Save(model.Expense{
		Type:     req.FormValue("type"),
		UserId:   req.FormValue("userId"),
		Name:     req.FormValue("name"),
		Amount:   float32(amount),
		Status:   req.FormValue("status"),
		Category: req.FormValue("category"),
		Group:    req.FormValue("group"),
		Comment:  req.FormValue("comment"),
		Date:     date,
		Creation: creation,
	})
	b, _ := json.Marshal(expense)
	if err != nil {
		// TODO: return decent HTTP response
		r.logger.Printf("unable to marshal location")
		r.writeServerError(res)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	// res.Header().Set("Location", fmt.Sprintf("/api/location/%d", location.Id))
	res.WriteHeader(http.StatusCreated)
	res.Write(b)
}

func (r *Router) HandleDeleteExpense(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	expenseId, err := validate.IsUint(id)
	if err != nil {
		r.logger.Printf("Invalid expense ID found in url")
		// TODO: write error response
		r.writeServerError(res)
		return
	}
	success, err := r.expenseRepo.DeleteExpense(expenseId)
	if err != nil {
		r.logger.Printf("something went wrong when deleting expense")
		// TODO: write error response
		r.writeServerError(res)
		return
	}
	b, _ := json.Marshal(success)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}
