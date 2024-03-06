package main

import (
	"log"

	"github.com/charmbracelet/huh"
)

func CaptureContact() Contact {
	contact := Contact{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("First Name").
				Value(&contact.FirstName),

			huh.NewText().
				Title("Last Name").
				Value(&contact.LastName),
		),
	)

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	return contact
}

func CaptureJobs() []Job {
	jobs := []Job{}
	jobs = append(jobs, CaptureJob())
	return jobs
}

func CaptureJob() Job {
	job := Job{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Employer").
				Value(&job.Employer),

			huh.NewText().
				Title("Job Title").
				Value(&job.JobTitle),
		),
	)

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	return job
}
