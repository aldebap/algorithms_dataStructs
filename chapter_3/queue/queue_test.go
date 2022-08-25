////////////////////////////////////////////////////////////////////////////////
//	queue_test.go  -  Ago-24-2022  -  aldebap
//
//	Test cases for the Queue Data Structure
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"
	"testing"
)

//	a few test cases - words will be put to the queue and an asterisk indicates a get from the queue
var testScenarios = []struct {
	scenario    string
	input       string
	emptyResult bool
	output      string
}{
	{scenario: "No operations", input: "", emptyResult: true, output: ""},
	{scenario: "Single operation", input: "hello", emptyResult: false, output: "hello"},
	{scenario: "Multiple operations", input: "rubbish * hello world", emptyResult: false, output: "hello world"},
	{scenario: "Sedgewick's operation", input: "A * S A * M * p * l * E S * T * * * A * C K * *", emptyResult: true, output: ""},
}

func TestStackAsArray(t *testing.T) {

	t.Run(">>> test a series of queue operations", func(t *testing.T) {

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			//	execute operations in the input
			want := test.output
			testQueue := NewQueue()

			for _, operation := range strings.Split(test.input, " ") {
				if operation == "*" {
					testQueue.Get()
				} else if len(operation) > 0 {
					testQueue.Put(operation)
				}
			}

			//	check the empty() result
			if test.emptyResult != testQueue.IsEmpty() {
				t.Errorf("failed empty result of the queue: expected: %t result: %t", test.emptyResult, testQueue.IsEmpty())
			}

			//	pop all elements from the stack
			got := ""

			for {
				item := testQueue.Get()
				if item == nil {
					break
				}
				got += item.(string) + " "
			}
			got = strings.TrimRight(got, " ")
			fmt.Printf("[debug] queue result: %s\n", got)

			//	check the result
			if want != got {
				t.Errorf("failed puting and getting items from the queue: expected: %s result: %s", want, got)
			}
		}
	})
}
