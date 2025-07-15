package main

import "fmt"

type AVLNode struct {
	value  *User
	left   *AVLNode
	right  *AVLNode
	height int
}

func newAVLNode(user *User) *AVLNode {
	return &AVLNode{
		value:  user,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

type AVL struct {
	root *AVLNode
}

func NewAVL() *AVL {
	return &AVL{
		root: nil,
	}
}

func (avl *AVL) printNode(node *AVLNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	rightPrefix := prefix
	if isLeft {
		rightPrefix += "│   "
	} else {
		rightPrefix += "    "
	}
	avl.printNode(node.right, rightPrefix, false)

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
	avl.printNode(node.left, leftPrefix, true)
}

func (avl *AVL) PrintTree() {
	fmt.Println("------------ AVL TREE ------------")
	if avl.root == nil || avl.root.value == nil {
		fmt.Println("Tree is empty.")
		fmt.Println("----------------------------------------------")
		return
	}
	avl.printNode(avl.root.right, "", false)

	fmt.Println(avl.root.value.UserId)

	avl.printNode(avl.root.left, "", true)
	fmt.Println("----------------------------------")
}

func (avl *AVL) getMinValueNode(node *AVLNode) *AVLNode {
	curr := node
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (avl *AVL) height(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (avl *AVL) updateHeight(node *AVLNode) {
	if node != nil {
		node.height = 1 + max(avl.height(node.left), avl.height(node.right))
	}
}

func (avl *AVL) balanceFactor(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return avl.height(node.left) - avl.height(node.right)
}

func (avl *AVL) leftRotate(node *AVLNode) *AVLNode {
	rightChild := node.right

	node.right = rightChild.left
	rightChild.left = node

	avl.updateHeight(node)
	avl.updateHeight(node.right)

	return rightChild
}

func (avl *AVL) rightRotate(node *AVLNode) *AVLNode {
	leftChild := node.left

	node.left = leftChild.right
	leftChild.right = node

	avl.updateHeight(node)
	avl.updateHeight(node.left)

	return leftChild
}

func (avl *AVL) Insert(node *AVLNode, user *User) *AVLNode {
	if node == nil {
		return newAVLNode(user)
	}

	if user.UserId < node.value.UserId {
		node.left = avl.Insert(node.left, user)
	} else if user.UserId > node.value.UserId {
		node.right = avl.Insert(node.right, user)
	} else {
		return node
	}

	avl.updateHeight(node)

	balance := avl.balanceFactor(node)

	// Left Left Case
	if balance > 1 && user.UserId < node.left.value.UserId {
		return avl.rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && user.UserId > node.right.value.UserId {
		return avl.leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && user.UserId > node.left.value.UserId {
		node.left = avl.leftRotate(node.left)
		return avl.rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && user.UserId < node.right.value.UserId {
		node.right = avl.rightRotate(node.right)
		return avl.leftRotate(node)
	}

	return node
}

func (avl *AVL) Remove(node *AVLNode, key int) *AVLNode {

	if node == nil {
		return node
	}

	if key < node.value.UserId {
		node.left = avl.Remove(node.left, key)
	} else if key > node.value.UserId {
		node.right = avl.Remove(node.right, key)
	} else {
		if (node.left == nil) || (node.right == nil) {
			var temp *AVLNode
			if node.left != nil {
				temp = node.left
			} else {
				temp = node.right
			}
			if temp == nil {
				temp = node
				node = nil
			} else {
				*node = *temp
			}
		} else {
			var temp *AVLNode = avl.getMinValueNode(node.right)
			node.value.UserId = temp.value.UserId
			node.right = avl.Remove(node.right, temp.value.UserId)
		}
	}

	if node == nil {
		return node
	}

	avl.updateHeight(node)

	balance := avl.balanceFactor(node)

	// Left Left Case
	if balance > 1 && avl.balanceFactor(node.left) >= 0 {
		return avl.rightRotate(node)
	}

	// Left Right Case
	if balance > 1 && avl.balanceFactor(node.left) < 0 {
		node.left = avl.leftRotate(node.left)
		return avl.rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && avl.balanceFactor(node.right) <= 0 {
		return avl.leftRotate(node)
	}

	// Right Left Case
	if balance < -1 && avl.balanceFactor(node.right) > 0 {
		node.right = avl.rightRotate(node.right)
		return avl.leftRotate(node)
	}

	return node
}
