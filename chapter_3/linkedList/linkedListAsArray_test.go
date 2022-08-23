////////////////////////////////////////////////////////////////////////////////
//	linkedListAsArray_test.go  -  Ago-20-2022  -  aldebap
//
//	Test cases for the Linked List as Array Data Structure
////////////////////////////////////////////////////////////////////////////////

package main

import "testing"

//	TestLinkedListAsArray linked list as array test cases
func TestLinkedListAsArray(t *testing.T) {
	test_genericLinkedList(t, "Insert and iterate through linked list as array", NewLinkedListAsArray)
}
