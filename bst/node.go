package bst

import (
	"fmt"
	"github.com/emirpasic/gods/utils"
)

// Node stores left, right, and parent Node pointers;
// and NodeData, containing the key and the value the caller wishes to store.
type Node struct {
	left   *Node
	right  *Node
	parent *Node
	Data   *NodeData
}

// NodeData stores the key and the value of the Node.
type NodeData struct {
	Key   interface{}
	Value interface{}
}

// NewNode takes in a key and a value and returns a pointer to type Node.
// When creating a new node, the left and right children, as well as the parent node, are set to nil.
func NewNode(k, v interface{}) *Node {
	return &Node{
		left:   nil,
		right:  nil,
		parent: nil,
		Data: &NodeData{
			Key:   k,
			Value: v,
		},
	}
}

// dfs traverses the nodes in a depth-first search paradigm.
// The function prints by converting each node's key and value to a string.
func (node *Node) dfs() {
	if node == nil {
		return
	}
	fmt.Println("Key: " + utils.ToString(node.Data.Key))
	fmt.Println("Value: " + utils.ToString(node.Data.Value))
	node.leftChild().dfs()
	node.rightChild().dfs()
}

// inOrder traverses the nodes "in order," printing every node's value in order from smallest to greatest.
// The function prints by converting each node's key and value to a string.
func (node *Node) inOrder() {
	if node == nil {
		return
	}
	node.leftChild().inOrder()
	fmt.Println("Key: " + utils.ToString(node.Data.Key))
	fmt.Println("Value: " + utils.ToString(node.Data.Value))
	node.rightChild().inOrder()
}

// isRoot checks to see if Node's parent is nil.
// If the parent is nil, the function returns true, as the Node is the tree's root.
// Otherwise, the function returns false.
func (node *Node) isRoot() bool {
	if node.getParent() == nil {
		return true
	}

	return false
}

// isLeaf checks to see if Node is a leaf.
// If the Node's left Node and right Node are nil, the function returns true.
// Otherwise, the function returns false.
func (node *Node) isLeaf() bool {
	if node.leftChild() == nil && node.rightChild() == nil {
		return true
	}

	return false
}

// setLeftChild sets a Node's left child
func (node *Node) setLeftChild(left *Node) {
	node.left = left
}

// leftChild returns the Node's left child.
func (node *Node) leftChild() *Node {
	return node.left
}

// setRight sets a Node's right child
func (node *Node) setRightChild(right *Node) {
	node.right = right
}

// rightChild returns the Node's right child.
func (node *Node) rightChild() *Node {
	return node.right
}

// setParent sets a node's parent.
func (node *Node) setParent(parent *Node) {
	node.parent = parent
}

// getParent returns a node's parent.
func (node *Node) getParent() *Node {
	return node.parent
}

// clear marks a node's parent and children as nil, effectively severing it from the tree.
func (node *Node) clear() {
	node.setParent(nil)
	node.setLeftChild(nil)
	node.setRightChild(nil)
}

// setKey takes a key and sets it as the key for a node.
func (node *Node) setKey(key interface{}) {
	node.Data.Key = key
}

// key returns the value of a node.
func (node *Node) key() interface{} {
	return node.Data.Key
}

// setValue takes a value and sets it as the value for a node.
func (node *Node) setValue(value interface{}) {
	node.Data.Value = value
}

// value returns the value of a node.
func (node *Node) value() interface{} {
	return node.Data.Value
}

// successor returns the node with the smallest key greater than the node the method is called on
func (node *Node) successor() *Node {
	// successor is the furthest left child of the right subtree
	if node.rightChild() != nil {
		return node.subtreeMin(node.rightChild())
	}
	// otherwise, work up and to the right of the subtrees
	parent := node.getParent()
	temp := node
	for parent != nil && temp == parent.rightChild() {
		temp = parent
		parent = parent.getParent()
	}

	return parent
}

// predecessor returns the node with the largest key smaller than the node the method is called on
func (node *Node) predecessor() *Node {
	// successor is the furthest left child of the right subtree
	if node.leftChild() != nil {
		return node.subtreeMax(node.rightChild())
	}
	// otherwise, work up and to the right of the subtrees
	parent := node.getParent()
	temp := node
	for parent != nil && temp == parent.leftChild() {
		temp = parent
		parent = parent.getParent()
	}

	return parent
}

// subtreeMin returns the furthest left child of a subtree
func (node *Node) subtreeMin(child *Node) *Node {
	temp := child
	for temp.leftChild() != nil {
		temp = temp.leftChild()
	}

	return temp
}

// subtreeMax returns the furthest right child of a subtree
func (node *Node) subtreeMax(child *Node) *Node {
	temp := child
	for temp.rightChild() != nil {
		temp = temp.rightChild()
	}

	return temp
}
