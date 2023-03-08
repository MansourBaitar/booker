package http

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

const frontEndLocation = "web/dist"

func (r *Router) HandleServeUI(res http.ResponseWriter, req *http.Request) {
	r.logger.Printf("Received a request for '%s' by '%s'", req.URL.Path, req.RemoteAddr)

	location := filepath.Join(frontEndLocation, req.URL.Path)
	file, err := r.files.Open(location)
	if err != nil {
		// file does not exist fallback to 'index.html'
		r.serveIndex(res)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil || fileInfo.IsDir() {
		r.serveIndex(res)
		return
	}

	// write found file
	bytes, err := io.ReadAll(file)
	if err != nil {
		r.logger.Printf("Unable to read file '%s' from embedded structure", location)
		r.writeServerError(res)
		return
	}

	mime := determineMime(location)
	res.Header().Set("Content-Type", mime)
	res.WriteHeader(http.StatusOK)
	res.Write(bytes)
}

func (r *Router) serveIndex(res http.ResponseWriter) {
	location := filepath.Join(frontEndLocation, "index.html")
	bytes, err := r.files.ReadFile(location)
	if err != nil {
		r.logger.Println("Unable to read 'index.html' from embedded structure")
		r.writeServerError(res)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(bytes)
}

func determineMime(filename string) string {
	split := strings.Split(filename, ".")
	ext := split[len(split)-1]

	switch strings.ToLower(ext) {
	case "js":
		return "text/javascript"
	case "css":
		return "text/css"
	case "svg":
		return "image/svg+xml"
	default:
		return "text/plain"
	}
}
