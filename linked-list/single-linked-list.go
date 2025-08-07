package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (head *List[T]) ListAppend(val T) *List[T] {
	v := List[T]{nil, val}
	
	if head == nil {
		return &v
	}
	
	var tail *List[T] = head
	for ; tail.next != nil; {
		tail = tail.next
	}
	
	tail.next = &v
	return head
}

func PrintList[T comparable](head *List[T]) {
	for ;head != nil; {
		fmt.Printf("Val= %v\n", head.val)
		head = head.next
	}
}

func main() {
	var list *List[int]
	list = list.ListAppend( 1)
	list = list.ListAppend( 2)
	list = list.ListAppend( 3)
	PrintList(list)
}

