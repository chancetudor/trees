package min

// Node contains the key and the value the caller wishes to store.
type Node struct {
	Key   interface{}
	Value interface{}
}

// NewNode takes in a key and a value and returns a pointer to type Node.
// When creating a new node, the left and right children, as well as the parent node, are set to nil.
func NewNode(k, v interface{}) *Node {
	return &Node{
		Key:   k,
		Value: v,
	}
}

// setKey takes a key and sets it as the key for a node.
func (node *Node) setKey(key interface{}) {
	node.Key = key
}

// key returns the value of a node.
func (node *Node) key() interface{} {
	return node.Key
}

// setValue takes a value and sets it as the value for a node.
func (node *Node) setValue(value interface{}) {
	node.Value = value
}

// value returns the value of a node.
func (node *Node) value() interface{} {
	return node.Value
}
