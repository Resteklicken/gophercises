// https://courses.calhoun.io/lessons/les_goph_01

package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestNewQuestionFromRecord(t *testing.T) {
	tests := []struct {
		name string
		args [2]string
		want Question
	}{
		{"1", [2]string{"5+5", "10"}, Question{"5+5", 10}},
		{"1", [2]string{"What is 24-6?", "  18  "}, Question{"What is 24-6?", 18}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuestionFromRecord(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuestionFromRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuizItem_Validate(t *testing.T) {
	type fields struct {
		question Question
		answer   int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"1", fields{Question{"", 5}, 5}, true},
		{"2", fields{Question{"", 6}, 1}, false},
		{"1", fields{Question{"", -1}, 10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := QuizItem{
				question: tt.fields.question,
				answer:   tt.fields.answer,
			}
			if got := q.Validate(); got != tt.want {
				t.Errorf("QuizItem.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnsweredQuiz_Evaluate(t *testing.T) {
	tests := []struct {
		name      string
		a         AnsweredQuiz
		wantScore int
	}{
		{"All correct gives full score",
			AnsweredQuiz{
				{Question{"", 5}, 5},
				{Question{"", 1}, 1},
				{Question{"", -100}, -100},
			}, 3,
		},
		{"Empty quiz gives 0",
			AnsweredQuiz{}, 0,
		},
		{"Mistakes deduct score",
			AnsweredQuiz{
				{Question{"", 1}, 5},
				{Question{"", -2}, 2},
				{Question{"", 0}, 0},
			}, 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := tt.a.Evaluate(); gotScore != tt.wantScore {
				t.Errorf("AnsweredQuiz.Evaluate() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
