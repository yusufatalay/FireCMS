package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	// register handler functions here
	// later secure some of this routes for security purposes

	router.HandlerFunc(http.MethodGet, "/course/:courseName", app.getCourse)
	router.HandlerFunc(http.MethodGet, "/courses", app.getAllCourses)
	return app.enableCORS(router)
}
