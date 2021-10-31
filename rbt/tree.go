package rbt

import (
	"fmt"
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
	if !tree.IsEmpty() {
		tree.Root().dfs()
		return
	}

	fmt.Println("Empty tree: []")
}

// InOrderTraversal prints every node's value in order from smallest to greatest.
func (tree *RBT) InOrderTraversal() {
	if !tree.IsEmpty() {
		tree.Root().inOrder()
		return
	}

	fmt.Println("Empty tree: []")
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

	var sibling, successor *Node
	if nodeToDelete.leftChild() != nil && nodeToDelete.rightChild() != nil {
		successor = nodeToDelete.successor()
	} else {
		successor = nodeToDelete
	}

	if successor.leftChild() != nil {
		sibling = successor.leftChild()
	} else {
		sibling = successor.rightChild()
	}

	newParent := successor.getParent()
	if sibling != nil {
		sibling.setParent(newParent)
	}
	if successor.getParent() == nil {
		tree.setRoot(sibling)
	} else if successor == successor.getParent().leftChild() {
		successor.getParent().setLeftChild(sibling)
	} else {
		successor.getParent().setRightChild(sibling)
	}

	if successor != nodeToDelete {
		nodeToDelete.setKey(successor.key())
		nodeToDelete.setValue(successor.value())
	}

	if successor.getColor() == BLACK {
		tree.deleteFixup(sibling, newParent)
	}
	tree.setSize(tree.Size() - 1)

	return nodeToDeleteKey, nil
}

// replaceSubTree replaces one subtree as a child of its parent with
// another subtree. When replaceSubTree replaces the subtree rooted at node u with
// the subtree rooted at node v, node u’s parent becomes node v’s parent, and u’s
// parent ends up having as its appropriate child.
func (tree *RBT) replaceSubTree(toDelete *Node, replacement *Node) {
	switch {
	case toDelete == tree.Root():
		tree.setRoot(replacement)
	case toDelete == toDelete.getParent().leftChild(): // node to delete is left child
		toDelete.getParent().setLeftChild(replacement)
	default: // node to delete is right child
		toDelete.getParent().setRightChild(replacement)
	}
	replacement.setParent(toDelete.getParent())
}

// deleteFixup maintains the invariants of the red-black tree after deletion.
// func (tree *RBT) deleteFixup(x *Node) {
// 	for x != tree.Root() && x.getColor() == BLACK {
// 		switch {
// 		case x == x.getParent().leftChild():
// 			w := x.getParent().rightChild()
// 			if w.getColor() == RED {
// 				w.recolor()
// 				x.getParent().setColor(RED)
// 				tree.leftRotate(x.getParent())
// 				w = x.getParent().rightChild()
// 			}
// 			if w.leftChild().getColor() == BLACK && w.rightChild().getColor() == BLACK {
// 				w.setColor(RED)
// 				x = x.getParent()
// 			} else {
// 				if w.rightChild().getColor() == BLACK {
// 					w.leftChild().setColor(BLACK)
// 					w.setColor(RED)
// 					tree.rightRotate(w)
// 					w = x.getParent().rightChild()
// 				}
// 				w.setColor(x.getParent().getColor())
// 				x.getParent().setColor(BLACK)
// 				w.rightChild().setColor(BLACK)
// 				tree.leftRotate(x.getParent())
// 				x = tree.Root()
// 			}
// 		default:
// 			w := x.getParent().leftChild()
// 			if w.getColor() == RED {
// 				w.recolor()
// 				x.getParent().setColor(RED)
// 				tree.rightRotate(x.getParent())
// 				w = x.getParent().leftChild()
// 			}
// 			if w.rightChild().getColor() == BLACK && w.leftChild().getColor() == BLACK {
// 				w.setColor(RED)
// 				x = x.getParent()
// 			} else {
// 				if w.leftChild().getColor() == BLACK {
// 					w.rightChild().setColor(BLACK)
// 					w.setColor(RED)
// 					tree.leftRotate(w)
// 					w = x.getParent().leftChild()
// 				}
// 				w.setColor(x.getParent().getColor())
// 				x.getParent().setColor(BLACK)
// 				w.leftChild().setColor(BLACK)
// 				tree.rightRotate(x.getParent())
// 				x = tree.Root()
// 			}
// 		}
// 	}
// 	x.setColor(BLACK)
// }

// deleteFixup maintains the invariants of the red-black tree after deletion.
func (tree *RBT) deleteFixup(x, parent *Node) {
	var w *Node

	for x != tree.Root() && x.getColor() == BLACK {
		if x != nil {
			parent = x.getParent()
		}
		switch {
		case x == parent.leftChild():
			w = parent.rightChild()
			if w.getColor() == RED {
				w.setColor(BLACK)
				parent.setColor(RED)
				tree.leftRotate(parent)
				w = parent.rightChild()
			}
			if w.leftChild().getColor() == BLACK && w.rightChild().getColor() == BLACK {
				w.setColor(RED)
				x = parent
			} else {
				if w.rightChild().getColor() == BLACK {
					if w.leftChild() != nil {
						w.leftChild().setColor(BLACK)
					}
					w.setColor(RED)
					tree.rightRotate(w)
					w = parent.rightChild()
				}
				w.setColor(parent.getColor())
				parent.setColor(BLACK)
				if w.rightChild() != nil {
					w.rightChild().setColor(BLACK)
				}
				tree.leftRotate(parent)
				x = tree.Root()
			}
		case x == parent.rightChild():
			w = parent.leftChild()
			if w.getColor() == RED {
				w.setColor(BLACK)
				parent.setColor(RED)
				tree.rightRotate(parent)
				w = parent.leftChild()
			}
			if w.leftChild().getColor() == BLACK && w.rightChild().getColor() == BLACK {
				w.setColor(RED)
				x = parent
			} else {
				if w.leftChild().getColor() == BLACK {
					if w.rightChild() != nil {
						w.rightChild().setColor(BLACK)
					}
					w.setColor(RED)
					tree.leftRotate(w)
					w = parent.leftChild()
				}
				w.setColor(parent.getColor())
				parent.setColor(BLACK)
				if w.leftChild() != nil {
					w.leftChild().setColor(BLACK)
				}
				tree.rightRotate(parent)
				x = tree.Root()
			}
		}
	}
	if x != nil {
		x.setColor(BLACK)
	}
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

// findNode takes a key and returns the node associated with that key.
// Returns nil and an error if no node exists.
func (tree *RBT) findNode(key interface{}) (*Node, error) {
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
