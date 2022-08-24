////////////////////////////////////////////////////////////////////////////////
//	stack_test.go  -  Ago-24-2022  -  aldebap
//
//	Test cases for the Stack Data Structure
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"
	"testing"
)

//	a few test cases - words will be pushed to the stack and an asterisk indicates a pop from the stack
var testScenarios = []struct {
	scenario    string
	input       string
	emptyResult bool
	output      string
}{
	{scenario: "No operations", input: "", emptyResult: true, output: ""},
	{scenario: "Single operation", input: "world", emptyResult: false, output: "world"},
	{scenario: "Multiple operations", input: "world rubbish * hello", emptyResult: false, output: "hello world"},
	{scenario: "Sedgewick's operation", input: "A * S A * M * p * l * E S * T * * * A * C K * *", emptyResult: true, output: ""},
}

func TestStackAsArray(t *testing.T) {

	t.Run(">>> test a series of stack operations", func(t *testing.T) {

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			//	execute operations in the input
			want := test.output
			testStack := NewStack()

			for _, operation := range strings.Split(test.input, " ") {
				if operation == "*" {
					testStack.Pop()
				} else if len(operation) > 0 {
					testStack.Push(operation)
				}
			}

			//	check the empty() result
			if test.emptyResult != testStack.IsEmpty() {
				t.Errorf("failed empty result of the stack: expected: %t result: %t", test.emptyResult, testStack.IsEmpty())
			}

			//	pop all elements from the stack
			got := ""

			for {
				item := testStack.Pop()
				if item == nil {
					break
				}
				got += item.(string) + " "
			}
			got = strings.TrimRight(got, " ")
			fmt.Printf("[debug] stack result: %s\n", got)

			//	check the result
			if want != got {
				t.Errorf("failed pushing and poping items to the stack: expected: %s result: %s", want, got)
			}
		}
	})
}
