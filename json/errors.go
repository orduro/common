package json

import (
	"net/http"
	"time"
)

type ErrorResponse struct {
	StatusCode int    `json:"code,omitempty"`
	Status     string `json:"status,omitempty"`
	ErrorMsg   string `json:"message,omitempty"`
	Path       string `json:"path,omitempty"`
	Timestamp  string `json:"timestamp"`
}

func WriteError(w http.ResponseWriter, r *http.Request, status int, message string) {
	e := &ErrorResponse{
		ErrorMsg:   message,
		StatusCode: status,
		Status:     http.StatusText(status),
		Path:       r.URL.Path,
		Timestamp:  time.Now().UTC().Format("2006-01-02T15:04:05Z"),
	}
	Write(w, status, map[string]any{
		"error": e,
	})

}
