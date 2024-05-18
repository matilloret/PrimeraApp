package api

import (
	"log"
	"net/http"
)

type APIServer struct {
	port string
}

// generar un constructor de la estructura
func NewAPIServer(port string) *APIServer {
	return &APIServer{
		port: port,
	}
}

func (app *APIServer) Run() error {
	// Enrutador
	router := http.NewServeMux()

	//server
	log.Println("listening on port: ", app.port)

	server := http.Server{
		Addr:    app.port,
		Handler: router,
	}

	return server.ListenAndServe()
}