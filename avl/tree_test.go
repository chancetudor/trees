package avl

import (
	"math/rand"
	"testing"
	"time"
)

func TestAVL_Insert(t *testing.T) {
	tree := NewWithIntComparator()
	size := 100
	rand.Seed(time.Now().Unix())
	keyVals := make(map[interface{}]int)
	for i := 0; i < size; i++ {
		key := rand.Int()
		got, err := tree.Insert(key, i)
		if err != nil {
			t.Errorf(err.Error())
		}
		keyVals[got] = i
	}

	if tree.Size() != size {
		t.Errorf("Size after insert = %v, want %d", tree.Size(), size)
	}

	for k, _ := range keyVals {
		if !tree.Search(k) {
			t.Errorf("Cannot find key %v in tree", k)
		}
	}
}
