package models

import (
	"context"
	"fmt"
	"time"
)

// this file contains courses spesific fucntions that either returns single course , array of courses or edits courses fromm given firestore.client

// get single course
func (c *ClientModel) Get(courseName string) (*Course, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	var course Course

	result, err := c.CL.Collection("Courses").Doc(courseName).Get(ctx)
	// create an course object from result
	course.CourseName = fmt.Sprint(result.Data()["Course Name"])
	course.CourseOwner = fmt.Sprint(result.Data()["Course Owner"])
	course.CoursePicture = fmt.Sprint(result.Data()["Course Picture"])
	course.Certificate = fmt.Sprint(result.Data()["Certificate"])
	course.OwnerPicture = fmt.Sprint(result.Data()["Course (owner) Picture (dummy pics)"])
	course.IntroductionLink = fmt.Sprint(result.Data()["Introduction Link"])
	course.Location = fmt.Sprint(result.Data()["Location"])
	course.Price = fmt.Sprint(result.Data()["Price"])
	course.ProfessionName = fmt.Sprint(result.Data()["Profession Name"])
	course.IsSaved = fmt.Sprint(result.Data()["isSaved"])
	course.Rating = int(result.Data()["rating"].(int64))

	if err != nil {
		return nil, err
	}
	return &course, nil

}

// GetAllCourses returns all the courses on the "Courses" collection
func (c *ClientModel) GetAllCourses() ([]*Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	var courses []*Course

	courseDocs, err := c.CL.Collection("Courses").Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	for _, doc := range courseDocs {
		var course Course
		course.CourseName = fmt.Sprint(doc.Data()["Course Name"])
		course.CourseOwner = fmt.Sprint(doc.Data()["Course Owner"])
		course.CoursePicture = fmt.Sprint(doc.Data()["Course Picture"])
		course.Certificate = fmt.Sprint(doc.Data()["Certificate"])
		course.OwnerPicture = fmt.Sprint(doc.Data()["Course (owner) Picture (dummy pics)"])
		course.IntroductionLink = fmt.Sprint(doc.Data()["Introduction Link"])
		course.Location = fmt.Sprint(doc.Data()["Location"])
		course.Price = fmt.Sprint(doc.Data()["Price"])
		course.ProfessionName = fmt.Sprint(doc.Data()["Profession Name"])
		course.IsSaved = fmt.Sprint(doc.Data()["isSaved"])
		course.Rating = int(doc.Data()["rating"].(int64))

		courses = append(courses, &course)
	}

	return courses, nil
}

// SaveCourse inserts a new course to the "Courses" collection if there is no other course with same name , otherwise it overrides the existing course
func (c *ClientModel) SaveCourse(course *Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Set() takes a contex and an interface as a parameter, to be sure about the data, I will convert the course object into a map[string]interface() instead of sending it as it is

	_, err := c.CL.Collection("Courses").Doc(course.CourseName).Set(ctx, map[string]interface{}{
		"Course Name":                         course.CourseName,
		"Course Owner":                        course.CourseOwner,
		"Course Picture":                      course.CoursePicture,
		"Certificate":                         course.Certificate,
		"Course (owner) Picture (dummy pics)": course.OwnerPicture,
		"Introduction Link":                   course.IntroductionLink,
		"Location":                            course.Location,
		"Price":                               course.Price,
		"Profession Name":                     course.ProfessionName,
		"isSaved":                             course.IsSaved,
		"rating":                              int64(course.Rating),
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteCourse deletes the given course document from the Courses collection. courseName is case sensetive
func (c *ClientModel) DeleteCourse(courseName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	_, err := c.CL.Collection("Courses").Doc(courseName).Delete(ctx)

	if err != nil {
		return err
	}

	return nil
}
