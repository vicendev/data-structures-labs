package binarytree

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/thediveo/go-asciitree"
)

var (
	ErrEmpty          = errors.New("binary tree is empty")
	ErrNotFound       = errors.New("binary tree key not found")
	ErrHashNotSupport = errors.New("binary tree fn is not supported")
)

type BinaryTreeService interface {
	// Insertion Methods
	Upsert(key int, value string) (oldValue string, replaced bool)
	Seed() int

	// Accesibility Methods
	Search(key int) (valueFound string, err error)

	// Deletion Methods
	Delete(key int) (deletedValue string, err error)

	// Utility Methods
	Reset()
	DebugTree()
}

type Node struct {
	Key   int
	Value string
	Left  *Node
	Right *Node
}

type binarytree struct {
	root *Node
	size int
}

func NewBinaryTree() BinaryTreeService {
	return &binarytree{
		size: 0,
	}
}

// Delete implements BinaryTreeService.
func (b *binarytree) Delete(key int) (deletedValue string, err error) {
	if b.size == 0 {
		return "", ErrEmpty
	}

	b.root, deletedValue, err = b.deleteNode(key, b.root)

	if err != nil {
		return "", err
	}

	b.size--
	return deletedValue, nil
}

func (b *binarytree) DebugTree() {
	b.PrintASCII()
}

// Reset implements BinaryTreeService.
func (b *binarytree) Reset() {
	b.root = nil
	b.size = 0
}

// Search implements BinaryTreeService.
func (b *binarytree) Search(key int) (valueFound string, err error) {

	if b.size == 0 {
		return "", ErrEmpty
	}

	node := b.findNode(key, b.root)

	if node != nil {
		return node.Value, nil
	}

	return "", ErrNotFound
}

// Seed implements BinaryTreeService.
func (b *binarytree) Seed() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := 20
	max := 200

	nums := make([]int, max)

	for i := 0; i < max; i++ {
		nums[i] = i + 1
	}

	r.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	result := nums[:n]

	for i := 0; i < len(result); i++ {
		resultString := strconv.Itoa(result[i])
		b.Upsert(result[i], "Data "+resultString)
	}

	return b.size
}

// Upsert implements BinaryTreeService.
func (b *binarytree) Upsert(key int, value string) (oldValue string, replaced bool) {

	newNode := &Node{
		Key:   key,
		Value: value,
	}
	if b.root == nil {
		b.root = newNode

		b.size++

		return "", false
	}

	node := b.findInsertPosition(key, b.root)

	if key == node.Key {
		oldValue = node.Value
		node.Value = value
		return oldValue, true
	}

	if key < node.Key {
		node.Left = newNode
	} else {
		node.Right = newNode
	}

	b.size++

	return "", false
}

func (b *binarytree) findInsertPosition(key int, node *Node) *Node {

	if node == nil {
		return nil
	}

	if key < node.Key {
		if node.Left == nil {
			return node
		}
		return b.findInsertPosition(key, node.Left)
	}

	if key > node.Key {
		if node.Right == nil {
			return node

		}
		return b.findInsertPosition(key, node.Right)
	}

	return node
}

func (b *binarytree) findNode(key int, node *Node) *Node {
	if node == nil {
		return nil
	}

	if key == node.Key {
		return node
	} else if key < node.Key {
		return b.findNode(key, node.Left)
	} else if key > node.Key {
		return b.findNode(key, node.Right)
	}

	return nil
}

func (b *binarytree) deleteNode(key int, node *Node) (*Node, string, error) {
	if node == nil {
		return nil, "", ErrNotFound
	}

	if key < node.Key {
		newLeft, value, err := b.deleteNode(key, node.Left)
		node.Left = newLeft
		return node, value, err
	}

	if key > node.Key {
		newRight, value, err := b.deleteNode(key, node.Right)
		node.Right = newRight
		return node, value, err
	}

	deletedValue := node.Value

	if node.Left == nil && node.Right == nil {
		return nil, deletedValue, nil
	}

	if node.Left == nil {
		return node.Right, deletedValue, nil
	}

	if node.Right == nil {
		return node.Left, deletedValue, nil
	}

	succesorNode := b.findMinimumNode(node.Right)

	node.Key = succesorNode.Key
	node.Value = succesorNode.Value

	node.Right, _, _ = b.deleteNode(node.Key, node.Right)

	return node, deletedValue, nil
}

func (b *binarytree) findMinimumNode(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

/* PRINT TREE*/

type printableNode struct {
	Label    string           `asciitree:"label"`
	Children []*printableNode `asciitree:"children"`
}

func makePrintable(n *Node) *printableNode {
	if n == nil {
		return nil
	}
	p := &printableNode{
		Label: fmt.Sprintf("%d:%s", n.Key, n.Value),
	}

	if n.Left != nil {
		p.Children = append(p.Children, makePrintable(n.Left))
	}
	if n.Right != nil {
		p.Children = append(p.Children, makePrintable(n.Right))
	}
	return p
}

func (b *binarytree) PrintASCII() {
	if b.root == nil {
		fmt.Println("(árbol vacío)")
		return
	}
	fmt.Println(asciitree.RenderFancy(makePrintable(b.root)))
}
