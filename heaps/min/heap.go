package min

import (
	"github.com/emirpasic/gods/utils"
	"math"
)

type Heap struct {
	store      []*Node          // the underlying data store
	comparator utils.Comparator // the key comparator
	size       int              // number of nodes in the tree
}

// NewWith returns a pointer to a Heap where the store is a zero value slice, size is 0,
// and the key comparator is set to the parameter passed in.
// The comparator format is taken from https://github.com/emirpasic/gods#comparator.
// Either import the package https://github.com/emirpasic/gods/utils and pass a comparator from the library,
// or write a custom comparator using guidelines from the gods README.
func NewWith(comparator utils.Comparator) *Heap {
	return &Heap{
		store:      []*Node{},
		comparator: comparator,
		size:       0,
	}
}

// NewWithIntComparator returns a pointer to a Heap where the store is a zero value slice, size is 0,
// and the key comparator is set to the IntComparator from package https://github.com/emirpasic/gods/utils.
func NewWithIntComparator() *Heap {
	return NewWith(utils.IntComparator)
}

// NewWithStringComparator returns a pointer to a Heap where the store is a zero value slice, size is 0,
// and the key comparator is set to the StringComparator from package https://github.com/emirpasic/gods/utils.
func NewWithStringComparator() *Heap {
	return NewWith(utils.StringComparator)
}

func (h *Heap) swap(curr int, replacement int) (interface{}, interface{}) {
	temp := h.store[curr]
	h.store[curr] = h.store[replacement]
	h.store[replacement] = temp

	return h.store[curr].key(), h.store[replacement].key()
}

// trickleUp rearranges the elements in the heap to maintain the heap property.
// The function takes the last element in the heap and bubbles it up until the heap property is satisfied.
func (h *Heap) trickleUp() {
	currIndex := len(h.store) - 1
	parentIndex := h.parentIndex(currIndex)
	currKey := h.store[currIndex].key()
	parentKey := h.store[parentIndex].key()
	compare := h.comparator(currKey, parentKey)

	for currIndex >= 0 && compare < 0 {
		currKey, parentKey = h.swap(currIndex, parentIndex)
		currIndex = h.parentIndex(currIndex)
		parentIndex = h.parentIndex(currIndex)
		compare = h.comparator(currKey, parentKey)
	}
	_, _ = h.swap(currIndex, parentIndex)
}

// trickleDown rearranges the elements in the heap to maintain the heap property.
// The function takes the first element in the heap and trickles it down
// until the heap property is satisfied.
func (h *Heap) trickleDown() {
	// if current node has no children, sifting is over;
	// if current node has one child: check, if heap property is broken,
	// then swap current node's value and child value; sift down the child;
	// if current node has two children: find the smallest of them.
	// If heap property is broken, then swap current node's value and selected child value; sift down the child.
	if h.leftChild(0) == -1 && h.rightChild(0) == -1 {
		return
	}
	currIndex := 0
	currKey := h.store[currIndex].key()
	leftChildKey := h.store[h.leftChild(currIndex)].key()
	rightChildKey := h.store[h.rightChild(currIndex)].key()
	for currIndex < len(h.store) {
		// three cases:
		// 1) left and right child; find smallest of them and swap
		// 2) left child; if heap property is broken, swap
		// 3) right child; if heap property is broken, swap
		switch {
		case h.leftChild(currIndex) != -1 && h.rightChild(currIndex) != -1:
			if h.comparator(leftChildKey, rightChildKey) < 0 { // left child smaller
				if h.comparator(leftChildKey, currKey) < 0 { // heap property broken
					currKey, leftChildKey = h.swap(currIndex, h.leftChild(currIndex))
					currIndex = h.leftChild(currIndex)
				}
			} else if h.comparator(h.store[h.leftChild(currIndex)].key(), h.store[h.rightChild(currIndex)].key()) > 0 { // right child smaller
				if h.comparator(rightChildKey, currKey) < 0 { // heap property broken
					currKey, rightChildKey = h.swap(currIndex, h.rightChild(currIndex))
					currIndex = h.rightChild(currIndex)
				}
			}
		case h.leftChild(0) != -1:
			if h.comparator(leftChildKey, currKey) < 0 { // heap property broken
				currKey, leftChildKey = h.swap(currIndex, h.leftChild(currIndex))
				currIndex = h.leftChild(currIndex)
			}
		case h.rightChild(0) != -1:
			if h.comparator(rightChildKey, currKey) < 0 { // heap property broken
				currKey, rightChildKey = h.swap(currIndex, h.rightChild(currIndex))
				currIndex = h.rightChild(currIndex)
			}
		}
	}
}

