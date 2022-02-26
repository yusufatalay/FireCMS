package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getGigSite(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	gigsitename := params.ByName("gigSiteName") // /gigsite/:gigSiteName

	gigsite, err := app.models.CL.GetGigSite(gigsitename)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot retrieve Gig Site"))
		return
	}

	err = app.writeJSON(w, http.StatusOK, gigsite, "gigsite")

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot load Gig Site"))
		return
	}

}

func (app *application) getAllGigSites(w http.ResponseWriter, r *http.Request) {

	gigsites, err := app.models.CL.GetAllGigSites()

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot retrieve Gig Sites"))
		return
	}

	err = app.writeJSON(w, http.StatusOK, gigsites, "gigsites")
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot load Gig Sites"))
		return
	}
}

func (app *application) saveGigSite(w http.ResponseWriter, r *http.Request) {

	var payload models.GigSite

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("given Gig Site contains error(s)"))
		return
	}

	err = app.models.CL.SaveGigSite(&payload)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot save the Gig Site"))
		return
	}

	var ok jsonResp

	ok.OK = true

	err = app.writeJSON(w, http.StatusOK, ok, "response")

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}

}

func (app *application) deleteGigSite(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	gigsitename := params.ByName("gigSiteName")

	err := app.models.CL.DeleteGigSite(gigsitename)

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, errors.New("cannot delete Gig Site"))
		return
	}

	var ok jsonResp

	ok.OK = true

	err = app.writeJSON(w, http.StatusOK, ok, "response")

	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}

}
