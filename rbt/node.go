package rbt

// Node stores left, right, and parent Node pointers; the node's color,
// and NodeData, containing the key and the value the caller wishes to store.
type Node struct {
	left   *Node
	right  *Node
	parent *Node
	Data   *NodeData
	Color  int
}

// NodeData stores the key and the value of the Node.
type NodeData struct {
	Key   interface{}
	Value interface{}
}

// NewNode takes in a key and a value and returns a pointer to type Node.
// When creating a new node, the left and right children, as well as the parent node, are set to nil.
func NewNode(color int, k, v interface{}) *Node {
	return &Node{
		left:   nil,
		right:  nil,
		parent: nil,
		Data: &NodeData{
			Key:   k,
			Value: v,
		},
		Color: color,
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
