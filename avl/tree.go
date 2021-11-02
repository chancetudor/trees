package avl

import (
	"fmt"
	"github.com/emirpasic/gods/utils"
)

/* Package avl implements an AVL tree in Go
* An AVL tree is a Node-based binary tree data structure which has the following properties:
* In an AVL tree, the heights of the two child subtrees of any node differ by at most one;
* if at any time they differ by more than one, rebalancing is done to restore this property.
* Lookup, insertion, and deletion all take O(log n) time in both the average and worst cases,
* where n is the number of nodes in the tree prior to the operation.
 */

// AVL stores the root Node of the tree, a key comparator, and the size of the tree.
// Duplicates are not allowed.
// Comparator format is taken from https://github.com/emirpasic/gods#comparator.
// Either import the package https://github.com/emirpasic/gods/utils and pass a comparator from the library,
// or write a custom comparator using guidelines from the gods README.
type AVL struct {
	root       *Node            // the root Node
	comparator utils.Comparator // the key comparator
	size       int              // number of nodes in the tree
}

// NewWith returns a pointer to a BST where root is nil, size is 0,
// and the key comparator is set to the parameter passed in.
// The comparator format is taken from https://github.com/emirpasic/gods#comparator.
// Either import the package https://github.com/emirpasic/gods/utils and pass a comparator from the library,
// or write a custom comparator using guidelines from the gods README.
func NewWith(comparator utils.Comparator) *AVL {
	return &AVL{
		root:       nil,
		comparator: comparator,
		size:       0,
	}
}

// NewWithIntComparator returns a pointer to a BST where root is nil, size is 0,
// and the key comparator is set to the IntComparator from package https://github.com/emirpasic/gods/utils.
// the comparator format is taken from https://github.com/emirpasic/gods#comparator.
func NewWithIntComparator() *AVL {
	return NewWith(utils.IntComparator)
}

// NewWithStringComparator returns a pointer to a BST where root is nil, size is 0,
// and the key comparator is set to the StringComparator from package https://github.com/emirpasic/gods/utils.
// the comparator format is taken from https://github.com/emirpasic/gods#comparator.
func NewWithStringComparator() *AVL {
	return NewWith(utils.StringComparator)
}

