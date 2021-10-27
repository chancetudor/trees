package bst

import (
	"github.com/emirpasic/gods/utils"
)

/* Package bst implements a binary search tree in Go
* A Binary Search Tree is a Node-based binary tree data structure which has the following properties:
* The left subtree of a Node contains only nodes with keys lesser than the Node’s key.
* The right subtree of a Node contains only nodes with keys greater than the Node’s key.
* The left and right subtree each must also be a binary search tree.
 */

// BST stores the root Node of the tree, a key comparator, and the size of the tree.
// Duplicates are not allowed.
// Comparator format is taken from https://github.com/emirpasic/gods#comparator.
// Either import the package https://github.com/emirpasic/gods/utils and pass a comparator from the library,
// or write a custom comparator using guidelines from the gods README.
type BST struct {
	root       *Node            // the root Node
	comparator utils.Comparator // the key comparator
	size       int              // number of nodes in the tree
}

// NewWith returns a pointer to a BST where root is nil, size is 0,
// and the key comparator is set to the parameter passed in.
// The comparator format is taken from https://github.com/emirpasic/gods#comparator.
// Either import the package https://github.com/emirpasic/gods/utils and pass a comparator from the library,
// or write a custom comparator using guidelines from the gods README.
func NewWith(comparator utils.Comparator) *BST {
	return &BST{
		root:       nil,
		comparator: comparator,
		size:       0,
	}
}

// NewWithIntComparator returns a pointer to a BST where root is nil, size is 0,
// and the key comparator is set to the IntComparator from package https://github.com/emirpasic/gods/utils.
// the comparator format is taken from https://github.com/emirpasic/gods#comparator.
func NewWithIntComparator() *BST {
	return NewWith(utils.IntComparator)
}

// NewWithStringComparator returns a pointer to a BST where root is nil, size is 0,
// and the key comparator is set to the StringComparator from package https://github.com/emirpasic/gods/utils.
// the comparator format is taken from https://github.com/emirpasic/gods#comparator.
func NewWithStringComparator() *BST {
	return NewWith(utils.StringComparator)
}

// DepthFirstTraversal (pre-order traversal) traverses the binary search tree by printing the root node,
// then recursively visiting the left and the right nodes of the current node.
func (tree *BST) DepthFirstTraversal() {
	tree.Root().dfs()
}

// InOrderTraversal prints every node's value in order from smallest to greatest.
func (tree *BST) InOrderTraversal() {
	tree.Root().inOrder()
}

