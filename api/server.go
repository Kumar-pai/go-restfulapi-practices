package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

type Server struct {
	*http.Server
}

func NewServer() (*Server, error) {
	log.Println("configuring Server")
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf(".env file not found")
	}

	appUrl := os.Getenv("APP_URL")
	if appUrl == ""{
		return nil, fmt.Errorf("app url is null")
	}
	log.Printf("APP URL: %s\n", appUrl)

	port := os.Getenv("APP_PORT")
	if port == ""{
		return nil, fmt.Errorf("app port is null")
	}
	log.Printf("APP PORT: %s\n", port)

	addr := appUrl + ":" + port

	svr := http.Server{
		Addr: addr,
	}

	return &Server{&svr}, nil
}

func (svr *Server) Start() {
	log.Println("starting server")
	// create a http server  goroutine
	go func() {
		if err := svr.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", svr.Addr)

	//create a listen quit channel
	quit := make(chan os.Signal)
	// register quit
	signal.Notify(quit)
	sig := <- quit
	log.Printf("Shutting down server ... Reason: %s", sig)

	// When use context backhround shutdown http server
	if err := svr.Shutdown(context.Background()); err != nil {
		// if shutdown return error trigger panic to close
		panic(err)
	}
	log.Println("Server stpped")
}