// Insert takes a key and a value of type interface, and inserts a new Node with that key and value.
// The function inserts by key; that is, the key of the new node is
// compared against current nodes to find the correct insertion point.
// The function returns the newly inserted node's key or an error, if there was one.
func (tree *AVL) Insert(key, value interface{}) (interface{}, error) {
	newNode := NewNode(key, value)

	// tree is empty, so we set the new node as the root and increase the size of the tree by 1.
	if tree.IsEmpty() {
		tree.setRoot(newNode)
		tree.setSize(tree.Size() + 1)
		newNode.setHeight(0)
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
	newNode.setHeight(newNode.calculateHeight())
	tree.fixup(newNode)
	tree.setSize(tree.Size() + 1)

	return newNode.key(), nil
}

// Delete takes a key, removes the node from the tree, and decrements the size of the tree.
// The function returns the key of the deleted node and an error, if there was one.
func (tree *AVL) Delete(key interface{}) (interface{}, error) {
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
	tree.fixup(nodeToDelete)
	tree.setSize(tree.Size() - 1)

	return nodeToDeleteKey, nil
}

// fixup rebalances the AVL tree to maintain the invariant:
// -1 <= getHeight(leftSubtree) - getHeight(rightSubtree) <= 1
func (tree *AVL) fixup(node *Node) {
	bf := node.BalanceFactor()
	if bf < -1 || bf > 1 {
		tree.rebalance(node)
	}
	if node.isRoot() {
		return
	}
	tree.fixup(node.getParent())
}

// replaceSubTree replaces the node to delete with a new root node of a subtree.
// That is, if a node to delete is the root node of a subtree,
// the function replaces it with a new root of that subtree.
func (tree *AVL) replaceSubTree(toDelete *Node, replacementNode *Node) {
	parent := toDelete.getParent()
	switch {
	case toDelete.isRoot():
		tree.setRoot(replacementNode)
	case toDelete == parent.leftChild(): // node to delete is left child
		parent.setLeftChild(replacementNode)
	default: // node to delete is right child
		parent.setRightChild(replacementNode)
	}
	if replacementNode != nil {
		replacementNode.setParent(toDelete.getParent())
	}
}

// pruneLeaf removes a leaf from the tree.
// If the node to delete is the root, tree.Clear() is called.
// The function does not decrement the size of the tree.
func (tree *AVL) pruneLeaf(toDelete *Node) {
	if toDelete.isRoot() {
		tree.Clear()
		return
	}
	parent := toDelete.getParent()
	switch {
	case toDelete == parent.leftChild(): // node to delete is left of parent
		parent.setLeftChild(nil)
	case toDelete == parent.rightChild(): // node to delete is right of parent
		parent.setRightChild(nil)
	}
	tree.fixup(toDelete)
	toDelete.clear()
}

// rebalance determines which rotations to perform to maintain the AVL invariant.
func (tree *AVL) rebalance(node *Node) {
	switch {
	case node.leftChild().getHeight() > node.rightChild().getHeight()+1:
		if node.leftChild().leftChild().getHeight() >= node.leftChild().rightChild().getHeight() { // left-left subtree bigger
			tree.rightRotate(node)
		} else { // right subtree bigger
			tree.leftRightRotate(node)
		}
	case node.rightChild().getHeight() > node.leftChild().getHeight()+1:
		if node.rightChild().rightChild().getHeight() >= node.rightChild().leftChild().getHeight() { // left-left subtree bigger
			tree.leftRotate(node)
		} else { // right subtree bigger
			tree.rightLeftRotate(node)
		}
	}
}

// leftRightRotate performs a left rotation on a node's left subtree, then a right rotation on the node itself.
func (tree *AVL) leftRightRotate(node *Node) {
	tree.leftRotate(node.leftChild())
	tree.rightRotate(node)
}

// rightLeftRotate performs a right rotation on a node's right subtree, then a left rotation on the node itself.
func (tree *AVL) rightLeftRotate(node *Node) {
	tree.rightRotate(node.rightChild())
	tree.leftRotate(node)
}

// leftRotate performs right rotations on the nodes
// in the tree to keep the RBT getHeight invariant.
// From CLRS: When we do a left rotation on a node x,
// we assume that its right child y is not nil; x may be any node in
// the tree whose right child is not nil.
// The left rotation “pivots” around the link from x to y.
// It makes y the new root of the subtree, with x as y’s left child and y’s
// left child as x’s right child.
func (tree *AVL) leftRotate(node *Node) {
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
	newParent.setHeight(newParent.calculateHeight())
	node.setHeight(node.calculateHeight())
}

// rightRotate performs right rotations on the nodes
// in the tree to keep the RBT getHeight invariant.
// From CLRS: When we do a right rotation on a node x,
// we assume that its left child y is not nil; x may be any node in
// the tree whose left child is not nil.
// The right rotation “pivots” around the link from x to y.
// It makes y the new root of the subtree, with x as y’s right child and y’s
// right child as x’s left child.
func (tree *AVL) rightRotate(node *Node) {
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
	newParent.setHeight(newParent.calculateHeight())
	node.setHeight(node.calculateHeight())
}

// Search takes a key and searches for the key in the tree.
// The function returns a boolean, stating whether the key was found or not.
func (tree *AVL) Search(key interface{}) bool {
	_, err := tree.findNode(key)
	if err != nil {
		return false
	}

	return true
}

// IsBalanced returns a bool representing whether
// the AVL tree maintains the invariant:
// -1 <= getHeight(leftSubtree) - getHeight(rightSubtree) <= 1
func (tree *AVL) IsBalanced() bool {
	switch {
	case tree.IsEmpty():
		return true
	case tree.Root().BalanceFactor() < -1 || tree.Root().BalanceFactor() > 1:
		return false
	default:
		return true
	}
}

// DepthFirstTraversal (pre-order traversal) traverses the binary search tree by printing the root node,
// then recursively visiting the left and the right nodes of the current node.
func (tree *AVL) DepthFirstTraversal() {
	if !tree.IsEmpty() {
		tree.Root().dfs()
		return
	}

	fmt.Println("Empty tree: []")
}

// InOrderTraversal prints every node's value in order from smallest to greatest.
func (tree *AVL) InOrderTraversal() {
	if !tree.IsEmpty() {
		tree.Root().inOrder()
		return
	}

	fmt.Println("Empty tree: []")
}

// findNode takes a key and returns the node associated with that key.
// Returns nil and an error if no node exists.
func (tree *AVL) findNode(key interface{}) (*Node, error) {
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
func (tree *AVL) Update(key interface{}, value interface{}) (interface{}, error) {
	matchingNode, err := tree.findNode(key)
	if err != nil {
		return nil, err
	}
	matchingNode.setValue(value)

	return matchingNode.value(), nil
}

// ReturnNodeValue takes a key and returns the value associated with the key or an error, if there was one.
func (tree *AVL) ReturnNodeValue(key interface{}) (interface{}, error) {
	matchingNode, err := tree.findNode(key)
	if err != nil {
		return nil, err
	}
	return matchingNode.value(), nil
}

// Clear sets the root node to nil and sets the size of the tree to 0.
func (tree *AVL) Clear() {
	tree.setRoot(nil)
	tree.setSize(0)
}

// Root returns the root of the tree, a pointer to type Node.
func (tree *AVL) Root() *Node {
	return tree.root
}

// setRoot takes in a pointer to a Node and sets the root of the tree to be that new Node.
func (tree *AVL) setRoot(newRoot *Node) {
	tree.root = newRoot
}

// Size returns the size, or number of nodes in the tree, of the tree.
func (tree *AVL) Size() int {
	return tree.size
}

// setSize sets a new size, or number of nodes in the tree, for the tree.
func (tree *AVL) setSize(newSize int) {
	tree.size = newSize
}

// IsEmpty returns a boolean stating whether the tree is empty or not.
func (tree *AVL) IsEmpty() bool {
	return tree.size == 0
}
