package api


import (
	"net/http"
	
	// is a lightweight http route
	"github.com/go-chi/chi/v5"
)

func NewRoute() (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Word!"))
	})
	
	return r, nil
}