package models

import "cloud.google.com/go/firestore"

// Models is the wrapper for the firesore client
type Models struct {
	CL ClientModel
}

type ClientModel struct {
	CL *firestore.Client
}

// NewModels returns models with client pool
func NewModels(cl *firestore.Client) Models {
	return Models{
		CL: ClientModel{
			CL: cl,
		},
	}
}

type Course struct {
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
	Rating           int    `json:"rating"`
}

type GigSite struct {
	BizModel       string   `json:"BizModel"`
	Comments       []string `json:"Comments-(2nd phase)"` // include this in further phase
	Definition     string   `json:"Definition / About"`
	CompanyLogo    string   `json:"Gig Company Logo(dummy pics)"`
	CompanyName    string   `json:"Gig Company Name"`
	LandingPage    string   `json:"Landing Page"`
	ProfessionName string   `json:"Profession Name"`
	Rating         string   `json:"Rating-(2nd phase)"`
	Remote         string   `json:"Remote"`
	Requirements   string   `json:"Requirements"`
	Tips           string   `json:"Tips"`
	Where          string   `json:"Where"`
	IsSaved        string   `json:"isSaved"`
}

type Job struct {
	Earning string   `json:"Earning"`
	GoodAt  []string `json:"Good At"`
	Like    []string `json:"Like"`
	Picture string   `json:"Picture"`
	Title   string   `json:"title"`
}

func getArrayFromInterface(inter interface{}) []string {

	var result []string

	if inter == nil {
		return nil
	}
	for _, v := range inter.([]interface{}) {
		result = append(result, v.(string))
	}

	return result
}
