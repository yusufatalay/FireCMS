package models

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
)

func (c *ClientModel) GetJob(jobName string) (*Job, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	var job Job

	result, err := c.CL.Collection("Jobs").Doc(jobName).Get(ctx)

	if err != nil {
		return nil, err
	}

	job.Earning = fmt.Sprint(result.Data()["Earning"])
	job.GoodAt = getArrayFromInterface(result.Data()["Good At"])
	job.Like = getArrayFromInterface(result.Data()["Like"])
	job.Picture = fmt.Sprint(result.Data()["Picture"])
	job.Title = fmt.Sprint(result.Data()["title"])

	return &job, nil
}

func (c *ClientModel) GetAllJobs() ([]*Job, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	var jobs []*Job
	jobDocs, err := c.CL.Collection("Jobs").Documents(ctx).GetAll()

	if err != nil {
		return nil, err
	}

	for _, jobDoc := range jobDocs {
		var job Job
		job.Earning = fmt.Sprint(jobDoc.Data()["Earning"])
		job.GoodAt = getArrayFromInterface(jobDoc.Data()["Good At"])
		job.Like = getArrayFromInterface(jobDoc.Data()["Like"])
		job.Picture = fmt.Sprint(jobDoc.Data()["Picture"])
		job.Title = fmt.Sprint(jobDoc.Data()["title"])
		
		jobs = append(jobs, &job)
	}

	return jobs, nil

}

func (c *ClientModel) SaveJob(job *Job) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	_, err := c.CL.Collection("Jobs").Doc(job.Title).Set(ctx, map[string]interface{}{
		"Earning": job.Earning,
		"Good At": job.GoodAt,
		"Like":    job.Like,
		"Picture": job.Picture,
		"title":   job.Title,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *ClientModel) DeleteJob(jobName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	_, err := c.CL.Collection("Jobs").Doc(jobName).Delete(ctx, firestore.Exists)

	if err != nil {
		return err
	}
	return nil
}