// Insert takes a key and a value of type interface, and inserts a new Node with that key and value.
// The function inserts by key; that is, the key of the new node is
// compared against current nodes to find the correct insertion point.
// The function returns the newly inserted node's key or an error, if there was one.
func (tree *BST) Insert(key, value interface{}) (interface{}, error) {
	newNode := NewNode(key, value)

	// tree is empty, so we set the new node as the root and increase the size of the tree by 1.
	if tree.IsEmpty() {
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

	tree.setSize(tree.Size() + 1)

	return newNode.key(), nil
}

// Search takes a key and searches for the key in the tree.
// The function returns a boolean, stating whether the key was found or not.
func (tree *BST) Search(key interface{}) bool {
	_, err := tree.findNode(key)
	if err != nil {
		return false
	}

	return true
}

// findNode takes a key and returns the node associated with that key.
// Returns nil and an error if no node exists.
func (tree *BST) findNode(key interface{}) (*Node, error) {
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

// Update takes a key and a value and updates a node with the existing key with the new value.
// Returns the new value of the node or an error, if there was one.
func (tree *BST) Update(key interface{}, value interface{}) (interface{}, error) {
	matchingNode, err := tree.findNode(key)
	if err != nil {
		return nil, err
	}
	matchingNode.setValue(value)

	return matchingNode.value(), nil
}

// ReturnNodeValue takes a key and returns the value associated with the key or an error, if there was one.
func (tree *BST) ReturnNodeValue(key interface{}) (interface{}, error) {
	matchingNode, err := tree.findNode(key)
	if err != nil {
		return nil, err
	}
	return matchingNode.value(), nil
}

// Delete takes a key, removes the node from the tree, and decrements the size of the tree.
// The function returns the key of the deleted node and an error, if there was one.
func (tree *BST) Delete(key interface{}) (interface{}, error) {
	nodeToDelete, err := tree.findNode(key)
	// node with key does not exist
	if err != nil {
		return nil, err
	}
	nodeToDeleteKey := nodeToDelete.key()

	// node is already a leaf
	if nodeToDelete.isLeaf() {
		tree.pruneLeaf(nodeToDelete)
		return nodeToDeleteKey, nil
	}

	switch {
	case nodeToDelete.leftChild() == nil: // the node to delete only has a right subtree
		tree.replaceSubTree(nodeToDelete, nodeToDelete.rightChild())
	case nodeToDelete.rightChild() == nil: // the node to delete only has a left subtree
		tree.replaceSubTree(nodeToDelete, nodeToDelete.leftChild())
	default: // the node to delete has two subtrees
		successor := nodeToDelete.successor()
		if successor.getParent() != nodeToDelete {
			tree.replaceSubTree(successor, successor.rightChild())
			successor.setRightChild(nodeToDelete.rightChild())
			successor.rightChild().setParent(successor)
		}

		tree.replaceSubTree(nodeToDelete, successor)
		successor.setLeftChild(nodeToDelete.leftChild())
		successor.leftChild().setParent(successor)
	}

	tree.setSize(tree.Size() - 1)
	return nodeToDeleteKey, nil
}

// replaceSubTree replaces the node to delete with a new root node of a subtree.
// That is, if a node to delete is the root node of a subtree,
// the function replaces it with a new root of that subtree.
func (tree *BST) replaceSubTree(toDelete *Node, replacementNode *Node) {
	parent := toDelete.getParent()
	switch {
	case toDelete.isRoot():
		tree.setRoot(replacementNode)
	// case tree.comparator(toDelete.key(), toDelete.getParent().leftChild().key()) == 0: // node to delete is left child
	// 	toDelete.getParent().setLeftChild(child)
	case toDelete == parent.leftChild(): // node to delete is left child
		parent.setLeftChild(replacementNode)
	default: // node to delete is right child
		parent.setRightChild(replacementNode)
	}
	if replacementNode != nil {
		replacementNode.setParent(toDelete.getParent())
	}
	// toDelete.clear()
}

// pruneLeaf removes a leaf from the tree.
// If the node to delete is the root, tree.Clear() is called.
// The function does not decrement the size of the tree.
func (tree *BST) pruneLeaf(toDelete *Node) {
	if toDelete.isRoot() {
		tree.Clear()
		return
	}
	parent := toDelete.getParent()
	// compare := tree.comparator(parent.key(), toDelete.key())
	switch {
	case toDelete == parent.leftChild(): // node to delete is left of parent
		parent.setLeftChild(nil)
		toDelete.clear()
	case toDelete == parent.rightChild(): // node to delete is right of parent
		parent.setRightChild(nil)
		toDelete.clear()
		// case compare < 0: // node to delete is left of parent
		// 	parent.setLeftChild(nil)
		// 	toDelete.setParent(nil)
		// case compare > 0:
		// 	parent.setRightChild(nil)
		// 	toDelete.setParent(nil)
	}
}

// Clear sets the root node to nil and sets the size of the tree to 0.
func (tree *BST) Clear() {
	tree.setRoot(nil)
	tree.setSize(0)
}

// Root returns the root of the tree, a pointer to type Node.
func (tree *BST) Root() *Node {
	return tree.root
}

// setRoot takes in a pointer to a Node and sets the root of the tree to be that new Node.
func (tree *BST) setRoot(newRoot *Node) {
	tree.root = newRoot
}

// Size returns the size, or number of nodes in the tree, of the tree.
func (tree *BST) Size() int {
	return tree.size
}

// setSize sets a new size, or number of nodes in the tree, for the tree.
func (tree *BST) setSize(newSize int) {
	tree.size = newSize
}

// IsEmpty returns a boolean stating whether the tree is empty or not.
func (tree *BST) IsEmpty() bool {
	return tree.size == 0
}
