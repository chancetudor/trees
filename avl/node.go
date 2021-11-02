package avl

import (
	"fmt"
	"github.com/emirpasic/gods/utils"
	"math"
)

// Node stores left, right, and parent Node pointers;
// and NodeData, containing the key and the value the caller wishes to store.
type Node struct {
	left   *Node
	right  *Node
	parent *Node
	Data   *NodeData
	height int
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
		height: 1, // new node is added at leaf initially
		Data: &NodeData{
			Key:   k,
			Value: v,
		},
	}
}

// BalanceFactor returns the difference in a node's left subtree's height and its right subtree's height.
// Function returns 0 if node is nil.
func (node *Node) BalanceFactor() int {
	if node == nil {
		return 0
	}
	return node.rightChild().getHeight() - node.leftChild().getHeight()
}

// getHeight returns the getHeight of the tree from a specific node.
func (node *Node) calculateHeight() int {
	if node == nil {
		return -1
	}

	return 1 + int(math.Max(
		float64(node.leftChild().calculateHeight()),
		float64(node.rightChild().calculateHeight())))
}

// setHeight stores the current getHeight of the node.
func (node *Node) setHeight(h int) {
	if node != nil {
		node.height = h
	}
}

// getHeight returns the getHeight stored in the node.
// if node is nil, the function returns 0.
func (node *Node) getHeight() int {
	if node != nil {
		return node.height
	}

	return 0
}

// dfs traverses the nodes in a depth-first search paradigm.
// The function prints by converting each node's key and value to a string.
func (node *Node) dfs() {
	if node == nil {
		return
	}
	if node.getParent() == nil {
		fmt.Println("ROOT")
	}
	fmt.Println("Key = " + utils.ToString(node.Data.Key) +
		" | " + "Value = " + utils.ToString(node.Data.Value))
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
	if node.getParent() == nil {
		fmt.Println("ROOT")
	}
	fmt.Println("Key = " + utils.ToString(node.Data.Key) +
		" | " + "Value = " + utils.ToString(node.Data.Value))
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
	if node != nil {
		node.left = left
	}
}

// leftChild returns the Node's left child.
func (node *Node) leftChild() *Node {
	if node != nil {
		return node.left
	}

	return nil
}

// setRight sets a Node's right child
func (node *Node) setRightChild(right *Node) {
	if node != nil {
		node.right = right
	}
}

// rightChild returns the Node's right child.
func (node *Node) rightChild() *Node {
	if node != nil {
		return node.right
	}

	return nil
}

// setParent sets a node's parent.
func (node *Node) setParent(parent *Node) {
	if node != nil {
		node.parent = parent
	}
}

// getParent returns a node's parent.
func (node *Node) getParent() *Node {
	if node != nil {
		return node.parent
	}

	return nil
}

// clear marks a node's parent and children as nil, effectively severing it from the tree.
func (node *Node) clear() {
	node.setParent(nil)
	node.setLeftChild(nil)
	node.setRightChild(nil)
}

// setKey takes a key and sets it as the key for a node.
func (node *Node) setKey(key interface{}) {
	if node != nil {
		node.Data.Key = key
	}
}

// key returns the value of a node.
func (node *Node) key() interface{} {
	if node != nil {
		return node.Data.Key
	}

	return nil
}

// setValue takes a value and sets it as the value for a node.
func (node *Node) setValue(value interface{}) {
	if node != nil {
		node.Data.Value = value
	}
}

// value returns the value of a node.
func (node *Node) value() interface{} {
	if node != nil {
		return node.Data.Value
	}

	return nil
}

// successor returns the node with the smallest key greater than the node the method is called on
func (node *Node) successor() *Node {
	// successor is the furthest left child of the right subtree
	if node.rightChild() != nil {
		return node.rightChild().subtreeMin()
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
		return node.rightChild().subtreeMax()
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
func (node *Node) subtreeMin() *Node {
	temp := node
	for temp.leftChild() != nil {
		temp = temp.leftChild()
	}

	return temp
}

// subtreeMax returns the furthest right child of a subtree
func (node *Node) subtreeMax() *Node {
	temp := node
	for temp.rightChild() != nil {
		temp = temp.rightChild()
	}

	return temp
}
