////////////////////////////////////////////////////////////////////////////////
//	selectionSort_test.go  -  Sep-1-2022  -  aldebap
//
//	Test cases for the "selection sort" algorithm
////////////////////////////////////////////////////////////////////////////////

package selectionSort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	t.Run(">>> Selection Sort", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			scenario string
			input    []string
			output   []string
		}{
			{
				scenario: "unordered list #1",
				input:    []string{"a", "s", "o", "r", "t", "i", "n", "g", "a", "r", "r", "a", "y"},
				output:   []string{"a", "a", "a", "g", "i", "n", "o", "r", "r", "r", "s", "t", "y"},
			},
			{
				scenario: "unordered list #2",
				input:    []string{"m", "l", "k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a"},
				output:   []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
			},
		}

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			want := test.output
			Sort(test.input)
			got := test.input

			//	check the result
			for i, sortedElement := range got {
				if sortedElement != want[i] {
					t.Errorf("fail sorting the : element %d expected: %s result: %s", i, want[i], sortedElement)
				}
			}
		}
	})
}
