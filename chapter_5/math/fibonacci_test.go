////////////////////////////////////////////////////////////////////////////////
//	fibonacci_test.go  -  Ago-31-2022  -  aldebap
//
//	Test cases for the recursive Fibonacci function
////////////////////////////////////////////////////////////////////////////////

package math

import (
	"fmt"
	"testing"
)

func TestRucursiveFibonacci(t *testing.T) {

	t.Run(">>> Recursive Fibonacci", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			scenario string
			input    uint8
			output   uint64
		}{
			{scenario: "0", input: 0, output: 1},
			{scenario: "1", input: 1, output: 1},
			{scenario: "2", input: 2, output: 2},
			{scenario: "3", input: 3, output: 3},
			{scenario: "4", input: 4, output: 5},
			{scenario: "5", input: 5, output: 8},
			{scenario: "6", input: 6, output: 13},
			{scenario: "7", input: 7, output: 21},
			{scenario: "8", input: 8, output: 34},
			{scenario: "9", input: 9, output: 55},
			{scenario: "10", input: 10, output: 89},
		}

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			want := test.output
			got := Fibonacci(test.input)

			//	check the result
			if want != got {
				t.Errorf("failed calculating the Fibonacci: expected: %d result: %d", want, got)
			}
		}
	})
}
