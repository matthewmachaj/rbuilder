package main

func main() {
	history := &History{}

	contact := CaptureContact()
	jobs := CaptureJobs()

	history.Contact = contact
	history.Jobs = jobs

	dao := YamlHistoryDao{}
	dao.SaveHistory(*history)
}
