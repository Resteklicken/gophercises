// https://courses.calhoun.io/lessons/les_goph_01

package main

import (
	"strconv"
	"strings"
)

func main() {

}

type Question struct {
	question string
	answer   int
}

func NewQuestionFromRecord(record [2]string) Question {
	trimmed := strings.TrimSpace(record[1])
	converted, err := strconv.Atoi(trimmed)
	if err != nil {
		panic(err)
	}
	return Question{record[0], converted}
}
