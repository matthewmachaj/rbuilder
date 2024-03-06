package promptui

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}
func gatherContact() Contact {
	firstName := gatherText("First Name")
	lastName := gatherText("Last Name")

	contact := Contact{
		firstName,
		lastName,
	}

	return contact
}

func gatherJobs() []Job {
	job := gatherJob()
	jobs := []Job{
		job,
	}
	return jobs
}

func gatherJob() Job {
	employer := gatherText("Employer")
	startDate := gatherText("Start Date")
	endDate := gatherText("End Date")
	jobTitle := gatherText("Job Title")

	return Job{
		Employer:  employer,
		StartDate: startDate,
		EndDate:   endDate,
		JobTitle:  jobTitle,
	}
}

func gatherText(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}
	text, err := prompt.Run()
	checkError(err)

	return text
}
