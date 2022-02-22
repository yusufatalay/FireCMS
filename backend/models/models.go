package models





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
	Comments       []string `json:"-"` // include this in further phase
	Definition     string   `json:"Definition / About"`
	Company        string   `json:"Gig Company"`
	Logo           string   `json:"Logo(dummy pics)"`
	CompanyName    string   `json:"Gig Company Name"`
	LandingPage    string   `json:"Landing Page"`
	ProfessionName string   `json:"Profession Name"`
	Rating         string   `json:"-"`
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
