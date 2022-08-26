////////////////////////////////////////////////////////////////////////////////
//	expression.go  -  Ago-25-2022  -  aldebap
//
//	Implementation of a simple expression parser
////////////////////////////////////////////////////////////////////////////////

package expression

import (
	"errors"
	"log"

	"github.com/aldebap/algorithms_dataStructs/chapter_3/stack"
)

//	Parse parses a simple expression using a stack
func Parse(expression string) (float64, error) {

	postfix, err := infix2postfix(expression)
	if err != nil {
		log.Fatalf("error parsing expression: %s", err.Error())
	}

	return evaluatePolishReverse(postfix)
}

//	infix2postfix read the infix expression and create a stack with the postfix version of it
func infix2postfix(expression string) (stack.Stack, error) {
	postfix := stack.New()
	operand := ""
	operator := stack.New()
	parenthesis := 0

	for _, char := range expression {
		switch char {
		case '+', '-', '*', '/':
			operator.Push(string(char))

		case '(':
			parenthesis++

		case ')':
			if parenthesis == 0 {
				return nil, errors.New("expression with missing open parenthesis")
			}
			if !operator.IsEmpty() {
				postfix.Push(operator.Pop())
			}
			parenthesis--

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			operand += string(char)

		case ' ':
			if len(operand) > 0 {
				postfix.Push(operand)
				operand = ""
			}
		}
	}
	//	make sure the parenthesis are balanced till the end
	if parenthesis != 0 {
		return nil, errors.New("expression with unbalanced parenthesis")
	}

	//	make sure the last operand and operators are pushed to the stack
	if len(operand) > 0 {
		postfix.Push(operand)
	}
	if !operator.IsEmpty() {
		postfix.Push(operator.Pop())
	}

	return postfix, nil
}

//	evaluatePolishReverse evaluate the Polish reverse expression (postfix) and return a numerical result
func evaluatePolishReverse(expression stack.Stack) (float64, error) {
	return 0, nil
}
