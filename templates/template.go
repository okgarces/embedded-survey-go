package templates

import (
	"embedded-survey-go/controllers"
	"embedded-survey-go/helpers"
	"fmt"
)


const (
	textField = "text"
	checkbox = "checkbox"
	radio = "radio"
)

func Render(questions controllers.Questionnaire) string  {
	var returnString string
	returnString = "<form method='POST'>"
	for i:=0; i < len(questions.Questions); i++{
		returnTemp := renderByInput(questions.Questions[i])(questions.Questions[i])
		returnTemp += "</br>"
		returnString += returnTemp
	}
	returnString += fmt.Sprintf("<button type='submit'> %s </button>", controllers.GlobalSettings.Config.Submit.Name )
	returnString += "</form>"
	return  returnString
}


func renderByInput(question controllers.Question) func(question2 controllers.Question) string {

	inputTypes := map[string] func(question2 controllers.Question) string{
		textField: renderText,
		radio: renderRadio,
		checkbox: renderCheckBox,
	}

	return inputTypes[question.InputType]
}

func renderText(question controllers.Question) string {
	str := "%s <br/> <input type='text' name='%s'>"
	name := helpers.GetName(question.Content)
	strReturn := fmt.Sprintf(str, question.Content, name)
	return strReturn
}

func renderRadio(question controllers.Question) string {
	mainStr := "<div> %s </div>"
	str := "<div> <input type='radio' id='%s' value='%s' name='%s' > <label> %s </label> </div>"
	name := helpers.GetName(question.Content)

	strReturn := fmt.Sprintf(mainStr, question.Content)
	for i:=0; i<len(question.Options); i++ {
		option := question.Options[i]
		optionName := helpers.GetName(option)
		strReturn += fmt.Sprintf(str, optionName, optionName, name, option )
	}

	return strReturn
}
 func renderCheckBox(question controllers.Question) string{
 	mainStr := "<div> %s </div>"
	str := "<div><input type='checkbox' id='%s' name='%s' value='%s' > <label for='%s'> %s </label> </div>"
	strReturn := fmt.Sprintf(mainStr, question.Content)
	name := helpers.GetName(question.Content)
	for i:=0; i<len(question.Options); i++ {
		option := question.Options[i]
		optionName := helpers.GetName(option)
		strReturn += fmt.Sprintf(str, optionName, name, optionName, optionName, option )
	}

	 return strReturn
 }
