////////////////////////////////////////////////////////////////////////////////
//	binaryTree.go  -  Ago-27-2022  -  aldebap
//
//	Binary Tree Data Structure
////////////////////////////////////////////////////////////////////////////////

package binaryTree

import (
	"github.com/aldebap/algorithms_dataStructs/chapter_3/queue"
	"github.com/aldebap/algorithms_dataStructs/chapter_3/stack"
)

type BinaryTree interface {
	Add(value interface{})
	Find(value interface{}) bool

	PreorderTraversal() []interface{}
	InorderTraversal() []interface{}
	PostorderTraversal() []interface{}
	LevelorderTraversal() []interface{}
}

type treeNode struct {
	value interface{}
	left  *treeNode
	right *treeNode
}

type BinaryTreeRoot struct {
	root          *treeNode
	compareValues func(firstValue, secondValue interface{}) int
}

//	New create a new binary tree
func New(compareValues func(firstValue, secondValue interface{}) int) BinaryTree {
	return &BinaryTreeRoot{
		root:          nil,
		compareValues: compareValues,
	}
}

//	Add add a new node to the binary tree keeping it ordered
func (t *BinaryTreeRoot) Add(value interface{}) {
	newNode := &treeNode{
		value: value,
		left:  nil,
		right: nil,
	}

	if t.root == nil {
		t.root = newNode
	} else {
		iterator := t.root

		//	use the compareValues function to insert a new in such a way that, for each node,
		//	the left node is lower than the node, and the right node is bigger
		for {
			if t.compareValues(value, iterator.value) < 0 {
				if iterator.left == nil {
					iterator.left = newNode
					break
				}

				iterator = iterator.left
			} else {
				if iterator.right == nil {
					iterator.right = newNode
					break
				}

				iterator = iterator.right
			}
		}
	}
}

//	Find search the tree and return true if the value is a node of the tree
func (t *BinaryTreeRoot) Find(value interface{}) bool {
	if t.root == nil {
		return false
	}

	iterator := t.root
	for {
		if t.compareValues(value, iterator.value) == 0 {
			return true
		}

		if t.compareValues(value, iterator.value) < 0 {
			if iterator.left == nil {
				return false
			}

			iterator = iterator.left
		} else {
			if iterator.right == nil {
				return false
			}

			iterator = iterator.right
		}
	}
}

const (
	visitNode         = 1
	visitLeftSubtree  = 2
	visitRightSubtree = 3
)

type traversalOperation struct {
	node      *treeNode
	operation uint8
}

//	PreorderTraversal perform a preorder traversal of the tree and return the nodes in a slice
func (t *BinaryTreeRoot) PreorderTraversal() []interface{} {
	traversalResult := make([]interface{}, 0)
	search := stack.New()

	//	preorder means, visit the root, then the left subtree, and then the right subtree
	search.Push(&traversalOperation{node: t.root, operation: visitRightSubtree})
	search.Push(&traversalOperation{node: t.root, operation: visitLeftSubtree})
	search.Push(&traversalOperation{node: t.root, operation: visitNode})

	for {
		if search.IsEmpty() {
			break
		}

		traversalOp := search.Pop().(*traversalOperation)

		switch traversalOp.operation {
		case visitNode:
			traversalResult = append(traversalResult, traversalOp.node.value)

		case visitLeftSubtree:
			if traversalOp.node.left != nil {
				search.Push(&traversalOperation{node: traversalOp.node.left, operation: visitRightSubtree})
				search.Push(&traversalOperation{node: traversalOp.node.left, operation: visitLeftSubtree})
				search.Push(&traversalOperation{node: traversalOp.node.left, operation: visitNode})
			}

		case visitRightSubtree:
			if traversalOp.node.right != nil {
				search.Push(&traversalOperation{node: traversalOp.node.right, operation: visitRightSubtree})
				search.Push(&traversalOperation{node: traversalOp.node.right, operation: visitLeftSubtree})
				search.Push(&traversalOperation{node: traversalOp.node.right, operation: visitNode})
			}
		}
	}

	return traversalResult
}

