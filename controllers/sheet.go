package controllers

import (
	"embedded-survey-go/helpers"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"time"
)

const (
	FirstReadRow = "A1:1"
	FirstWriteRoad = "A2"
)

func getServiceGoogle() *sheets.Service{
	ctx := context.Background()
	// Get Credentials from File
	clientOptions := option.WithCredentialsFile(GlobalURIFile)
	srv, _ := sheets.NewService(ctx, clientOptions)
	return srv

}

func writeQuestionsInFirstRow(){
	questions := ReadDatasource()
	questionsSlugs := make([]interface{}, 0)
	for i:=0; i < len(questions.Questions); i++{
		questionsSlugs = append(questionsSlugs, helpers.GetName(questions.Questions[i].Content))
	}

	srv := getServiceGoogle()
	spreadsheetId := GlobalSettings.Config.Spreadsheet.Id
	writeRange := FirstReadRow
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, questionsSlugs)
	_, _ = srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
}

// helpers tick used to avoid stack overflow
var internalTick = 0
var maxTicks = 5
func tick() int {
	internalTick += 1
	return internalTick
}
func rebootTick() {
	internalTick = 0
}

//////////
func ReadFirst() []string {

	srv := getServiceGoogle()
	spreadsheetId := GlobalSettings.Config.Spreadsheet.Id
	readRange := FirstReadRow
	values, _ := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	var orderQuestions []string
	for _, row := range values.Values {
		for _, column := range  row {
			orderQuestions = append(orderQuestions, column.(string))
		}
	}

	if len(orderQuestions) == 0 {
		if tick() > maxTicks{
			rebootTick()
			return nil
		}
		writeQuestionsInFirstRow()
		return ReadFirst()
	}

	return orderQuestions
}

func WriteInSheet(answers map[string][]string){
	orderQuestions := ReadFirst()
	newAnswers := make([]interface{}, 0)
	for i:=0; i < len(orderQuestions); i+=1{
		answer := answers[orderQuestions[i]]
		if len(answer) > 0 {
			newAnswers = append(newAnswers, helpers.ConcateNateList(answer) )
		}else{
			newAnswers = append(newAnswers, "")
		}
	}

	// Add Created date
	newAnswers = append(newAnswers, time.Now().String())

	srv := getServiceGoogle()
	spreadsheetId := GlobalSettings.Config.Spreadsheet.Id
	writeRange := FirstWriteRoad
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, newAnswers)
	_, _ = srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
}
