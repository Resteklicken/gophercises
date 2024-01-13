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
