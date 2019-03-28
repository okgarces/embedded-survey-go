package helpers

import "strings"

func GetName(str string) string {

	name := strings.ReplaceAll(str, " ", "_")
	name = strings.ToLower(name)
	return name
}

func ConcateNateList(answers []string)string{
	return strings.Join(answers, ";")
}