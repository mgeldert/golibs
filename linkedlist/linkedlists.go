package linkedlists

import (
	"fmt"
)

type LinkedListNode[T any] struct {
	next *LinkedListNode[T]
	previous *LinkedListNode[T]
	Data T
}

func (lln *LinkedListNode[T]) String() string {
	return fmt.Sprintf(
		"[%p] prev:%0 11p | %-30v | next:%0 11p",
		lln, lln.previous, lln.Data, lln.next,
	)
}

type LinkedList[T any] struct {
	head *LinkedListNode[T]
	tail *LinkedListNode[T]
	length int
	limit int
}

func NewLinkedList[T any](limit int) LinkedList[T] {
	return LinkedList[T]{
		head: nil,
		tail: nil,
		length: 0,
		limit: limit,
	}
}

func (ll LinkedList[T]) Head() *LinkedListNode[T] {
	return ll.head
}

func (ll LinkedList[T]) Tail() *LinkedListNode[T] {
	return ll.tail
}

func (ll LinkedList[T]) Length() int {
	return ll.length
}

func (ll LinkedList[T]) Limit() int {
	return ll.limit
}

func(ll LinkedList[T]) Get(index int) (T, error) {
	var zeroValue T
	if index >= ll.length {
		return zeroValue, fmt.Errorf("Index out of range")
	}

	current := ll.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.Data, nil
}

func(ll *LinkedList[T]) Append(value T) error {
	if ll.limit > 0 && ll.length == ll.limit {
		return fmt.Errorf("List is full (max size %d)", ll.limit)
	}
	ll.length++

	newNode := &LinkedListNode[T]{
		next: nil,
		previous: nil,
		Data: value,
	}

	if ll.head == nil {
		ll.head = newNode
		return nil
	}
	
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	newNode.previous = current
	current.next = newNode
	ll.tail = newNode
	return nil
}

func(ll *LinkedList[T]) InsertAtIndex(value T, index int) error {
	if ll.limit > 0 && ll.length == ll.limit {
		return fmt.Errorf("List is full (max size %d)", ll.limit)
	} else if index > ll.length {
		return fmt.Errorf("Index out of range")
	}
	
	newNode := &LinkedListNode[T]{
		next: nil,
		previous: nil,
		Data: value,
	}

	var previous *LinkedListNode[T] = nil
	current := &ll.head
	for i := 0; i < index; i++ {
		if *current == nil {
			return fmt.Errorf("Index out of range")
		}
		previous = *current
		current = &((*current).next)
	}
	
	ll.length++
	newNode.previous = previous
	newNode.next = *current
	if newNode.next == nil {
		ll.tail = newNode
	} else {
		(*current).previous = newNode
	}
	*current = newNode

	return nil
}

func(ll *LinkedList[T]) DeleteIndex(index int) error {
	if ll.length == 0 {
		return fmt.Errorf("List is empty")
	} else if index >= ll.length {
		return fmt.Errorf("Index out of range")
	}

	var previous *LinkedListNode[T] = nil
	current := &ll.head
	for i := 0; i < index; i++ {
		if *current == nil {
			return fmt.Errorf("Index out of range")
		}
		previous = *current
		current = &((*current).next)
	}

	if (*current).next == nil {
		ll.tail = (*current).previous
		*current = nil
	} else if previous == nil {
		(*current).next.previous = nil
		*current = (*current).next
	} else {
		(*current).next.previous = previous
		*current = (*current).next
	}

	ll.length--
	return nil
}

func(ll LinkedList[T]) Array() []T {
	arr := []T{}
	current := ll.head
	for current != nil {
		arr = append(arr, current.Data)
		current = current.next
	}
	return arr
}

func (ll LinkedList[T]) String() string {
	output := ""
	output += fmt.Sprintf("List length: %7d\n", ll.length)
	output += fmt.Sprintf("Max length:  %7d\n", ll.limit)
	output += fmt.Sprintf("Head: %p\n", ll.head)
	current := ll.head
	for current != nil {
		output += fmt.Sprintf("%v\n", current)
		current = current.next
	}
	output += fmt.Sprintf("Tail: %p\n", ll.tail)

	return output
}