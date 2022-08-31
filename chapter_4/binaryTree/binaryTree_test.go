////////////////////////////////////////////////////////////////////////////////
//	binaryTree_test.go  -  Ago-28-2022  -  aldebap
//
//	Test cases for the Binary Tree Data Structure
////////////////////////////////////////////////////////////////////////////////

package binaryTree

import (
	"fmt"
	"testing"
)

//	TestBinaryTree test cases
func TestBinaryTree(t *testing.T) {

	t.Run(">>> Add and find nodes in a tree", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			scenario string
			input    []string
			search   string
			exists   bool
		}{
			{scenario: "Empty tree", input: []string{}, search: "once", exists: false},
			{scenario: "Single element tree", input: []string{"once"}, search: "once", exists: true},
			{scenario: "Double elements tree", input: []string{"once", "a"}, search: "has", exists: false},
			{scenario: "Multiple elements tree", input: []string{"once", "a", "tree", "has", "been", "constructed"}, search: "been", exists: true},
		}

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			//	add the input to the binary tree
			want := test.exists
			testBinaryTree := New(func(firstValue, secondValue interface{}) int {
				if firstValue.(string) < secondValue.(string) {
					return -1
				}
				if firstValue.(string) > secondValue.(string) {
					return 1
				}

				return 0
			})

			for _, item := range test.input {
				testBinaryTree.Add(item)
			}

			//	find a value in the binary tree
			got := testBinaryTree.Find(test.search)

			//	check the result
			if want != got {
				t.Errorf("failed finding node in the binary tree: expected: %t result: %t", want, got)
			}
		}
	})
}

//	TestBinaryTreePreorderTraversal test cases
func TestBinaryTreePreorderTraversal(t *testing.T) {

	t.Run(">>> preorder traversal of binary trees", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			scenario string
			input    []string
			traverse string
		}{
			{scenario: "Single element tree", input: []string{"once"}, traverse: "once"},
			{scenario: "Double elements tree", input: []string{"once", "a"}, traverse: "once a"},
			{scenario: "Multiple elements tree", input: []string{"once", "a", "tree", "has", "been", "constructed"}, traverse: "once a has been constructed tree"},
		}

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			//	add the input to the binary tree
			want := test.traverse
			testBinaryTree := New(func(firstValue, secondValue interface{}) int {
				if firstValue.(string) < secondValue.(string) {
					return -1
				}
				if firstValue.(string) > secondValue.(string) {
					return 1
				}

				return 0
			})

			for _, item := range test.input {
				testBinaryTree.Add(item)
			}

			//	perform a preorder traversal of the tree and concatenate the result
			traversalResult := testBinaryTree.PreorderTraversal()
			got := ""

			for _, value := range traversalResult {
				if len(got) > 0 {
					got += " "
				}
				got += value.(string)
			}

			//	check the result
			if want != got {
				t.Errorf("failed traversing the binary tree: expected: %s result: %s", want, got)
			}
		}
	})
}
