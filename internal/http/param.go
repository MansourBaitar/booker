package http

import (
	"app/internal/validate"
	"log"
	"net/http"
	"strconv"
)

// queryParam defines a function that will look for a query parameter and
// extract the value of that parameter or return the default value if not found.
func queryParam(req *http.Request, name string, defaultValue string) string {
	value := req.URL.Query().Get(name)
	log.Printf("value is: %s", value)
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func queryParamUint(req *http.Request, name string, defaultValue uint) uint {
	param := queryParam(req, name, strconv.FormatUint(uint64(defaultValue), 10))
	log.Printf("param is: %s", param)
	i, err := validate.IsUint(param)
	if err != nil {
		return defaultValue
	}

	return i
}
