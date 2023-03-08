package http

import (
	"app/db"
	"app/internal"
	"app/internal/model"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Router struct {
	files  *embed.FS
	logger *log.Logger

	userRepo    model.UserRepository
	expenseRepo model.ExpenseRepository
	holidayRepo model.HolidayRepository
	absenceRepo model.AbsenceRepository
}

// func NewHttpRouter(frontEnd *embed.FS, conn *sql.DB, storage storage.Storage) *Router {

func NewHttpRouter(frontEnd *embed.FS, conn *sql.DB) *Router {
	// return &Router{
	// files:       frontEnd,
	// 	logger:   internal.NewDefaultLogger(),
	// 	userRepo: db.NewUserRepo(conn),
	// storage:     	storage,
	// }

	return &Router{
		files:       frontEnd,
		logger:      internal.NewDefaultLogger(),
		userRepo:    db.NewUserRepo(conn),
		expenseRepo: db.NewExpenseRepo(conn),
		holidayRepo: db.NewHolidayRepo(conn),
		absenceRepo: db.NewAbsenceRepo(conn),
	}
}

func (r *Router) Start(port uint16) {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter()

	// configure cors
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})

	// register routes
	api.HandleFunc("/users", r.HandleGetUsers).Methods(http.MethodGet)
	api.HandleFunc("/users", r.HandleCreateUser).Methods(http.MethodPost)
	api.HandleFunc("/users/login", r.HandleLogin).Methods(http.MethodPost)

	api.HandleFunc("/absences", r.HandleGetAbsences).Methods(http.MethodGet)
	api.HandleFunc("/absences", r.HandleCreateAbsence).Methods(http.MethodPost)
	api.HandleFunc("/absences/{id}", r.HandleDeleteAbsence).Methods(http.MethodDelete)
	api.HandleFunc("/absences/{id}", r.HandleGetOneAbsence).Methods(http.MethodGet)

	api.HandleFunc("/holidays", r.HandleGetHolidays).Methods(http.MethodGet)
	api.HandleFunc("/holidays", r.HandleCreateHoliday).Methods(http.MethodPost)
	api.HandleFunc("/holidays/{id}", r.HandleDeleteHoliday).Methods(http.MethodDelete)
	api.HandleFunc("/holidays/{id}", r.HandleGetOneHoliday).Methods(http.MethodGet)

	api.HandleFunc("/expenses", r.HandleGetExpenses).Methods(http.MethodGet)
	api.HandleFunc("/expenses", r.HandleCreateExpense).Methods(http.MethodPost)
	api.HandleFunc("/expenses/{id}", r.HandleDeleteExpense).Methods(http.MethodDelete)
	api.HandleFunc("/expenses/{id}", r.HandleGetOneExpense).Methods(http.MethodGet)

	router.PathPrefix("/").HandlerFunc(r.HandleServeUI).Methods(http.MethodGet)

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	r.logger.Printf("Starting HTTP server on '%s'", addr)

	srv := &http.Server{
		Handler:      handlers.CORS(credentials, methods, origins, headers)(router),
		Addr:         addr,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	r.logger.Fatalln(srv.ListenAndServe())
}
