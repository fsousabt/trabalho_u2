package main

import "fmt"

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

type RedBlackTreeNode struct {
	value  *User
	left   *RedBlackTreeNode
	right  *RedBlackTreeNode
	color  Color
	parent *RedBlackTreeNode
}

func newRedBlackTreeNode(user *User, color Color) *RedBlackTreeNode {
	return &RedBlackTreeNode{
		color:  color,
		value:  user,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}

type RedBlackTree struct {
	root *RedBlackTreeNode
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		root: nil,
	}
}

func (rbt *RedBlackTree) printNode(node *RedBlackTreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	rightPrefix := prefix
	if isLeft {
		rightPrefix += "│   "
	} else {
		rightPrefix += "    "
	}
	rbt.printNode(node.right, rightPrefix, false)

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("└── ")
	} else {
		fmt.Print("┌── ")
	}
	fmt.Println(node.value.UserId)

	leftPrefix := prefix
	if isLeft {
		leftPrefix += "    "
	} else {
		leftPrefix += "│   "
	}
	rbt.printNode(node.left, leftPrefix, true)
}

func (rbt *RedBlackTree) PrintTree() {
	fmt.Println("------------ RED BLACK TREE ------------")
	if rbt.root == nil || rbt.root.value == nil {
		fmt.Println("Tree is empty.")
		fmt.Println("----------------------------------------------")
		return
	}
	rbt.printNode(rbt.root.right, "", false)

	fmt.Println(rbt.root.value.UserId)

	rbt.printNode(rbt.root.left, "", true)
	fmt.Println("----------------------------------")
}

func (rbt *RedBlackTree) leftRotate(node *RedBlackTreeNode) *RedBlackTreeNode {
	rightChild := node.right
	node.right = rightChild.left
	if rightChild.left != nil {
		rightChild.left.parent = node
	}
	rightChild.parent = node.parent
	if node.parent == nil {
		rbt.root = rightChild
	} else if node == node.parent.left {
		node.parent.left = rightChild
	} else {
		node.parent.right = rightChild
	}
	rightChild.left = node
	node.parent = rightChild
	return rightChild
}

func (rbt *RedBlackTree) rightRotate(node *RedBlackTreeNode) *RedBlackTreeNode {
	leftChild := node.left
	node.left = leftChild.right
	if leftChild.right != nil {
		leftChild.right.parent = node
	}
	leftChild.parent = node.parent
	if node.parent == nil {
		rbt.root = leftChild
	} else if node == node.parent.right {
		node.parent.right = leftChild
	} else {
		node.parent.left = leftChild
	}
	leftChild.right = node
	node.parent = leftChild
	return leftChild
}

func isRed(node *RedBlackTreeNode) bool {
	if node == nil {
		return false
	}
	return node.color == RED
}

func (rbt *RedBlackTree) fixInsert(node *RedBlackTreeNode) {
	parentNode := node.parent
	for parentNode != nil && isRed(parentNode) {
		grandparentNode := parentNode.parent

		if parentNode == grandparentNode.left {
			uncle := grandparentNode.right

			if uncle != nil && isRed(uncle) {
				grandparentNode.color = RED
				parentNode.color = BLACK
				uncle.color = BLACK
				node = grandparentNode
			} else {
				if node == parentNode.right {
					node = parentNode
					rbt.leftRotate(node)
				}

				parentNode.color = BLACK
				grandparentNode.color = RED
				rbt.rightRotate(grandparentNode)
			}
		} else {
			uncle := grandparentNode.left

			if uncle != nil && isRed(uncle) {
				grandparentNode.color = RED
				parentNode.color = BLACK
				uncle.color = BLACK
				node = grandparentNode
			} else {
				if node == parentNode.left {
					node = parentNode
					rbt.rightRotate(node)
				}

				parentNode.color = BLACK
				grandparentNode.color = RED
				rbt.leftRotate(grandparentNode)
			}
		}
		parentNode = node.parent
	}

	rbt.root.color = BLACK
}

