////////////////////////////////////////////////////////////////////////////////
//	linkedList_test.go  -  Ago-20-2022  -  aldebap
//
//	Test cases for the Linked List Data Structure
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

//	a few test cases
var testScenarios = []struct {
	scenario string
	input    []string
}{
	{scenario: "Empty list", input: []string{}},
	{scenario: "Single element list", input: []string{"hello"}},
	{scenario: "Double elements list", input: []string{"hello", "world"}},
	{scenario: "Multiple elements list", input: []string{"linked", "list", "data", "structure"}},
}

//	TestGenericLinkedList test cases
func test_genericLinkedList(t *testing.T, testCase string, newLinkedList func() GenericLinkedList) {

	t.Run(">>> "+testCase, func(t *testing.T) {

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			//	add the input to the linked list
			want := ""
			testLinkedList := newLinkedList()

			for _, item := range test.input {
				want += item + " "
				testLinkedList.AddElement(item)
			}

			//	iterate throughout the list
			got := ""

			testLinkedList.First()
			for {
				item := testLinkedList.Next()
				if item == nil {
					break
				}
				got += item.(string) + " "
			}
			fmt.Printf("[debug] iteration result: %s\n", got)

			//	check the result
			if want != got {
				t.Errorf("failed adding and iterating the list: expected: %s result: %s", want, got)
			}
		}
	})
}

//	TestLinkedList linked list test cases
func TestLinkedList(t *testing.T) {
	test_genericLinkedList(t, "Insert and iterate through linked list", NewLinkedList)
}
