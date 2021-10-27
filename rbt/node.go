package rbt

const BLACK = 0
const RED = 1
const LEFT = 2
const RIGHT = 3

// Node stores left, right, and parent Node pointers; the node's color,
// and NodeData, containing the key and the value the caller wishes to store.
type Node struct {
	left   *Node
	right  *Node
	parent *Node
	Data   *NodeData
	color  int
}

// NodeData stores the key and the value of the Node.
type NodeData struct {
	Key   interface{}
	Value interface{}
}

// NewNode takes in a key and a value and returns a pointer to type Node.
// When creating a new node, the left and right children, as well as the parent node, are set to nil.
func NewNode(k, v interface{}, color int) *Node {
	return &Node{
		left:   nil,
		right:  nil,
		parent: nil,
		Data: &NodeData{
			Key:   k,
			Value: v,
		},
		color: color,
	}
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

// setGrandparent sets a node's grandparent.
func (node *Node) setGrandparent(grandparent *Node) {
	node.parent.parent = grandparent
}

// grandparent returns a node's grandparent.
func (node *Node) grandparent() *Node {
	if node != nil {
		return node.parent.parent
	}

	return nil
}

// uncle returns a node's uncle and a side flag, for use in insertFixup.
func (node *Node) uncle() (*Node, int) {
	// return grandparent's right child
	if node.grandparent().leftChild() == node.getParent() {
		return node.grandparent().rightChild(), RIGHT
	}
	// else return grandparent's left child
	return node.grandparent().leftChild(), LEFT
}

// recolor flips the color of the node.
// i.e., if a node is black, recolor colors the node red
// and if a node is red, recolor colors the node black.
func (node *Node) recolor() {
	if node.color == BLACK {
		node.setColor(RED)
		return
	}

	node.setColor(BLACK)
}

// setColor sets a node's color to either red or black.
func (node *Node) setColor(newColor int) {
	if node != nil {
		node.color = newColor
	}
}

// getColor returns a node's color, either red (1) or black (0).
// Returns black if the node is nil.
func (node *Node) getColor() int {
	if node != nil {
		return node.color
	}

	return BLACK
}

// clear marks a node's parent and children as nil, effectively severing it from the tree.
func (node *Node) clear() {
	node.setParent(nil)
	node.setLeftChild(nil)
	node.setRightChild(nil)
	node.setColor(BLACK)
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
