// https://courses.calhoun.io/lessons/les_goph_01

package main

import "strconv"

func main() {

}

type Question struct {
	question string
	answer   int
}

func NewQuestionFromRecord(record [2]string) Question {
	conv, err := strconv.Atoi(record[1])
	if err != nil {
		panic(err)
	}
	return Question{record[0], conv}
}
