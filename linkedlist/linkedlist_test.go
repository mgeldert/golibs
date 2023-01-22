package linkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	ll = NewLinkedList[string](5)
)

func TestAppendToLinkedList(t *testing.T) {
	var err error
	var value string

	var appendTests = []struct{
		newValue string
		expectedList []string
		expectedErr error
	}{
		{"foo", []string{"foo"}, nil},
		{"bar", []string{"foo", "bar"}, nil},
		{"bin", []string{"foo", "bar", "bin"}, nil},
		{"bash", []string{"foo", "bar", "bin", "bash"}, nil},
		{"woo", []string{"foo", "bar", "bin", "bash", "woo"}, nil},
		{"baz", []string{"foo", "bar", "bin", "bash", "woo"}, fmt.Errorf("List is full (max size 5)")},
	}

	for i, tt := range appendTests {
		if err = ll.Append(tt.newValue); err != nil {
			if err.Error() != tt.expectedErr.Error() {
				t.Errorf("Unexpected error: expected %#v, got %#v", tt.expectedErr, err)
			}
		}

		if tt.expectedErr == nil {
			if value, err = ll.Get(i); err != nil {
				t.Errorf("Failed to retrieve a valid entry")
			} else if value != tt.newValue {
				t.Errorf("Invalid error: expected '%s', got '%s'", tt.newValue, value)
			}
		}

		listArray := ll.Array()
		if !reflect.DeepEqual(listArray, tt.expectedList) {
			t.Errorf("Array mismatch: expected %v, got %v", tt.expectedList, listArray)
		}
	}
}

func TestDeleteFromLinkedList(t *testing.T) {
	var deleteTests = []struct{
		index int
		expectedList []string
		expectedErr error
	}{
		{2, []string{"foo", "bar", "bash", "woo"}, nil},
		{0, []string{"bar", "bash", "woo"}, nil},
		{2, []string{"bar", "bash"}, nil},
		{2, []string{"bar", "bash"}, fmt.Errorf("Index out of range")},
	}

	for _, tt := range deleteTests {
		if err := ll.DeleteIndex(tt.index); err != nil {
			if err.Error() != tt.expectedErr.Error() {
				t.Errorf("Unexpected error: expected %#v, got %#v", tt.expectedErr, err)
			}
		}
		
		listArray := ll.Array()
		if !reflect.DeepEqual(listArray, tt.expectedList) {
			t.Errorf("Array mismatch: expected %v, got %v", tt.expectedList, listArray)
		}
	}
}

func TestInsertAtIndexOfLinkedList(t *testing.T) {
	var insertTests = []struct{
		index int
		value string
		expectedList []string
		expectedErr error
	}{
		{3, "outofrange", []string{"bar", "bash"}, fmt.Errorf("Index out of range")},
		{0, "beginning", []string{"beginning", "bar", "bash"}, nil},
		{3, "end", []string{"beginning", "bar", "bash", "end"}, nil},
		{2, "middle", []string{"beginning", "bar", "middle", "bash", "end"}, nil},
		{2, "no", []string{"beginning", "bar", "middle", "bash", "end"}, fmt.Errorf("List is full (max size 5)")},
	}

	for _, tt := range insertTests {
		if err := ll.InsertAtIndex(tt.value, tt.index); err != nil {
			if err.Error() != tt.expectedErr.Error() {
				t.Errorf("Unexpected error: expected %#v, got %#v", tt.expectedErr, err)
			}
		}
		
		listArray := ll.Array()
		if !reflect.DeepEqual(listArray, tt.expectedList) {
			t.Errorf("Array mismatch: expected %v, got %v", tt.expectedList, listArray)
		}
	}
}

func TestSetOfLinkedList(t *testing.T) {
	var setTests = []struct{
		index int
		value string
		expectedList []string
		expectedErr error
	}{
		{2, "MIDDLE", []string{"beginning", "bar", "MIDDLE", "bash", "end"}, nil},
		{5, "outofrange", []string{"beginning", "bar", "MIDDLE", "bash", "end"}, fmt.Errorf("Index out of range")},
	}

	for _, tt := range setTests {
		if err := ll.Set(tt.value, tt.index); err != nil {
			if err.Error() != tt.expectedErr.Error() {
				t.Errorf("Unexpected error: expected %#v, got %#v", tt.expectedErr, err)
			}
		}
		
		listArray := ll.Array()
		if !reflect.DeepEqual(listArray, tt.expectedList) {
			t.Errorf("Array mismatch: expected %v, got %v", tt.expectedList, listArray)
		}
	}
}

func TestLengthAndLimit(t *testing.T) {
	if ll.Length() != 5 {
		t.Errorf("Expected length mismatch: expected 5, got %d", ll.Length())
	}
	if ll.Limit() != 5 {
		t.Errorf("Limit mismatch: expected 5, got %d", ll.Limit())
	}
}