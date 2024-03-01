package api


import (
	"net/http"
	
	// is a lightweight http route
	"github.com/go-chi/chi/v5"
	"github.com/Kumar-pai/go-restfulapi-practices/database"
)

func NewRoute() (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Word!"))
	})
	
	_, err := database.DBConn()
	if err != nil {
		return nil, err
	}

	return r, nil
}