// Insert adds an item to a heap while maintaining its heap property.
// Insertion is done by key; that is, the keys that are inserted are
// what maintains the heap property.
func (h *Heap) Insert(key, val interface{}) interface{} {
	newNode := NewNode(key, val)
	h.store = append(h.store, newNode)
	h.trickleUp()
	h.size = len(h.store)

	return newNode.key()
}

// DeleteMin removes the minimum from the heap and returns the key and value of the deleted node.
func (h *Heap) DeleteMin() (interface{}, interface{}) {
	// Copy the last value in the array to the root;
	_, _ = h.swap(0, len(h.store)-1)
	// Decrease heap's size by 1;
	h.setSize(h.Size() - 1)
	h.setStore(h.store[0 : len(h.store)-1])
	// Sift down root's value. Sifting is done as following:
	h.trickleDown()

	return nil, nil
}

// IsEmpty returns true if heap is empty and false if it has a node.
func (h *Heap) IsEmpty() bool {
	if h.size == 0 {
		return true
	}

	return false
}

// Size returns the size of the heap.
func (h *Heap) Size() int {
	return len(h.store)
}

// setSize sets a new size for the heap.
func (h *Heap) setSize(newSize int) {
	h.size = newSize
}

// GetMin returns the minimum value in the heap.
func (h *Heap) GetMin() interface{} {
	return h.store[0].value()
}

// Root returns the root node of the heap.
func (h *Heap) Root() *Node {
	return h.store[0]
}

// parentIndex returns the index of the given node's parent.
func (h *Heap) parentIndex(index int) int {
	return int(math.Floor(float64((index - 1) / 2)))
}

// leftChild returns the index of the given node's left child.
// If the child index is greater than the size of the heap, -1 is returned.
func (h *Heap) leftChild(index int) int {
	left := 2 * (index + 1)
	if left > h.Size() {
		return -1
	}

	return left
}

// rightChild returns the index of the given node's left child.
// If the child index is greater than the size of the heap, -1 is returned.
func (h *Heap) rightChild(index int) int {
	right := 2 * (index + 2)
	if right > h.Size() {
		return -1
	}

	return right
}

func (h *Heap) setStore(newHeap []*Node) {
	h.store = newHeap
}

// // Update takes a key and a value and updates a node with the existing key with the new value.
// // Returns the new value of the node or an error, if there was one.
// func (h *Heap) Update(key interface{}, value interface{}) (interface{}, error) {
// 	matchingNode, err := h.findNode(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	matchingNode.setValue(value)
//
// 	return matchingNode.value(), nil
// }
//
// // ReturnNodeValue takes a key and returns the value associated with the key or an error, if there was one.
// func (h *Heap) ReturnNodeValue(key interface{}) (interface{}, error) {
// 	matchingNode, err := h.findNode(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return matchingNode.value(), nil
// }
//
// // Search takes a key and searches for the key in the tree.
// // The function returns a boolean, stating whether the key was found or not.
// func (h *Heap) Search(key interface{}) bool {
// 	_, err := h.findNode(key)
// 	if err != nil {
// 		return false
// 	}
//
// 	return true
// }
//
// // findNode takes a key and returns the node associated with that key.
// // Returns nil and an error if no node exists.
// func (h *Heap) findNode(key interface{}) (*Node, error) {
// 	tempNode := h.Root()
// 	for tempNode != nil {
// 		compare := h.comparator(key, tempNode.key())
// 		switch {
// 		case compare < 0:
// 			tempNode = tempNode.leftChild()
// 		case compare > 0:
// 			tempNode = tempNode.rightChild()
// 		case compare == 0:
// 			return tempNode, nil
// 		}
// 	}
//
// 	return nil, NewNilNodeError(key)
// }
