package http

import (
	"app/internal/model"
	"app/internal/validate"
	"encoding/json"
	"net/http"
	"time"
)

func (r *Router) HandleGetUsers(res http.ResponseWriter, req *http.Request) {
	users, err := r.userRepo.FindAll()
	if err != nil {
		// TODO: write error response
		r.writeServerError(res)
		return
	}

	b, err := json.Marshal(users)
	res.Header().Set("Content-Type", "application/json")
	res.Write(b)
}

func (r *Router) HandleLogin(res http.ResponseWriter, req *http.Request) {
	isEmail := validate.IsMail(req.FormValue("email"))
	if isEmail != true {
		// TODO: return decent HTTP response
		r.logger.Printf("email was malformed")
		r.writeServerError(res)
		return
	}
	user, err := r.userRepo.GetUserByEmail(req.FormValue("email"), req.FormValue("password"))
	b, err := json.Marshal(user)
	if err != nil {
		// TODO: return decent HTTP response
		r.logger.Printf("unable to marshal user")
		r.writeServerError(res)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(b)
}

func (r *Router) HandleCreateUser(res http.ResponseWriter, req *http.Request) {

	isEmail := validate.IsMail(req.FormValue("email"))
	if isEmail != true {
		// TODO: return decent HTTP response
		r.logger.Printf("email was malformed")
		r.writeServerError(res)
		return
	}

	birthDate, err := validate.IsDate(req.FormValue("birthDate"))
	if err != nil {
		r.logger.Printf("unable to parse birth date, %v", err)
		// TODO: throw bad request
		r.writeServerError(res)
		return
	}

	user, err := r.userRepo.Save(model.User{
		FirstName: req.FormValue("firstName"),
		LastName:  req.FormValue("lastName"),
		Email:     req.FormValue("email"),
		Password:  req.FormValue("password"),
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	})
	b, err := json.Marshal(user)
	if err != nil {
		// TODO: return decent HTTP response
		r.logger.Printf("unable to marshal user")
		r.writeServerError(res)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	// res.Header().Set("Location", fmt.Sprintf("/api/users/%d", user.Id))
	res.WriteHeader(http.StatusCreated)
	res.Write(b)
}
