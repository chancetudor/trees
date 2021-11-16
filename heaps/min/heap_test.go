package min

import (
	"fmt"
	"testing"
)

func TestHeap_Insert(t *testing.T) {
	heap := NewWithIntComparator()
	k := heap.Insert(1, 2)
	if k != 1 {
		t.Errorf("Got %d, want %d", k, 1)
	}
	k = heap.Insert(2, 3)
	if k != 2 {
		t.Errorf("Got %d, want %d", k, 2)
	}
	k = heap.Insert(3, 4)
	if k != 3 {
		t.Errorf("Got %d, want %d", k, 3)
	}
	k = heap.Insert(0, 0)
	if k != 0 {
		t.Errorf("Got %d, want %d", k, 0)
	}
	k = heap.Insert(-1, 1241)
	if k != -1 {
		t.Errorf("Got %d, want %d", k, -1)
	}
	k = heap.Insert(-100, 1241)
	if k != -100 {
		t.Errorf("Got %d, want %d", k, -100)
	}
	if heap.Root().key() != -100 {
		t.Errorf("Got %d, want %d", heap.Root().key(), -100)
	}
}

func TestHeap_DeleteMin(t *testing.T) {
	heap := NewWithIntComparator()
	k := heap.Insert(1, 2)
	if k != 1 {
		t.Errorf("Got %d, want %d", k, 1)
	}
	k = heap.Insert(2, 3)
	if k != 2 {
		t.Errorf("Got %d, want %d", k, 2)
	}
	k = heap.Insert(3, 4)
	if k != 3 {
		t.Errorf("Got %d, want %d", k, 3)
	}
	k = heap.Insert(0, 0)
	if k != 0 {
		t.Errorf("Got %d, want %d", k, 0)
	}
	k = heap.Insert(-1, 1241)
	if k != -1 {
		t.Errorf("Got %d, want %d", k, -1)
	}
	k = heap.Insert(-100, 1241)
	if k != -100 {
		t.Errorf("Got %d, want %d", k, -100)
	}
	for _, v := range heap.store {
		fmt.Println(v.Key, v.Value)
	}
	k, _ = heap.DeleteMin()
	if k != -100 {
		t.Errorf("Got %d, want %d", k, -100)
	}
	for _, v := range heap.store {
		fmt.Println(v.Key, v.Value)
	}
	if heap.Root().key() != -1 {
		t.Errorf("Got %v, want %d", heap.Root().key(), -1)
	}
}
