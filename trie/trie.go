package trie

import (
	"github.com/emirpasic/gods/utils"
	"strings"
)

type Trie struct {
	root       *Node
	comparator utils.Comparator
	size       int // represents how many full words are in the Trie
}

// NewWithStringComparator returns a pointer to a Trie where root is initialized, size is 0,
// and the key comparator is set to the StringComparator from package https://github.com/emirpasic/gods/utils.
// the comparator format is taken from https://github.com/emirpasic/gods#comparator.
func NewWithStringComparator() *Trie {
	return &Trie{
		root:       NewNode(),
		comparator: utils.StringComparator,
		size:       0,
	}
}

// Insert takes a string and inserts the string into the Trie. The function returns a boolean.
// If the string is not already in the Trie, true is returned. Otherwise, false is returned.
func (t *Trie) Insert(word string) bool {
	if len(word) == 0 {
		return false
	}
	chars := strings.Split(word, "")
	temp := t.root
	for _, char := range chars {
		if child, ok := temp.child(char); ok {
			temp = child
		} else {
			newChild := NewNode()
			temp.setChild(char, newChild)
			temp = newChild
		}
		if temp.isTerminal() {
			return false
		}
	}
	temp.markEnd()
	t.size++

	return true
}

// Delete takes a word and deletes it from the Trie. The function returns
// true if the word has been deleted and false if the word does not exist.
func (t *Trie) Delete(word string) bool {
	if t.Size() == 0 || len(word) == 0 {
		return false
	}
	// chars := strings.Split(word, "")
	// TODO implement

	t.size--

	return true
}

// IsPrefix
func (t *Trie) IsPrefix(word string) bool {
	if len(word) == 0 || t.Size() == 0 {
		return false
	}
	chars := strings.Split(word, "")
	temp := t.root
	for _, char := range chars {
		child, ok := temp.child(char)
		if !ok {
			return false
		}
		temp = child
	}

	return true
}

// Exists takes a word and returns false if the word is not in the Trie
// and false otherwise.
func (t *Trie) Exists(word string) bool {
	if len(word) == 0 || t.Size() == 0 {
		return false
	}
	chars := strings.Split(word, "")
	temp := t.root
	for _, char := range chars {
		child, ok := temp.child(char)
		if !ok {
			return false
		}
		temp = child
	}

	return temp.isTerminal()
}

// Size returns the total number of words in the Trie.
func (t *Trie) Size() int {
	return t.size
}
