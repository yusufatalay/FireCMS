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
	router.HandlerFunc(http.MethodPost, "/savecourse", app.saveCourse)
	router.HandlerFunc(http.MethodGet, "/deletecourse/:courseName", app.deleteCourse)

	router.HandlerFunc(http.MethodGet, "/job/:jobName", app.getJob)
	router.HandlerFunc(http.MethodGet, "/jobs", app.getAllJobs)
	router.HandlerFunc(http.MethodPost, "/savejob", app.saveJob)
	router.HandlerFunc(http.MethodGet, "/deletejob/:jobName", app.deleteJob)

	router.HandlerFunc(http.MethodGet, "/gigsite/:gigSiteName", app.getGigSite)
	router.HandlerFunc(http.MethodGet, "/gigsites", app.getAllGigSites)
	router.HandlerFunc(http.MethodPost, "/savegigsite", app.saveGigSite)
	router.HandlerFunc(http.MethodGet, "/deletegigsite/:gigSiteName", app.deleteGigSite)

	return app.enableCORS(router)
}
