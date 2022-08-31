////////////////////////////////////////////////////////////////////////////////
//	factorial_test.go  -  Ago-31-2022  -  aldebap
//
//	Test cases for the recursive factorial function
////////////////////////////////////////////////////////////////////////////////

package math

import (
	"fmt"
	"testing"
)

func TestRucursiveFactorial(t *testing.T) {

	t.Run(">>> Recursive factorial", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			scenario string
			input    uint8
			output   uint64
		}{
			{scenario: "0!", input: 0, output: 1},
			{scenario: "1!", input: 1, output: 1},
			{scenario: "2!", input: 2, output: 2},
			{scenario: "3!", input: 3, output: 6},
			{scenario: "4!", input: 4, output: 24},
			{scenario: "5!", input: 5, output: 120},
			{scenario: "6!", input: 6, output: 720},
		}

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			want := test.output
			got := Factorial(test.input)

			//	check the result
			if want != got {
				t.Errorf("failed calculating the factorial: expected: %d result: %d", want, got)
			}
		}
	})
}
