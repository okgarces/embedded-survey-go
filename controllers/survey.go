package controllers

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Settings struct {
	Config SettingsConfig `json:"settings"`
}

type SettingsConfig struct {
	Spreadsheet SheetSettings  `json:"spreadsheet"`
	Submit      SubmitSettings `json:"submit"`
}

type SheetSettings struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type SubmitSettings struct {
	Name string `json:"name"`
}

/* Structs to get Questionnaire and Options */
type Questionnaire struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	InputType string   `json:"type"`
	Content   string   `json:"content"`
	Options   []string `json:"options"`
}

// Define Global Settings
var GlobalSettings Settings
/**
Read Datasource with questions
 */
func ReadDatasource(url string) Questionnaire {
	jsonFile, _ := os.Open(url)
	// Bytes and Unmarshal Json
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var questions Questionnaire
	// Do not get err
	_ = json.Unmarshal([]byte(byteValue), &questions)
	_ = json.Unmarshal([]byte(byteValue), &GlobalSettings)

	return questions
}
