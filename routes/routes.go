package routes

import (
	"fluent-go/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HelloHandler)
}
