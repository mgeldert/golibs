package linkedlist

import (
	"fmt"
)

type LinkedListNode[T any] struct {
	next *LinkedListNode[T]
	previous *LinkedListNode[T]
	data T
}

func (lln LinkedListNode[T]) Data() T {
	return lln.data
}

func (lln *LinkedListNode[T]) String() string {
	return fmt.Sprintf(
		"[%p] prev:%011p | %-30v | next:%011p",
		lln, lln.previous, lln.data, lln.next,
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

func (ll *LinkedList[T]) getIndexPointer(index int) (*LinkedListNode[T], error) {

	current := ll.head
	for i := 0; i < index; i++ {
		if current == nil {
			return nil, fmt.Errorf("Unexpected nil pointer - list broken")
		}
		current = current.next
	}

	return current, nil
}

func (ll *LinkedList[T]) getIndexPointerPointer(index int) (**LinkedListNode[T], *LinkedListNode[T], error) {

	var previous *LinkedListNode[T] = nil
	current := &ll.head
	for i := 0; i < index; i++ {
		if *current == nil {
			return nil, nil, fmt.Errorf("Unexpected nil pointer - list broken")
		}
		previous = *current
		current = &((*current).next)
	}

	return current, previous, nil
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

	if pointer, err := ll.getIndexPointer(index); err != nil {
		return zeroValue, err
	} else {
		return pointer.data, nil
	}
}

func(ll LinkedList[T]) Set(value T, index int) error {
	if index >= ll.length {
		return fmt.Errorf("Index out of range")
	}

	if pointer, err := ll.getIndexPointer(index); err != nil {
		return err
	} else {
		(*pointer).data = value
	}

	return nil
}

func(ll *LinkedList[T]) InsertAtIndex(value T, index int) error {
	if ll.limit > 0 && ll.length == ll.limit {
		return fmt.Errorf("List is full (max size %d)", ll.limit)
	} else if index > ll.length {
		return fmt.Errorf("Index out of range")
	}

	current, previous, err := ll.getIndexPointerPointer(index)
	if err != nil {
		return err
	}

	newNode := &LinkedListNode[T]{data: value,}
	if *current != nil {
		newNode.next = *current
		newNode.previous = (*current).previous
		(*current).previous = newNode
	} else {
		newNode.previous = previous
		ll.tail = newNode
	}
	*current = newNode

	ll.length++
	return nil
}

func(ll *LinkedList[T]) Append(value T) error {
	if ll.limit > 0 && ll.length == ll.limit {
		return fmt.Errorf("List is full (max size %d)", ll.limit)
	}

	newNode := &LinkedListNode[T]{previous: ll.tail, data: value,}
	if ll.tail != nil {
		(*ll.tail).next = newNode
	}
	if ll.head == nil {
		ll.head = newNode
	}
	ll.tail = newNode

	ll.length++
	return nil
}

func(ll *LinkedList[T]) DeleteIndex(index int) error {
	if ll.length == 0 {
		return fmt.Errorf("List is empty")
	} else if index >= ll.length {
		return fmt.Errorf("Index out of range")
	}

	current, previous, err := ll.getIndexPointerPointer(index)
	if err != nil {
		return err
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
		arr = append(arr, current.data)
		current = current.next
	}
	return arr
}

func (ll LinkedList[T]) String() string {
	output := ""
	output += fmt.Sprintf("List length: %8d\n", ll.length)
	output += fmt.Sprintf("Max. length: %8d\n", ll.limit)
	output += fmt.Sprintf("Head: %p\n", ll.head)
	current := ll.head
	for current != nil {
		output += fmt.Sprintf("     %v\n", current)
		current = current.next
	}
	output += fmt.Sprintf("Tail: %p\n", ll.tail)

	return output
}