//	InorderTraversal perform a inorder traversal of the tree and return the nodes in a slice
func (t *BinaryTreeRoot) InorderTraversal() []interface{} {
	traversalResult := make([]interface{}, 0)
	search := stack.New()

	//	inorder means, visit the left subtree, then root, and then the right subtree
	search.Push(&traversalOperation{node: t.root, operation: visitRightSubtree})
	search.Push(&traversalOperation{node: t.root, operation: visitNode})
	search.Push(&traversalOperation{node: t.root, operation: visitLeftSubtree})

	for {
		if search.IsEmpty() {
			break
		}

		traversalOp := search.Pop().(*traversalOperation)

		switch traversalOp.operation {
		case visitNode:
			traversalResult = append(traversalResult, traversalOp.node.value)

		case visitLeftSubtree:
			if traversalOp.node.left != nil {
				search.Push(&traversalOperation{node: traversalOp.node.left, operation: visitRightSubtree})
				search.Push(&traversalOperation{node: traversalOp.node.left, operation: visitNode})
				search.Push(&traversalOperation{node: traversalOp.node.left, operation: visitLeftSubtree})
			}

		case visitRightSubtree:
			if traversalOp.node.right != nil {
				search.Push(&traversalOperation{node: traversalOp.node.right, operation: visitRightSubtree})
				search.Push(&traversalOperation{node: traversalOp.node.right, operation: visitNode})
				search.Push(&traversalOperation{node: traversalOp.node.right, operation: visitLeftSubtree})
			}
		}
	}

	return traversalResult
}

//	PostorderTraversal perform a postorder traversal of the tree and return the nodes in a slice
func (t *BinaryTreeRoot) PostorderTraversal() []interface{} {
	traversalResult := make([]interface{}, 0)
	search := stack.New()

	//	postorder means, visit the left subtree, then root, and then the right subtree
	search.Push(&traversalOperation{node: t.root, operation: visitNode})
	search.Push(&traversalOperation{node: t.root, operation: visitRightSubtree})
	search.Push(&traversalOperation{node: t.root, operation: visitLeftSubtree})

	for {
		if search.IsEmpty() {
			break
		}

		traversalOp := search.Pop().(*traversalOperation)

		switch traversalOp.operation {
		case visitNode:
			traversalResult = append(traversalResult, traversalOp.node.value)

		case visitLeftSubtree:
			if traversalOp.node.left != nil {
				search.Push(&traversalOperation{node: traversalOp.node.left, operation: visitNode})
				search.Push(&traversalOperation{node: traversalOp.node.left, operation: visitRightSubtree})
				search.Push(&traversalOperation{node: traversalOp.node.left, operation: visitLeftSubtree})
			}

		case visitRightSubtree:
			if traversalOp.node.right != nil {
				search.Push(&traversalOperation{node: traversalOp.node.right, operation: visitNode})
				search.Push(&traversalOperation{node: traversalOp.node.right, operation: visitRightSubtree})
				search.Push(&traversalOperation{node: traversalOp.node.right, operation: visitLeftSubtree})
			}
		}
	}

	return traversalResult
	return nil
}

//	LevelorderTraversal perform a level-order traversal of the tree and return the nodes in a slice
func (t *BinaryTreeRoot) LevelorderTraversal() []interface{} {
	traversalResult := make([]interface{}, 0)
	search := queue.New()

	//	level-order means, visit the nodes from top to bottom, and from left to right
	search.Put(&traversalOperation{node: t.root, operation: visitNode})
	search.Put(&traversalOperation{node: t.root, operation: visitLeftSubtree})
	search.Put(&traversalOperation{node: t.root, operation: visitRightSubtree})

	for {
		if search.IsEmpty() {
			break
		}

		traversalOp := search.Get().(*traversalOperation)

		switch traversalOp.operation {
		case visitNode:
			traversalResult = append(traversalResult, traversalOp.node.value)

		case visitLeftSubtree:
			if traversalOp.node.left != nil {
				search.Put(&traversalOperation{node: traversalOp.node.left, operation: visitNode})
				search.Put(&traversalOperation{node: traversalOp.node.left, operation: visitLeftSubtree})
				search.Put(&traversalOperation{node: traversalOp.node.left, operation: visitRightSubtree})
			}

		case visitRightSubtree:
			if traversalOp.node.right != nil {
				search.Put(&traversalOperation{node: traversalOp.node.right, operation: visitNode})
				search.Put(&traversalOperation{node: traversalOp.node.right, operation: visitLeftSubtree})
				search.Put(&traversalOperation{node: traversalOp.node.right, operation: visitRightSubtree})
			}
		}
	}

	return traversalResult
	return nil
}
