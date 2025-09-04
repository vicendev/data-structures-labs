package linkedlist

import (
	"errors"
)

var (
	ErrEmpty         = errors.New("single linked list is empty")
	ErrIndexNotFound = errors.New("single linked list index not found")
	ErrNotFound      = errors.New("single linked list node not found")
)

type SingleLinkedListService interface {
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
	GetAt(index int) (*node, error)
	Find(value string) (*node, error)
	IndexOf(value string) (int, error)
}

type node struct {
	Value string
	Next  *node
}

type linkedList struct {
	head *node
	tail *node
	size int
}

func NewSingleLinkedList() SingleLinkedListService {
	return &linkedList{}
}

// AddFirst implements SingleLinkedListService.
func (l *linkedList) AddFirst(newValue string) {
	newNode := &node{Value: newValue, Next: l.head}

	l.head = newNode

	if l.tail == nil {
		l.tail = newNode
	}

	print(l.head.Value)
	l.size++
}

// AddLast implements SingleLinkedListService.
func (l *linkedList) AddLast(newValue string) {
	newNode := &node{Value: newValue, Next: nil}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}

	l.size++
}

// Clear implements SingleLinkedListService.
func (l *linkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Find implements SingleLinkedListService.
func (l *linkedList) Find(searchValue string) (*node, error) {
	println(l.head.Value)
	if err := l.validateEmpty(); err != nil {
		return nil, err
	}

	foundNode, _, _ := l.findByValue(searchValue)

	if foundNode != nil {
		return foundNode, nil
	}

	return nil, ErrNotFound
}

// GetAt implements SingleLinkedListService.
func (l *linkedList) GetAt(nodeIndex int) (*node, error) {
	if err := l.validateEmpty(); err != nil {
		return nil, err
	}

	foundNode, _ := l.findByIndex(nodeIndex)

	if foundNode != nil {
		return foundNode, nil
	}

	return nil, ErrIndexNotFound
}

// IndexOf implements SingleLinkedListService.
func (l *linkedList) IndexOf(searchValue string) (int, error) {
	if err := l.validateEmpty(); err != nil {
		return -1, err
	}

	_, _, index := l.findByValue(searchValue)

	if index >= 0 {
		return index, nil
	}

	return -1, nil
}

// InsertAt implements SingleLinkedListService.
func (l *linkedList) InsertAt(nodeIndex int, value string) error {
	if err := l.validateIndex(nodeIndex, true); err != nil {
		return err
	}

	newNode := &node{Value: value}

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
		l.tail.Next = newNode
		l.tail = newNode

		l.size++
		return nil
	}

	foundNode, prevNode := l.findByIndex(nodeIndex)
	if foundNode != nil {
		newNode.Next = foundNode
		prevNode.Next = newNode

		l.size++
		return nil
	}

	return ErrIndexNotFound
}

func (l *linkedList) InsertAfter(searchValue string, newValue string) error {
	if err := l.validateEmpty(); err != nil {
		return err
	}

	foundNode, _, _ := l.findByValue(searchValue)

	if foundNode == nil {
		return ErrNotFound
	}

	newNode := &node{Value: newValue, Next: foundNode.Next}
	foundNode.Next = newNode

	if foundNode == l.tail {
		l.tail = newNode
	}

	l.size++

	return nil
}

// Remove implements SingleLinkedListService.
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

// RemoveAt implements SingleLinkedListService.
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

// RemoveFirst implements SingleLinkedListService.
func (l *linkedList) RemoveFirst() (string, error) {
	if err := l.validateEmpty(); err != nil {
		return "", err
	}

	removedValue := l.head.Value
	l.unlinkNode(l.head, nil)

	l.size--
	return removedValue, nil
}

// RemoveLast implements SingleLinkedListService.
func (l *linkedList) RemoveLast() (string, error) {
	if err := l.validateEmpty(); err != nil {
		return "", err
	}

	foundNode, prevNode, _ := l.findLastWithPrev()

	removedValue := foundNode.Value
	l.unlinkNode(foundNode, prevNode)

	l.size--
	return removedValue, nil
}

func (l *linkedList) findByValue(value string) (*node, *node, int) {
	index := 0
	currentNode := l.head
	var prevNode *node = nil
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

func (l *linkedList) findByIndex(index int) (*node, *node) {
	currentNode := l.head
	var prevNode *node = nil

	for i := 0; i < index; i++ {
		if currentNode == nil {
			return nil, nil
		}
		prevNode = currentNode
		currentNode = currentNode.Next
	}

	return currentNode, prevNode
}

func (l *linkedList) findLastWithPrev() (*node, *node, int) {
	index := 0
	currentNode := l.head
	var prevNode *node = nil
	for currentNode != nil {

		if currentNode.Next == nil {
			return currentNode, prevNode, index
		}

		prevNode = currentNode
		currentNode = currentNode.Next
		index++
	}

	return nil, nil, -1
}

func (l *linkedList) unlinkNode(currentNode *node, prevNode *node) {
	if currentNode == l.head && currentNode == l.tail {
		l.head = nil
		l.tail = nil

		return
	}

	if currentNode == l.head {
		l.head = currentNode.Next
		return
	}

	if currentNode == l.tail {
		l.tail = prevNode
		if l.tail != nil {
			l.tail.Next = nil
		}
		return
	}

	prevNode.Next = currentNode.Next
}

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
	if index >= l.size {
		return ErrIndexNotFound
	}
	return nil
}