func (rbt *RedBlackTree) Insert(user *User) {
	newNode := newRedBlackTreeNode(user, RED)
	if rbt.root == nil {
		newNode.color = BLACK
		rbt.root = newNode
		return
	}

	currentNode := rbt.root
	var parentNode *RedBlackTreeNode

	for currentNode != nil {
		parentNode = currentNode
		if newNode.value.UserId < currentNode.value.UserId {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	newNode.parent = parentNode
	if newNode.value.UserId < parentNode.value.UserId {
		parentNode.left = newNode
	} else {
		parentNode.right = newNode
	}

	rbt.fixInsert(newNode)
}

func (rbt *RedBlackTree) Search(key int) *RedBlackTreeNode {
	curr := rbt.root
	for curr != nil {
		if key < curr.value.UserId {
			curr = curr.left
		} else if key > curr.value.UserId {
			curr = curr.right
		} else {
			return curr
		}
	}
	return nil
}

func (rbt *RedBlackTree) getMinNode(node *RedBlackTreeNode) *RedBlackTreeNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (rbt *RedBlackTree) transplant(u, v *RedBlackTreeNode) {
	if u.parent == nil {
		rbt.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (rbt *RedBlackTree) fixDelete(node *RedBlackTreeNode, parent *RedBlackTreeNode) {
	for node != rbt.root && !isRed(node) {
		if parent == nil {
			break
		}

		if node == parent.left {
			sibling := parent.right

			if isRed(sibling) {
				sibling.color = BLACK
				parent.color = RED
				rbt.leftRotate(parent)
				sibling = parent.right
			}

			if sibling == nil || (!isRed(sibling.left) && !isRed(sibling.right)) {
				if sibling != nil {
					sibling.color = RED
				}
				node = parent
				parent = node.parent
			} else {
				if !isRed(sibling.right) {
					if sibling.left != nil {
						sibling.left.color = BLACK
					}
					sibling.color = RED
					rbt.rightRotate(sibling)
					sibling = parent.right
				}

				if sibling != nil {
					sibling.color = parent.color
				}
				parent.color = BLACK
				if sibling != nil && sibling.right != nil {
					sibling.right.color = BLACK
				}
				rbt.leftRotate(parent)
				node = rbt.root
			}
		} else {
			sibling := parent.left

			if isRed(sibling) {
				sibling.color = BLACK
				parent.color = RED
				rbt.rightRotate(parent)
				sibling = parent.left
			}

			if sibling == nil || (!isRed(sibling.left) && !isRed(sibling.right)) {
				if sibling != nil {
					sibling.color = RED
				}
				node = parent
				parent = node.parent
			} else {
				if !isRed(sibling.left) {
					if sibling.right != nil {
						sibling.right.color = BLACK
					}
					sibling.color = RED
					rbt.leftRotate(sibling)
					sibling = parent.left
				}

				if sibling != nil {
					sibling.color = parent.color
				}
				parent.color = BLACK
				if sibling != nil && sibling.left != nil {
					sibling.left.color = BLACK
				}
				rbt.rightRotate(parent)
				node = rbt.root
			}
		}
	}

	if node != nil {
		node.color = BLACK
	}
}

func (rbt *RedBlackTree) Remove(id int) {
	nodeToDelete := rbt.Search(id)
	if nodeToDelete == nil {
		return
	}

	originalColor := nodeToDelete.color
	var fixNode *RedBlackTreeNode
	var fixNodeParent *RedBlackTreeNode

	if nodeToDelete.left == nil {
		fixNode = nodeToDelete.right
		fixNodeParent = nodeToDelete.parent
		rbt.transplant(nodeToDelete, nodeToDelete.right)
	} else if nodeToDelete.right == nil {
		fixNode = nodeToDelete.left
		fixNodeParent = nodeToDelete.parent
		rbt.transplant(nodeToDelete, nodeToDelete.left)
	} else {
		successor := rbt.getMinNode(nodeToDelete.right)
		originalColor = successor.color
		fixNode = successor.right

		if successor.parent == nodeToDelete {
			if fixNode != nil {
				fixNode.parent = successor
			}
			fixNodeParent = successor
		} else {
			rbt.transplant(successor, successor.right)
			successor.right = nodeToDelete.right
			successor.right.parent = successor
			fixNodeParent = successor.parent
		}

		rbt.transplant(nodeToDelete, successor)
		successor.left = nodeToDelete.left
		successor.left.parent = successor
		successor.color = nodeToDelete.color
	}

	if originalColor == BLACK {
		rbt.fixDelete(fixNode, fixNodeParent)
	}
}
