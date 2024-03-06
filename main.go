package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func HuhForm(history *History) {
	job := Job{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("First Name").
				Value(&history.Contact.FirstName),

			huh.NewText().
				Title("Last Name").
				Value(&history.Contact.LastName),
		),

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

	history.Jobs = append(history.Jobs, job)
}

func main() {
	history := &History{}
	HuhForm(history)

	fmt.Println(history)

	dao := YamlHistoryDao{}
	dao.SaveHistory(*history)
}
