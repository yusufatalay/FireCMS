package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"

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

type JobPayload struct {
	Earning string   `json:"Earning"`
	GoodAt  []string `json:"Good At"`
	Like    []string `json:"Like"`
	Picture string   `json:"Picture"`
	Title   string   `json:"title"`
}

func (app *application) saveJob(w http.ResponseWriter, r *http.Request) {
	var payload JobPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot save the job"))

		return
	}

	var job models.Job

	job.Earning = payload.Earning
	job.GoodAt = payload.GoodAt
	job.Like = payload.Like
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