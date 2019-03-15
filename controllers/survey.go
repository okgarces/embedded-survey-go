package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	textFieldType = "textfield"
	checkbox = "combobox"
	radio = "radio"
)

/* Structs to get Questions and Options */

type Questions struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	InputType string `json:"type"`
	QuestionContent string `json:"content"`
	Options []string `json:"options"`
}

/**
Read Datasource with questions
 */
func ReadDatasource() Questions{
	jsonFile, err := os.Open("datasource/questions.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// Bytes and Unmarshal Json
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var questions Questions
	err = json.Unmarshal([]byte(byteValue), &questions)

	if err != nil {
		fmt.Println(err)
	}


	return questions
}