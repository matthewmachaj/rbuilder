package main

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var OUT_FILE = filepath.Join("data", "history.yaml")

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
	JobTitle  string `yaml:"jobTitle"`
	StartDate string `yaml:"endDate"`
	EndDate   string `yaml:"startDate"`
}

type HistoryDao interface {
	GetHistory() (History, error)
	SaveHistory(h History)
}

type YamlHistoryDao struct{}

func (y *YamlHistoryDao) GetHistory() (History, error) {
	return History{}, nil
}

func (y *YamlHistoryDao) SaveHistory(history History) {
	encodeYaml(history, OUT_FILE)
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
