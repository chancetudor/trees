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
			if h.isSmaller(leftChildKey, rightChildKey) && h.isSmaller(leftChildKey, currKey) { // left child smaller & heap property is broken
				currKey, leftChildKey = h.swap(currIndex, h.leftChild(currIndex))
				currIndex = h.leftChild(currIndex)
			} else if h.isSmaller(rightChildKey, currKey) { // right child smaller & heap property is broken
				currKey, rightChildKey = h.swap(currIndex, h.rightChild(currIndex))
				currIndex = h.rightChild(currIndex)
			}
		case h.leftChild(currIndex) != -1:
			if h.isSmaller(leftChildKey, currKey) { // heap property broken
				currKey, leftChildKey = h.swap(currIndex, h.leftChild(currIndex))
				currIndex = h.leftChild(currIndex)
			}
		case h.rightChild(currIndex) != -1:
			if h.isSmaller(rightChildKey, currKey) { // heap property broken
				currKey, rightChildKey = h.swap(currIndex, h.rightChild(currIndex))
				currIndex = h.rightChild(currIndex)
			}
		default:
			return
		}
	}
}

// isSmaller returns true if key1 is smaller than key2 and false otherwise.
func (h *Heap) isSmaller(key1 interface{}, key2 interface{}) bool {
	if h.comparator(key1, key2) < 0 {
		return true
	}

	return false
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
	k, v := h.Root().key(), h.Root().value()
	// Copy the last value in the array to the root;
	_, _ = h.swap(0, len(h.store)-1)
	// Decrease heap's size by 1 and overwrite underlying slice.
	h.setSize(h.Size() - 1)
	h.store[len(h.store)-1] = nil
	h.setStore(h.store[:len(h.store)-1])
	h.trickleDown()

	return k, v
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
