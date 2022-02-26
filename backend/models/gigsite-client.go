package models

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
)

func (c *ClientModel) GetGigSite(siteName string) (*GigSite, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	result, err := c.CL.Collection("Gig Sites").Doc(siteName).Get(ctx)

	if err != nil {
		return nil, err
	}

	var gigSite GigSite

	gigSite.BizModel = fmt.Sprint(result.Data()["BizModel"])
	gigSite.Comments = []string{""} // not implemented for now
	gigSite.Definition = fmt.Sprint(result.Data()["Definition / About"])
	gigSite.CompanyLogo = fmt.Sprint(result.Data()["Gig Company Logo(dummy pics)"])
	gigSite.CompanyName = fmt.Sprint(result.Data()["Gig Company Name"])
	gigSite.LandingPage = fmt.Sprint(result.Data()["Landing Page"])
	gigSite.ProfessionName = fmt.Sprint(result.Data()["Profession Name"])
	gigSite.Rating = fmt.Sprint(result.Data()["Rating-(2nd phase)"])
	gigSite.Remote = fmt.Sprint(result.Data()["Remote"])
	gigSite.Requirements = fmt.Sprint(result.Data()["Requirements"])
	gigSite.Tips = fmt.Sprint(result.Data()["Tips"])
	gigSite.Where = fmt.Sprint(result.Data()["Where"])
	gigSite.IsSaved = fmt.Sprint(result.Data()["isSaved"])

	return &gigSite, nil

}

func (c *ClientModel) GetAllGigSites() ([]*GigSite, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	gigsitedocs, err := c.CL.Collection("Gig Sites").Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	var gigsites []*GigSite

	for _, result := range gigsitedocs {

		var gigSite GigSite
		gigSite.BizModel = fmt.Sprint(result.Data()["BizModel"])
		gigSite.Comments = []string{""} // not implemented for now
		gigSite.Definition = fmt.Sprint(result.Data()["Definition / About"])
		gigSite.CompanyLogo = fmt.Sprint(result.Data()["Gig Company Logo(dummy pics)"])
		gigSite.CompanyName = fmt.Sprint(result.Data()["Gig Company Name"])
		gigSite.LandingPage = fmt.Sprint(result.Data()["Landing Page"])
		gigSite.ProfessionName = fmt.Sprint(result.Data()["Profession Name"])
		gigSite.Rating = fmt.Sprint(result.Data()["Rating-(2nd phase)"])
		gigSite.Remote = fmt.Sprint(result.Data()["Remote"])
		gigSite.Requirements = fmt.Sprint(result.Data()["Requirements"])
		gigSite.Tips = fmt.Sprint(result.Data()["Tips"])
		gigSite.Where = fmt.Sprint(result.Data()["Where"])
		gigSite.IsSaved = fmt.Sprint(result.Data()["isSaved"])

		gigsites = append(gigsites, &gigSite)
	}

	return gigsites, nil
}

func (c *ClientModel) SaveGigSite(gigSite *GigSite) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	_, err := c.CL.Collection("Gig Sites").Doc(fmt.Sprintf("%s %s", gigSite.CompanyName, gigSite.ProfessionName)).Set(ctx, map[string]interface{}{
		"BizModel":                     gigSite.BizModel,
		"Comments-(2nd phase)":         gigSite.Comments,
		"Definition / About":           gigSite.Definition,
		"Gig Company Logo(dummy pics)": gigSite.CompanyLogo,
		"Gig Company Name":             gigSite.CompanyName,
		"Landing Page":                 gigSite.LandingPage,
		"Profession Name":              gigSite.ProfessionName,
		"Rating-(2nd phase)":           gigSite.Rating,
		"Remote":                       gigSite.Remote,
		"Requirements":                 gigSite.Requirements,
		"Tips":                         gigSite.Tips,
		"Where":                        gigSite.Where,
		"isSaved":                      gigSite.IsSaved,
	})

	if err != nil {
		return err
	}

	return nil

}

func (c *ClientModel) DeleteGigSite(siteName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.CL.Collection("Gig Sites").Doc(siteName).Delete(ctx, firestore.Exists)

	if err != nil {
		return err
	}
	return nil
}
