package survey

import (
	"embedded-survey-go/controllers"
	"embedded-survey-go/templates"
)

// Render receives a map.
// map should include a datasource value
func Render(config map[string]interface{}) string {
	// Set Global Variable Settings accessible from any part
	controllers.GlobalURIFile = config["datasource"].(string)
	questions := controllers.ReadDatasource()
	working := controllers.GlobalSettings.Config.Spreadsheet.Name
	println("Erda", working)
	return templates.Render(questions)
}

// Function that Process answers
// The answers come in a map.
// We can decide here how to process
// Now we are going to store at Google Sheet
func Process(answers map[string][]string) {
	controllers.WriteInSheet(answers)
}
