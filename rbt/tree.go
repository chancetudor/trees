package rbt

import "github.com/emirpasic/gods/utils"

/* Package rbt implements a red-black tree in Go
* A red-black tree is a Node-based, balanced binary tree data structure which has the following properties:
* A node is either red or black.
* The root and leaves (nil) are black.
* If a node is red, then its children are black.
* All paths from a node to its nil descendants contain the same number of black nodes.
 */

// RBT stores the root Node of the tree, a key comparator, and the size of the tree.
// Duplicates are not allowed.
// Comparator format is taken from https://github.com/emirpasic/gods#comparator.
// Either import the package https://github.com/emirpasic/gods/utils and pass a comparator from the library,
// or write a custom comparator using guidelines from the gods README.
type RBT struct {
	root       *Node            // the root Node
	comparator utils.Comparator // the key comparator
	size       int              // number of nodes in the tree
}

// NewWith returns a pointer to a RBT where root is nil, size is 0,
// and the key comparator is set to the parameter passed in.
// The comparator format is taken from https://github.com/emirpasic/gods#comparator.
// Either import the package https://github.com/emirpasic/gods/utils and pass a comparator from the library,
// or write a custom comparator using guidelines from the gods README.
func NewWith(comparator utils.Comparator) *RBT {
	return &RBT{
		root:       nil,
		comparator: comparator,
		size:       0,
	}
}

// NewWithIntComparator returns a pointer to a RBT where root is nil, size is 0,
// and the key comparator is set to the IntComparator from package https://github.com/emirpasic/gods/utils.
// the comparator format is taken from https://github.com/emirpasic/gods#comparator.
func NewWithIntComparator() *RBT {
	return NewWith(utils.IntComparator)
}

// NewWithStringComparator returns a pointer to a RBT where root is nil, size is 0,
// and the key comparator is set to the StringComparator from package https://github.com/emirpasic/gods/utils.
// the comparator format is taken from https://github.com/emirpasic/gods#comparator.
func NewWithStringComparator() *RBT {
	return NewWith(utils.StringComparator)
}

// Insert takes a key and a value of type interface, and inserts a new Node with that key and value.
// The function inserts by key; that is, the key of the new node is
// compared against current nodes to find the correct insertion point.
// The function returns the newly inserted node's key or an error, if there was one.
func (tree *RBT) Insert(key, value interface{}) (interface{}, error) {
	newNode := NewNode(key, value, -1)

	// tree is empty, so we set the new node as the root and increase the size of the tree by 1.
	if tree.IsEmpty() {
		newNode.setColor(BLACK) // Case 0: recolor the root
		tree.setRoot(newNode)
		tree.setSize(tree.Size() + 1)
		return newNode.key(), nil
	}

	// key already exists in the tree
	if tree.Search(key) {
		return nil, NewDuplicateError(key)
	}

	parent := new(Node)     // this will eventually be set as the newNode's parent
	tempNode := tree.Root() // to determine when we've hit a leaf
	for tempNode != nil {
		parent = tempNode
		compare := tree.comparator(newNode.key(), parent.key())
		switch {
		case compare < 0:
			tempNode = tempNode.leftChild()
		case compare > 0:
			tempNode = tempNode.rightChild()
		}
	}

	newNode.setParent(parent)
	compare := tree.comparator(newNode.key(), parent.key())
	switch {
	case compare < 0:
		parent.setLeftChild(newNode)
	case compare > 0:
		parent.setRightChild(newNode)
	}

	newNode.setRightChild(nil)
	newNode.setLeftChild(nil)
	newNode.setColor(RED)
	tree.insertFixup(newNode)
	tree.setSize(tree.Size() + 1)

	return newNode.key(), nil
}

// insertFixup performs rotations and recolorations after insertion.
// Cases:
// 1. newNode's uncle is red: recolor node's parent, grandparent, and uncle.
// 2. newNode's uncle is black (triangle): rotate node's parent in
// the opposite direction of newNode's placement.
// 3. newNode's uncle is black (line): rotate node's grandparent in
// the opposite direction of newNode's placement, then recolor original parent and grandparent.
func (tree *RBT) insertFixup(node *Node) {
	for node.getParent().getColor() == RED {
		uncle, side := node.uncle()
		if side == RIGHT {
			if uncle.getColor() == RED { // case 1
				node.getParent().setColor(BLACK)
				uncle.recolor()
				node.grandparent().setColor(RED)
				node = node.grandparent()
			} else if node == node.getParent().rightChild() { // case 2
				node = node.getParent()
				tree.leftRotate(node)
			} else { // case 3
				node.getParent().setColor(BLACK)
				node.grandparent().setColor(RED)
				tree.rightRotate(node)
			}
		} else if side == LEFT {
			if uncle.getColor() == RED { // case 1
				node.getParent().setColor(BLACK)
				uncle.recolor()
				node.grandparent().setColor(RED)
				node = node.grandparent()
			} else if node == node.getParent().leftChild() { // case 2
				node = node.getParent()
				tree.rightRotate(node)
			} else { // case 3
				node.getParent().setColor(BLACK)
				node.grandparent().setColor(RED)
				tree.leftRotate(node)
			}
		}
	}
	tree.Root().setColor(BLACK)
}

