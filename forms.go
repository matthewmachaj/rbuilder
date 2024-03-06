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
	newJob := true

	for newJob {
		job, shouldAddNewJob := CaptureJob()
		jobs = append(jobs, job)
		newJob = shouldAddNewJob
	}
	return jobs
}

func CaptureJob() (Job, bool) {
	job := Job{}
	shouldAddNewJob := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Employer").
				Value(&job.Employer),

			huh.NewText().
				Title("Job Title").
				Value(&job.JobTitle),

			huh.NewConfirm().
				Title("Would you like to add another job?").
				Affirmative("Yes").
				Negative("No").
				Value(&shouldAddNewJob),
		),
	)

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	return job, shouldAddNewJob
}
