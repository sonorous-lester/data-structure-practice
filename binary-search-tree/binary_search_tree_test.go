package binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	tree := &Node{data: 1}
	tree.insert(1)
	log.Printf("tree is: %v", tree)
}

func TestBinarySearchTree_Find(t *testing.T) {
	tree := &Node{data: 25}
	tree.insert(2)
	tree.insert(39)
	tree.insert(54)
	tree.insert(14)
	log.Printf("tree is: %v", tree)

	expected := true
	actual := tree.find(14)
	assert.Equal(t, expected, actual)
}

func TestBinarySearchTree_NotFind(t *testing.T) {
	tree := &Node{data: 25}
	tree.insert(2)
	tree.insert(39)
	tree.insert(54)
	tree.insert(14)
	log.Printf("tree is: %v", tree)

	expected := false
	actual := tree.find(15)
	assert.Equal(t, expected, actual)
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (node *Node) insert(value int) {
	switch {
	case value > node.data:
		{
			if node.right == nil {
				node.right = &Node{data: value}
			}else {
				node.right.insert(value)
			}
		}
	case value < node.data:
		{
			if node.left == nil {
				node.left = &Node{data: value}
			}else {
				node.left.insert(value)
			}
		}
	default:
		{
			log.Printf("value: %d already in the tree", value)
		}
	}
}
func (node *Node) find(target int) (ok bool) {
	if node == nil {
		return false
	}
	switch {
	case target > node.data:
		{
			if node.right == nil {
				return false
			}
			return node.right.find(target)
		}
	case target < node.data:
		{
			if node.left == nil {
				return false
			}
			return node.left.find(target)
		}
	case target == node.data:
		return true
	default:
		return false
	}
}