package main

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"messge"`
}

func (app *application) getCourse(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	courseName := params.ByName("courseName") // :4000/course/<courseName>
	app.logger.Println("course wanted ", courseName)
	course, err := app.models.CL.Get(courseName)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, course, "course")

	if err != nil {
		app.logger.Println(errors.New("cannot write to the browser"))
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := app.models.CL.GetAllCourses()

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, courses, "courses")

	if err != nil {
		app.logger.Println(errors.New("cannot write to the browser"))
		app.errorJSON(w, err)
		return
	}

}
