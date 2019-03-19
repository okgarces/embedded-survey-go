package main

import (
	"bufio"
	"embedded-survey-go/controllers"
	"embedded-survey-go/templates"
	"fmt"
	"os"
)

func main() {
	questions := controllers.ReadDatasource()
	fmt.Println(templates.Render(questions))
	f, _ := os.Create("index.html")
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(templates.Render(questions))
	w.Flush()
}
