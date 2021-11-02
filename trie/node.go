package trie

type Node struct {
	children map[interface{}]*Node
	terminal bool
}

func NewNode() *Node {
	return &Node{
		children: make(map[interface{}]*Node),
		terminal: false,
	}
}
