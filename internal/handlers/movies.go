package handlers

import (
	"fmt"
	"net/http"
)

type MoviesHandler struct{}

func (m *MoviesHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /movies", m.GetAllMovies)
	router.HandleFunc("GET /movies/{id}", m.GetmovieById)
	router.HandleFunc("DELETE /movies", m.DeleteAllMovies)
	router.HandleFunc("POST /movies", m.CreateMovie)
}

func (m *MoviesHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Todas las peliculas"))
}

func (m *MoviesHandler) GetmovieById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue(id)

	fmt.Fprintf(w, "Movie con id %s", id)
}

func (m *MoviesHandler) DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Todas las peliculas eliminadas"))
}

func (m *MoviesHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var payload models.MovieModel
}