// Delete takes a key, removes the node from the tree, and decrements the size of the tree.
// The function returns the key of the deleted node and an error, if there was one.
func (tree *RBT) Delete(key interface{}) (interface{}, error) {
	nodeToDelete, err := tree.findNode(key)
	// node with key does not exist
	if err != nil {
		return nil, err
	}
	nodeToDeleteKey := nodeToDelete.key()

	tree.setSize(tree.Size() - 1)
	return nodeToDeleteKey, nil
}

// Search takes a key and searches for the key in the tree.
// The function returns a boolean, stating whether the key was found or not.
func (tree *RBT) Search(key interface{}) bool {
	_, err := tree.findNode(key)
	if err != nil {
		return false
	}

	return true
}

// ReturnNodeValue takes a key and returns the value associated with the key or an error, if there was one.
func (tree *RBT) ReturnNodeValue(key interface{}) (interface{}, error) {
	matchingNode, err := tree.findNode(key)
	if err != nil {
		return nil, err
	}
	return matchingNode.value(), nil
}

// Update takes a key and a value and updates a node with the existing key with the new value.
// Returns the new value of the node or an error, if there was one.
func (tree *RBT) Update(key interface{}, value interface{}) (interface{}, error) {
	matchingNode, err := tree.findNode(key)
	if err != nil {
		return nil, err
	}
	matchingNode.setValue(value)

	return matchingNode.value(), nil
}

// Clear sets the root node to nil and sets the size of the tree to 0.
func (tree *RBT) Clear() {
	tree.setRoot(nil)
	tree.setSize(0)
}

// Root returns the root of the tree, a pointer to type Node.
func (tree *RBT) Root() *Node {
	return tree.root
}

// IsEmpty returns a boolean stating whether the tree is empty or not.
func (tree *RBT) IsEmpty() bool {
	return tree.size == 0
}

// Size returns the size, or number of nodes in the tree, of the tree.
func (tree *RBT) Size() int {
	return tree.size
}

func (tree *RBT) leftRotate(node *Node) {
	newParent := node.rightChild()
	node.setRightChild(newParent.leftChild())
	if newParent.leftChild() != nil {
		newParent.leftChild().setParent(node)
	}
	newParent.setParent(node.getParent())
	if node.isRoot() {
		tree.setRoot(newParent)
	} else if node == node.getParent().leftChild() {
		node.getParent().setLeftChild(newParent)
	} else {
		node.getParent().setRightChild(newParent)
	}
	newParent.setLeftChild(node)
	node.setParent(newParent)
}

func (tree *RBT) rightRotate(node *Node) {
	newParent := node.leftChild()
	node.setLeftChild(newParent.rightChild())
	if newParent.rightChild() != nil {
		newParent.rightChild().setParent(node)
	}
	newParent.setParent(node.getParent())
	if node.isRoot() {
		tree.setRoot(newParent)
	} else if node == node.getParent().rightChild() {
		node.getParent().setRightChild(newParent)
	} else {
		node.getParent().setLeftChild(newParent)
	}
	newParent.setRightChild(node)
	node.setParent(newParent)
}

// findNode takes a key and returns the node associated with that key.
// Returns nil and an error if no node exists.
func (tree *RBT) findNode(key interface{}) (*Node, error) {
	tempNode := tree.Root()
	for tempNode != nil {
		compare := tree.comparator(key, tempNode.key())
		switch {
		case compare < 0:
			tempNode = tempNode.leftChild()
		case compare > 0:
			tempNode = tempNode.rightChild()
		case compare == 0:
			return tempNode, nil
		}
	}

	return nil, NewNilNodeError(key)
}

// setSize sets a new size, or number of nodes in the tree, for the tree.
func (tree *RBT) setSize(newSize int) {
	tree.size = newSize
}

// setRoot takes in a pointer to a Node and sets the root of the tree to be that new Node.
func (tree *RBT) setRoot(newRoot *Node) {
	tree.root = newRoot
}