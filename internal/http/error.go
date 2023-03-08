package http

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

func (r *Router) writeApiError(res http.ResponseWriter, status int, apiErr apiError) {
	bytes, err := json.Marshal(apiErr)
	if err != nil {
		r.writeServerError(res)
		return
	}

	res.WriteHeader(status)
	res.Write(bytes)
}

func (r *Router) writeServerError(res http.ResponseWriter) {
	http.Error(res, "Internal Server Error", http.StatusInternalServerError)
}

func (r *Router) writeNotImplemented(res http.ResponseWriter) {
	http.Error(res, "Not Implemented", http.StatusNotImplemented)
}
