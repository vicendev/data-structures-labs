package linkedlist

import (
	"errors"
)

var (
	ErrEmpty         = errors.New("double linked list is empty")
	ErrIndexNotFound = errors.New("double linked list index not found")
	ErrNotFound      = errors.New("double linked list node not found")
)

type DoubleLinkedListService interface {
	// Insertion Methods
	AddFirst(value string)
	AddLast(value string)
	InsertAt(index int, value string) error
	InsertAfter(searchValue string, newValue string) error

	// Deletion Methods
	RemoveFirst() (string, error)
	RemoveLast() (string, error)
	RemoveAt(index int) (string, error)
	Remove(value string) (string, error)
	Clear()

	// Accessibility Methods
	GetAt(index int) (*Node, error)
	Find(value string) (*Node, error)
	IndexOf(value string) (int, error)
}

type Node struct {
	Value string
	Next  *Node
	Prev  *Node
}

type linkedList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleLinkedList() DoubleLinkedListService {
	return &linkedList{}
}

func (l *linkedList) AddFirst(newValue string) {
	newNode := &Node{Value: newValue}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Next = l.head
		newNode.Next.Prev = newNode
		l.head = newNode
	}

	l.size++
}

// AddLast implements DoubleLinkedListService.
func (l *linkedList) AddLast(newValue string) {
	newNode := &Node{Value: newValue}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Prev = l.tail
		l.tail.Next = newNode
		l.tail = newNode
	}

	l.size++
}

// Clear implements DoubleLinkedListService.
func (l *linkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Find implements DoubleLinkedListService.
func (l *linkedList) Find(value string) (*Node, error) {
	if err := l.validateEmpty(); err != nil {
		return nil, err
	}

	foundNode, _, _ := l.findByValue(value)

	if foundNode != nil {
		return foundNode, nil
	}

	return nil, ErrNotFound
}

func (l *linkedList) GetAt(nodeIndex int) (*Node, error) {
	if err := l.validateEmpty(); err != nil {
		return nil, err
	}

	foundNode, _ := l.findByIndex(nodeIndex)

	if foundNode != nil {
		return foundNode, nil
	}

	return nil, ErrIndexNotFound
}

func (l *linkedList) IndexOf(value string) (int, error) {
	if err := l.validateEmpty(); err != nil {
		return -1, err
	}

	_, _, index := l.findByValue(value)

	if index >= 0 {
		return index, nil
	}

	return -1, nil
}

func (l *linkedList) InsertAfter(searchValue string, newValue string) error {
	if err := l.validateEmpty(); err != nil {
		return err
	}

	foundNode, _, _ := l.findByValue(searchValue)

	if foundNode == nil {
		return ErrNotFound
	}

	newNode := &Node{Value: newValue, Next: foundNode.Next, Prev: foundNode}
	foundNode.Next = newNode

	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}

	if foundNode == l.tail {
		l.tail = newNode
	}

	l.size++

	return nil
}

// InsertAt implements DoubleLinkedListService.
func (l *linkedList) InsertAt(nodeIndex int, newValue string) error {
	if err := l.validateIndex(nodeIndex, true); err != nil {
		return err
	}

	newNode := &Node{Value: newValue}

	if nodeIndex == 0 {
		newNode.Next = l.head
		l.head = newNode

		if l.tail == nil {
			l.tail = newNode
		}

		l.size++
		return nil
	}

	if nodeIndex == l.size {
		newNode.Prev = l.tail
		l.tail.Next = newNode
		l.tail = newNode

		l.size++

		return nil
	}

	foundNode, prevNode := l.findByIndex(nodeIndex)
	if foundNode != nil {
		newNode.Next = foundNode
		foundNode.Prev = newNode

		newNode.Prev = prevNode
		prevNode.Next = newNode

		l.size++
		return nil
	}

	return ErrIndexNotFound
}

// Remove implements DoubleLinkedListService.
func (l *linkedList) Remove(searchValue string) (string, error) {
	if err := l.validateEmpty(); err != nil {
		return "", err
	}

	foundNode, prevNode, _ := l.findByValue(searchValue)

	if foundNode == nil {
		return "", ErrNotFound
	}

	removedValue := foundNode.Value
	l.unlinkNode(foundNode, prevNode)

	l.size--

	return removedValue, nil
}

// RemoveAt implements DoubleLinkedListService.
func (l *linkedList) RemoveAt(nodeIndex int) (string, error) {
	if err := l.validateEmpty(); err != nil {
		return "", err
	}

	if err := l.validateIndex(nodeIndex, false); err != nil {
		return "", err
	}

	foundNode, prevNode := l.findByIndex(nodeIndex)

	removedValue := foundNode.Value
	l.unlinkNode(foundNode, prevNode)

	l.size--

	return removedValue, nil
}

// RemoveFirst implements DoubleLinkedListService.
func (l *linkedList) RemoveFirst() (string, error) {
	if err := l.validateEmpty(); err != nil {
		return "", err
	}

	removedValue := l.head.Value
	l.unlinkNode(l.head, nil)

	l.size--
	return removedValue, nil
}

// RemoveLast implements DoubleLinkedListService.
func (l *linkedList) RemoveLast() (string, error) {
	if err := l.validateEmpty(); err != nil {
		return "", err
	}

	removedValue := l.tail.Value
	l.unlinkNode(l.tail, nil)

	l.size--
	return removedValue, nil
}

/* Private Functions */

func (l *linkedList) findByValue(value string) (*Node, *Node, int) {
	index := 0
	currentNode := l.head
	var prevNode *Node = nil
	for currentNode != nil {

		if currentNode.Value == value {
			return currentNode, prevNode, index
		}

		prevNode = currentNode
		currentNode = currentNode.Next
		index++
	}

	return nil, nil, -1
}

func (l *linkedList) findByIndex(index int) (*Node, *Node) {
	currentNode := l.head
	var prevNode *Node = nil

	for i := 0; i < index; i++ {
		if currentNode == nil {
			return nil, nil
		}
		prevNode = currentNode
		currentNode = currentNode.Next
	}

	return currentNode, prevNode
}

func (l *linkedList) unlinkNode(currentNode *Node, prevNode *Node) {
	if currentNode == l.head && currentNode == l.tail {
		l.head = nil
		l.tail = nil

		return
	}

	if currentNode == l.head {
		l.head = currentNode.Next
		l.head.Prev = nil
		return
	}

	if currentNode == l.tail {
		l.tail = prevNode
		if l.tail != nil {
			l.tail.Next = nil
			l.tail.Prev = prevNode.Prev
		}
		return
	}

	prevNode.Next = currentNode.Next
	prevNode.Next.Prev = currentNode.Prev
}

/* Validations */

func (l *linkedList) validateEmpty() error {
	if l.head == nil {
		return ErrEmpty
	}

	return nil
}

func (l *linkedList) validateIndex(index int, allowEqualSize bool) error {

	if index < 0 {
		return ErrIndexNotFound
	}
	if allowEqualSize && index > l.size {
		return ErrIndexNotFound
	}
	if index > l.size {
		return ErrIndexNotFound
	}
	return nil
}
