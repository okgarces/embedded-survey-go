package main

import (
	"bufio"
	"embedded-survey-go/survey"
	"fmt"
	"os"
)

func main() {
	questionsStr := survey.Render(map[string]interface{}{"datasource": "datasource/questions.json"})
	fmt.Println(questionsStr)
	f, _ := os.Create("index.html")
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(questionsStr)
	w.Flush()
}


