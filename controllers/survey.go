package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* Structs to get Questionnaire and Options */
type Questionnaire struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	InputType string   `json:"type"`
	Content   string   `json:"content"`
	Options   []string `json:"options"`
}

/**
Read Datasource with questions
 */
func ReadDatasource() Questionnaire {
	jsonFile, err := os.Open("datasource/questions.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// Bytes and Unmarshal Json
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var questions Questionnaire
	err = json.Unmarshal([]byte(byteValue), &questions)

	if err != nil {
		fmt.Println(err)
	}


	return questions
}