////////////////////////////////////////////////////////////////////////////////
//	expression_test.go  -  Ago-25-2022  -  aldebap
//
//	Test cases for the simple expression parser
////////////////////////////////////////////////////////////////////////////////

package expression

import (
	"fmt"
	"strings"
	"testing"
)

//	Test_infix2postfix test cases for the conversion from infix -> postfix
func Test_infix2postfix(t *testing.T) {

	//	a few test cases
	var testScenarios = []struct {
		scenario string
		input    string
		output   string
	}{
		{scenario: "addition", input: "2 + 5", output: "2 5 +"},
		{scenario: "subtraction", input: "5 - 2", output: "5 2 -"},
		{scenario: "multiplication", input: "2 * 5", output: "2 5 *"},
		{scenario: "division", input: "10 / 2", output: "10 2 /"},
		{scenario: "one parenthesis", input: "( 4 + 6 ) / 2", output: "4 6 + 2 /"},
		{scenario: "two parenthesis", input: "( 4 + ( 2 * 3 ) ) / 2", output: "4 2 3 * + 2 /"},
		{scenario: "unbalanced parenthesis", input: "( 4 + 6 / 2", output: "expression with unbalanced parenthesis"},
	}

	t.Run(">>> test conversion from infix -> postfix", func(t *testing.T) {

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			//	execute conversion from infix -> postfix
			postfix, err := infix2postfix(test.input)
			if err != nil {
				if err.Error() != test.output {
					t.Errorf("unexpected error converting from infix -> postfix: %s", err)
				}
				continue
			}
			want := test.output

			got := ""

			for {
				item := postfix.Pop()
				if item == nil {
					break
				}
				got = item.(string) + " " + got
			}
			got = strings.TrimRight(got, " ")
			fmt.Printf("[debug] postfix result: %s\n", got)

			//	check the result
			if want != got {
				t.Errorf("fail converting from infix -> postfix: expected: %s result: %s", want, got)
			}
		}
	})
}
