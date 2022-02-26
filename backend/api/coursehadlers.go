package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)



// getCourse gets single course with the given name in the URL.
func (app *application) getCourse(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	courseName := params.ByName("courseName") // :4000/course/<courseName>
	course, err := app.models.CL.GetCourse(courseName)

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

// getAllCourses returns all the courses located in the "Course" collection.
func (app *application) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := app.models.CL.GetAllCourses()

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, courses, "courses")

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot write to the browser"))
		return
	}

}

type CoursePayload struct {
	CourseName       string `json:"Course Name"`
	CourseOwner      string `json:"Course Owner"`
	CoursePicture    string `json:"Course Picture"`
	Certificate      string `json:"Certificate"`
	OwnerPicture     string `json:"Course (owner) Picture (dummy pics)"`
	IntroductionLink string `json:"Introduction Link"`
	Location         string `json:"Location"`
	Price            string `json:"Price"`
	ProfessionName   string `json:"Profession Name"`
	IsSaved          string `json:"isSaved"`
	Rating           string `json:"rating"`
}

// saveCourse saves given course in the r's body to the Firestore.
// If the course already exists in the database then it overrides.
func (app *application) saveCourse(w http.ResponseWriter, r *http.Request) {
	var payload CoursePayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}

	var course models.Course

	course.CourseName = payload.CourseName
	course.CourseOwner = payload.CourseOwner
	course.CoursePicture = payload.CoursePicture
	course.Certificate = payload.Certificate
	course.OwnerPicture = payload.OwnerPicture
	course.IntroductionLink = payload.IntroductionLink
	course.Location = payload.Location
	course.Price = payload.Price
	course.ProfessionName = payload.ProfessionName
	course.IsSaved = payload.IsSaved
	course.Rating, _ = strconv.Atoi(payload.Rating)

	err = app.models.CL.SaveCourse(&course)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("could not save the course"))
		return
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")

	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

// deleteCourse deletes the course with the course name gathered from the URL.
func (app *application) deleteCourse(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	courseName := params.ByName("courseName") // portNo/deleteCourse/<courseName>

	err := app.models.CL.DeleteCourse(courseName)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("could not delete the course"))
		return
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")

	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
