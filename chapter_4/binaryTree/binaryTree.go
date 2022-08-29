////////////////////////////////////////////////////////////////////////////////
//	binaryTree.go  -  Ago-27-2022  -  aldebap
//
//	Binary Tree Data Structure
////////////////////////////////////////////////////////////////////////////////

package binaryTree

type BinaryTree interface {
	Add(value interface{})
	Find(value interface{}) bool
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
