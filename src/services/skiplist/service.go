package skiplist

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	ErrEmpty    = errors.New("skiplist is empty")
	ErrNotFound = errors.New("skiplist node not found")
)

type SkipListService interface {
	// Insertion Methods
	Insert(key int, value string) (old string, replaced bool)
	Seed() (size int)

	// Accessibility Methods
	Search(key int) (valueFound string, err error)
	Contains(key int) (found bool, err error)

	// Deletion Methods
	Delete(key int) (deletedValue string, err error)
}

type Node struct {
	Key   int
	Value string
	Next  []*Node
}

type skipList struct {
	head        *Node
	maxLevel    int
	level       int
	probability float64
	size        int
}

func NewSkipList() SkipListService {
	skipList := &skipList{
		head:        &Node{},
		maxLevel:    5,
		level:       0,
		probability: 0.5,
		size:        0,
	}
	skipList.head.Next = make([]*Node, skipList.maxLevel)

	return skipList
}

// Contains implements SkipListService.
func (s *skipList) Contains(key int) (found bool, err error) {
	if err := s.validateEmpty(); err != nil {
		return false, err
	}
	baseList := s.traverseList(key)
	currentNode := baseList[0]

	if currentNode.Next[0] != nil && currentNode.Next[0].Key == key {
		return true, nil
	}

	return false, nil
}

// Delete implements SkipListService.
func (s *skipList) Delete(key int) (deletedValue string, err error) {
	if err := s.validateEmpty(); err != nil {
		return "", err
	}
	update := s.traverseList(key)

	currentNode := update[0]
	successorNode := currentNode.Next[0]

	if successorNode == nil || successorNode.Key != key {
		return "", ErrNotFound
	}

	deletedValue = successorNode.Value
	for i := 0; i < len(successorNode.Next); i++ {
		if update[i].Next[i] == successorNode {
			update[i].Next[i] = successorNode.Next[i]
		}
	}

	for s.level > 0 && s.head.Next[s.level] == nil {
		s.level--
	}

	s.size--

	return deletedValue, nil
}

// Insert implements SkipListService.
func (s *skipList) Insert(key int, value string) (oldValue string, replaced bool) {
	update := s.traverseList(key)

	currentNode := update[0]
	successorNode := currentNode.Next[0]

	if successorNode != nil && successorNode.Key == key {
		oldValue = successorNode.Value
		successorNode.Value = value
		return oldValue, true
	}

	height := s.randomLevel()

	newNode := &Node{
		Key:   key,
		Value: value,
		Next:  make([]*Node, height+1),
	}

	if s.level < height {
		for i := s.level + 1; i <= height; i++ {
			update[i] = s.head
		}

		s.level = height
	}

	for i := height; i >= 0; i-- {
		newNode.Next[i] = update[i].Next[i]
		update[i].Next[i] = newNode
	}

	s.size++

	fmt.Println(newNode, "level: ", height)
	return "", false
}

// Search implements SkipListService.
func (s *skipList) Search(key int) (valueFound string, err error) {
	if err := s.validateEmpty(); err != nil {
		return "", err
	}

	baseList := s.traverseList(key)
	currentNode := baseList[0]

	if currentNode.Next[0] != nil && currentNode.Next[0].Key == key {
		return currentNode.Next[0].Value, nil
	}

	return "", ErrNotFound
}

func (s *skipList) Seed() (size int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := 500
	max := 2000

	nums := make([]int, max)

	for i := 0; i < max; i++ {
		nums[i] = i + 1
	}

	r.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	result := nums[:n]
	print(result)
	for i := 0; i < len(result); i++ {
		s.Insert(result[i], "Data "+strconv.Itoa(result[i]))
	}

	return s.size
}

/* Private Methods */
func (s *skipList) randomLevel() int {
	level := 0

	for rand.Float64() < s.probability && level < s.maxLevel-1 {
		level++
	}

	return level
}

func (s *skipList) traverseList(key int) []*Node {
	update := make([]*Node, s.maxLevel)

	currentNode := s.head
	for i := s.level; i >= 0; i-- {
		for currentNode.Next[i] != nil && currentNode.Next[i].Key < key {
			currentNode = currentNode.Next[i]
		}

		update[i] = currentNode
	}

	return update
}

/* Validations */
func (s *skipList) validateEmpty() error {
	if s.size == 0 {
		return ErrEmpty
	}

	return nil
}

/* Utils */
func (n *Node) String() string {
	if n == nil {
		return "<nil>"
	}

	keys := []int{}
	for _, next := range n.Next {
		if next != nil {
			keys = append(keys, next.Key)
		} else {
			keys = append(keys, -1)
		}
	}

	return fmt.Sprintf("Node{Key:%d, Value:%s, NextKeys:%v},", n.Key, n.Value, keys)
}
