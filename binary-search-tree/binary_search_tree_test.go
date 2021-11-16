package binary_search_tree

import (
	"log"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	tree := &Node{data: 1}
	tree.insert(1)
	log.Printf("tree is: %v", tree)
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
