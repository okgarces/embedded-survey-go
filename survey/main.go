package survey

import (
	"embedded-survey-go/controllers"
	"embedded-survey-go/templates"
)

// Render receives a map.
// map should include a datasource value
func Render(config map[string]interface{}) string {
	datasourceString := config["datasource"].(string)
	questions := controllers.ReadDatasource(datasourceString)
	return templates.Render(questions)
}
