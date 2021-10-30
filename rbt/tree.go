package rbt

import (
	"github.com/emirpasic/gods/utils"
)

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

// DepthFirstTraversal (pre-order traversal) traverses the binary search tree by printing the root node,
// then recursively visiting the left and the right nodes of the current node.
func (tree *RBT) DepthFirstTraversal() {
	tree.Root().dfs()
}

// InOrderTraversal prints every node's value in order from smallest to greatest.
func (tree *RBT) InOrderTraversal() {
	tree.Root().inOrder()
}

// Insert takes a key and a value of type interface, and inserts a new Node with that key and value.
// The function inserts by key; that is, the key of the new node is
// compared against current nodes to find the correct insertion point.
// The function returns the newly inserted node's key or an error, if there was one.
func (tree *RBT) Insert(key, value interface{}) (interface{}, error) {
	newNode := NewNode(key, value, -1)
	// key already exists in the tree
	if tree.Search(key) {
		return nil, NewDuplicateError(key)
	}

	var parent *Node        // this will eventually be set as the newNode's parent
	tempNode := tree.Root() // to determine when we've hit a leaf
	for tempNode != nil {
		parent = tempNode
		switch {
		case tree.comparator(newNode.key(), parent.key()) < 0:
			tempNode = tempNode.leftChild()
		case tree.comparator(newNode.key(), parent.key()) > 0:
			tempNode = tempNode.rightChild()
		}
	}

	newNode.setParent(parent)
	if parent == nil {
		tree.setRoot(newNode)
	} else if tree.comparator(newNode.key(), parent.key()) < 0 {
		parent.setLeftChild(newNode)
	} else {
		parent.setRightChild(newNode)
	}
	// switch {
	// case parent == nil:
	// 	tree.setRoot(newNode)
	// case tree.comparator(newNode.key(), parent.key()) < 0:
	// 	parent.setLeftChild(newNode)
	// case tree.comparator(newNode.key(), parent.key()) > 0:
	// 	parent.setRightChild(newNode)
	// }
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
		if node.getParent() == node.grandparent().leftChild() {
			uncle := node.grandparent().rightChild()
			if uncle.getColor() == RED { // case 1
				node.getParent().setColor(BLACK)
				uncle.setColor(BLACK)
				node.grandparent().setColor(RED)
				node = node.grandparent()
			} else if node == node.getParent().rightChild() { // case 2
				node = node.getParent()
				tree.leftRotate(node)
			} else { // case 3
				node.getParent().setColor(BLACK)
				node.grandparent().setColor(RED)
				tree.rightRotate(node.grandparent())
			}
		} else { // node's parent is a right child, so uncle is on the left side
			uncle := node.grandparent().leftChild()
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
				tree.leftRotate(node.grandparent())
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
	toDeleteChild := new(Node)
	toDeleteCopy := nodeToDelete
	originalColor := toDeleteCopy.getColor()

	if nodeToDelete.leftChild() == nil {
		toDeleteChild = nodeToDelete.rightChild()
		tree.replaceSubTree(nodeToDelete, nodeToDelete.rightChild())
	} else if nodeToDelete.rightChild() == nil {
		toDeleteChild = nodeToDelete.leftChild()
		tree.replaceSubTree(nodeToDelete, nodeToDelete.leftChild())
	} else {
		toDeleteCopy = nodeToDelete.subtreeMin(nodeToDelete.rightChild())
		originalColor = toDeleteCopy.getColor()
		toDeleteChild = toDeleteCopy.rightChild()
		if toDeleteCopy.getParent() == nodeToDelete {
			toDeleteChild.setParent(toDeleteCopy)
		} else {
			tree.replaceSubTree(toDeleteCopy, toDeleteCopy.rightChild())
			toDeleteCopy.setRightChild(nodeToDelete.rightChild())
			toDeleteCopy.rightChild().setParent(toDeleteCopy)
		}
		tree.replaceSubTree(nodeToDelete, toDeleteCopy)
		toDeleteCopy.setLeftChild(nodeToDelete.leftChild())
		toDeleteCopy.leftChild().setParent(toDeleteCopy)
		toDeleteCopy.setColor(nodeToDelete.getColor())
	}

	if originalColor == BLACK {
		tree.deleteFixup(toDeleteChild)
	}
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

// IsBalanced returns a bool representing whether
// all paths from a node to its nil descendants contain
// the same number of black nodes.
func (tree *RBT) IsBalanced() bool {
	switch {
	case tree.IsEmpty():
		return true
	case tree.BlackHeight() < 0:
		return false
	default:
		return true
	}
}

// BlackHeight returns an int representing the black height of the tree.
func (tree *RBT) BlackHeight() int {
	if tree.IsEmpty() {
		return 0
	}

	return tree.Root().blackHeight()
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

// leftRotate performs right rotations on the nodes
// in the tree to keep the RBT height invariant.
// From CLRS: When we do a left rotation on a node x,
// we assume that its right child y is not nil; x may be any node in
// the tree whose right child is not nil.
// The left rotation “pivots” around the link from x to y.
// It makes y the new root of the subtree, with x as y’s left child and y’s
// left child as x’s right child.
func (tree *RBT) leftRotate(node *Node) {
	newParent := node.rightChild()
	node.setRightChild(newParent.leftChild())
	if newParent.leftChild() != nil {
		newParent.leftChild().setParent(node)
	}
	newParent.setParent(node.getParent())
	if node.getParent() == nil {
		tree.setRoot(newParent)
	}
	switch {
	case node == node.getParent().leftChild():
		node.getParent().setLeftChild(newParent)
	default:
		node.getParent().setRightChild(newParent)
	}
	newParent.setLeftChild(node)
	node.setParent(newParent)
}

// rightRotate performs right rotations on the nodes
// in the tree to keep the RBT height invariant.
// From CLRS: When we do a right rotation on a node x,
// we assume that its left child y is not nil; x may be any node in
// the tree whose left child is not nil.
// The right rotation “pivots” around the link from x to y.
// It makes y the new root of the subtree, with x as y’s right child and y’s
// right child as x’s left child.
func (tree *RBT) rightRotate(node *Node) {
	newParent := node.leftChild()
	node.setLeftChild(newParent.rightChild())
	if newParent.rightChild() != nil {
		newParent.rightChild().setParent(node)
	}
	newParent.setParent(node.getParent())
	if node.getParent() == nil {
		tree.setRoot(newParent)
	}
	switch {
	case node == node.getParent().rightChild():
		node.getParent().setRightChild(newParent)
	default:
		node.getParent().setLeftChild(newParent)
	}
	newParent.setRightChild(node)
	node.setParent(newParent)
}

// replaceSubTree replaces one subtree as a child of its parent with
// another subtree. When replaceSubTree replaces the subtree rooted at node u with
// the subtree rooted at node v, node u’s parent becomes node v’s parent, and u’s
// parent ends up having as its appropriate child.
func (tree *RBT) replaceSubTree(toDelete *Node, replacement *Node) {
	//parent := toDelete.getParent()
	switch {
	case toDelete.getParent() == nil:
		tree.setRoot(replacement)
	case toDelete == toDelete.getParent().leftChild(): // node to delete is left child
		toDelete.getParent().setLeftChild(replacement)
	default: // node to delete is right child
		toDelete.getParent().setRightChild(replacement)
	}
	replacement.setParent(toDelete.getParent())
	// toDelete.clear()
}

// deleteFixup maintains the invariants of the red-black tree after deletion.
func (tree *RBT) deleteFixup(child *Node) {
	for child != tree.Root() && child.getColor() == BLACK {
		if child == child.getParent().leftChild() {
			childSibling := child.getParent().rightChild()
			if childSibling.getColor() == RED {
				childSibling.recolor()
				child.getParent().setColor(RED)
				tree.leftRotate(child.getParent())
				childSibling = child.getParent().rightChild()
			}
			if childSibling.leftChild().getColor() == BLACK && childSibling.rightChild().getColor() == BLACK {
				childSibling.setColor(RED)
				child = child.getParent()
			} else if childSibling.rightChild().getColor() == BLACK {
				childSibling.leftChild().setColor(BLACK)
				childSibling.setColor(RED)
				tree.rightRotate(childSibling)
				childSibling = child.getParent().rightChild()
			}
			childSibling.setColor(child.getParent().getColor())
			child.getParent().setColor(BLACK)
			childSibling.rightChild().setColor(BLACK)
			tree.leftRotate(child.getParent())
			child = tree.Root()
		} else {
			childSibling := child.getParent().leftChild()
			if childSibling.getColor() == RED {
				childSibling.recolor()
				child.getParent().setColor(RED)
				tree.rightRotate(child.getParent())
				childSibling = child.getParent().leftChild()
			}
			if childSibling.rightChild().getColor() == BLACK && childSibling.leftChild().getColor() == BLACK {
				childSibling.setColor(RED)
				child = child.getParent()
			} else if childSibling.leftChild().getColor() == BLACK {
				childSibling.rightChild().setColor(BLACK)
				childSibling.setColor(RED)
				tree.leftRotate(childSibling)
				childSibling = child.getParent().leftChild()
			}
			childSibling.setColor(child.getParent().getColor())
			child.getParent().setColor(BLACK)
			childSibling.leftChild().setColor(BLACK)
			tree.rightRotate(child.getParent())
			child = tree.Root()
		}
		child.setColor(BLACK)
	}
}

// findNode takes a key and returns the node associated with that key.
// Returns nil and an error if no node exists.
func (tree *RBT) findNode(key interface{}) (*Node, error) {
	// if tree.IsEmpty() {
	// 	return nil, NewNilNodeError(key)
	// }
	tempNode := tree.Root()
	for tempNode != nil {
		switch {
		case tree.comparator(key, tempNode.key()) < 0: // left
			tempNode = tempNode.leftChild()
		case tree.comparator(key, tempNode.key()) > 0: // right
			tempNode = tempNode.rightChild()
		case tree.comparator(key, tempNode.key()) == 0: // match
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
