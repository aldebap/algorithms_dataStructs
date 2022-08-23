////////////////////////////////////////////////////////////////////////////////
//	linkedList.go  -  Ago-20-2022  -  aldebap
//
//	Linked List Data Structure
////////////////////////////////////////////////////////////////////////////////

package main

type GenericLinkedList interface {
	AddElement(value interface{})
	First()
	Next() interface{}
}

type linkedListElement struct {
	value interface{}
	next  *linkedListElement
}

type LinkedList struct {
	head     *linkedListElement
	iterator *linkedListElement
}

//	NewLinkedList create a new linked list
func NewLinkedList() GenericLinkedList {
	return &LinkedList{
		head:     nil,
		iterator: nil,
	}
}

//	AddElement add a new element to the linked list
func (l *LinkedList) AddElement(value interface{}) {
	newElement := &linkedListElement{
		value: value,
		next:  nil,
	}

	if l.head == nil {
		l.head = newElement
	} else {
		//	iterate to the last element of the list
		iterator := l.head
		for {
			if iterator.next == nil {
				break
			}
			iterator = iterator.next
		}

		//	make the last element to point to the new one
		iterator.next = newElement
	}
}

//	First point the iterator to the first element of the linked list
func (l *LinkedList) First() {
	l.iterator = l.head
}

//	Next return the element poited by the iterator, and move it to the next one
func (l *LinkedList) Next() interface{} {
	if l.iterator == nil {
		return nil
	}

	value := l.iterator.value
	l.iterator = l.iterator.next

	return value
}
