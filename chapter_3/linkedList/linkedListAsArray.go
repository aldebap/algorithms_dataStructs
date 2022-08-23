////////////////////////////////////////////////////////////////////////////////
//	linkedListAsArray.go  -  Ago-22-2022  -  aldebap
//
//	Linked List as Array Data Structure
////////////////////////////////////////////////////////////////////////////////

package main

type linkedListAsArrayElement struct {
	value interface{}
	next  int
}

type LinkedListAsArray struct {
	elements []linkedListAsArrayElement
	iterator int
}

//	NewLinkedListAsArray create a new linked list
func NewLinkedListAsArray() GenericLinkedList {
	return &LinkedListAsArray{
		elements: make([]linkedListAsArrayElement, 0),
		iterator: -1,
	}
}

//	AddElement add a new element to the linked list
func (l *LinkedListAsArray) AddElement(value interface{}) {
	lastElement := len(l.elements) - 1

	l.elements = append(l.elements, linkedListAsArrayElement{
		value: value,
		next:  -1,
	})

	//	make the last element to point to the new one
	if lastElement != -1 {
		l.elements[lastElement].next = len(l.elements) - 1
	}
}

//	First point the iterator to the first element of the linked list
func (l *LinkedListAsArray) First() {
	l.iterator = 0
}

//	Next return the element poited by the iterator, and move it to the next one
func (l *LinkedListAsArray) Next() interface{} {
	if l.iterator == -1 || l.iterator >= len(l.elements) {
		return nil
	}

	value := l.elements[l.iterator].value
	l.iterator = l.elements[l.iterator].next

	return value
}
