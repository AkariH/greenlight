package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /v1/healthcheck healthcheckHandler Show application information
// POST /v1/movies createMovieHandler Create a new movie
// GET /v1/movies/:id showMovieHandler Show the details of a specific movie

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	app.logger.Println("get request from create movie")
	fmt.Fprintln(w, "hello from create movie")

}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	app.logger.Println("get request from show movie")
	id, err := app.readIDParam(r)
	if err != nil {
		app.logger.Println(err)
		http.NotFound(w, r)
		return
	}
	// Otherwise, interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
