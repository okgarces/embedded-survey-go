package main

import (
	"embedded-survey-go/controllers"
	"fmt"
)

func main() {
	questions := controllers.ReadDatasource()

	for i:=0; i < len(questions.Questions); i++{
		fmt.Println(questions.Questions[i].QuestionContent)
		fmt.Println(questions.Questions[i].InputType)
		fmt.Println(questions.Questions[i].Options)
	}

}
