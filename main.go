package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v3"
)

const OUT_FILE = "history.yaml"

type History struct {
	Contact Contact
	Jobs    []Job
}

type Contact struct {
	FirstName string `yaml:"firstName"`
	LastName  string `yaml:"lastName"`
}

type Job struct {
	Employer  string
	StartDate string `yaml:"endDate"`
	EndDate   string `yaml:"startDate"`
	JobTitle  string `yaml:"jobTitle"`
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
}

func encodeYaml(history History, out string) {
	file, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)

	err = enc.Encode(history)
	if err != nil {
		log.Fatalf("error encoding: %v", err)
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

func main() {
	// validate := func(input string) error {
	// 	_, err := strconv.ParseFloat(input, 64)
	// 	if err != nil {
	// 		return errors.New("Invalid number")
	// 	}
	// 	return nil
	// }

	contact := gatherContact()
	jobs := gatherJobs()

	history := History{
		Contact: contact,
		Jobs:    jobs,
	}

	encodeYaml(history, OUT_FILE)
}
