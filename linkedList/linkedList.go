package linkedlist

import (
	"errors"
	"fmt"
)

var ElementNotFoundError = errors.New("element not found")

// Definición de tipos

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
}

// Constructores

func newNode[T comparable](value T) *Node[T] {
	return &Node[T]{Val: value}
}

func New[T comparable](args ...T) *LinkedList[T] {
	if len(args) == 0 {
		return &LinkedList[T]{}
	} else {
		l := &LinkedList[T]{
			Head: newNode(args[0]),
		}

		currentNode := l.Head
		for _, v := range args[1:] {
			currentNode.Next = newNode(v)
			currentNode = currentNode.Next
		}

		return l
	}
}

func (l LinkedList[T]) ForEach(fn func(v T)) {
	currentNode := l.Head

	for currentNode != nil {
		fn(currentNode.Val)
		currentNode = currentNode.Next
	}
}

func (l *LinkedList[T]) ForEachMut(fn func(v *T)) {
	currentNode := l.Head

	for currentNode != nil {
		fn(&currentNode.Val)
		currentNode = currentNode.Next
	}
}

func (l *LinkedList[T]) Push(v T) {
	if l.Head == nil {

		l.Head = newNode(v)
		return

	}

	currentNode := l.Head

	for {

		if currentNode.Next == nil {
			currentNode.Next = newNode(v)
			break
		}

		currentNode = currentNode.Next

	}
}

func (l *LinkedList[T]) InsertAtBeginning(v T) {
	oldHead := l.Head

	l.Head = newNode(v)
	l.Head.Next = oldHead
}

func (l *LinkedList[T]) InsertAfterValue(target, v T) error {
	currentNode := l.Head

	for currentNode != nil {

		if currentNode.Val == target {
			oldNext := currentNode.Next

			currentNode.Next = newNode(v)
			currentNode.Next.Next = oldNext

			return nil
		}

		currentNode = currentNode.Next

	}

	return ElementNotFoundError
}

func (l *LinkedList[T]) RemoveHead() {
	l.Head = l.Head.Next
}

func (l *LinkedList[T]) RemoveByValue(target T) error {
	currentNode := l.Head

	for currentNode.Next != nil {
		if node := currentNode.Next; node.Val == target {

			currentNode.Next = currentNode.Next.Next
			return nil

		}
	}

	return ElementNotFoundError
}

func (l LinkedList[T]) Size() int {
	count := 0

	l.ForEach(func(v T) {
		count++
	})

	return count
}

func (l LinkedList[T]) Print() {
	var s []byte

	l.ForEach(func(v T) {
		s = fmt.Append(s, v, "-> ")
	})

	fmt.Println(string(s))
}

func (l LinkedList[T]) ToSlice() []T {
	s := []T{}
	l.ForEach(func(v T) {
		s = append(s, v)
	})
	return s
}
