package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		number int
		want   string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{9, "Fizz"},
		{10, "Buzz"},
		{15, "Fizzbuzz!"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d gives %q", tt.number, tt.want), func(t *testing.T) {
			got := FizzBuzz(tt.number)

			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
