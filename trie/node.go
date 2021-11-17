package trie

type Node struct {
	children map[string]*Node
	terminal bool
}

func NewNode() *Node {
	return &Node{
		children: make(map[string]*Node),
		terminal: false,
	}
}

func (n *Node) markEnd() {
	n.terminal = true
}

func (n *Node) isTerminal() bool {
	if n.terminal {
		return true
	}

	return false
}

func (n *Node) child(c string) (*Node, bool) {
	if n.children[c] != nil {
		return n.children[c], true
	}

	return nil, false
}

func (n *Node) setChild(c string, child *Node) {
	n.children[c] = child
}
