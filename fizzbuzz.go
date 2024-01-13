package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Enter a number: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Please enter a valid number.")
	}

	for i := 1; i <= number; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

func FizzBuzz(number int) string {
	switch {
	case number%15 == 0:
		return "Fizzbuzz!"
	case number%3 == 0:
		return "Fizz"
	case number%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", number)
	}
}
