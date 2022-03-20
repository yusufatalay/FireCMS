package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getJob(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	jobName := params.ByName("jobName") // /job/:jobName

	job, err := app.models.CL.GetJob(jobName)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot get job"))

		return
	}

	err = app.writeJSON(w, http.StatusOK, job, "job")

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot get job"))

		return
	}
}

func (app *application) getAllJobs(w http.ResponseWriter, r *http.Request) {

	jobs, err := app.models.CL.GetAllJobs()
	if err != nil {

		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot get jobs"))

		return
	}

	err = app.writeJSON(w, http.StatusOK, jobs, "jobs")

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot get jobs"))
		return
	}
}

func (app *application) saveJob(w http.ResponseWriter, r *http.Request) {
	var payload models.Job

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot save the job"))

		return
	}

	var job models.Job

	job.Earning = payload.Earning
	job.GoodAt = payload.GoodAt
	job.Like = strings.Split(payload.Like[0], ",")
	job.Picture = payload.Picture
	job.Title = payload.Title

	err = app.models.CL.SaveJob(&job)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot save the job"))

		return
	}
	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot save the job"))
		return
	}

}

func (app *application) deleteJob(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	jobName := params.ByName("jobName") // /deleteJob/:jobName

	err := app.models.CL.DeleteJob(jobName)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot delete Job"))
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
