package rbt

import (
	"github.com/emirpasic/gods/utils"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestRBT_Insert(t *testing.T) {
	tree := &RBT{
		root:       nil,
		comparator: utils.IntComparator,
		size:       0,
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		key := rand.Int()
		got, err := tree.Insert(key, i)
		if err != nil {
			t.Errorf("Insert() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, key) {
			t.Errorf("Insert() got = %v, want %v", got, key)
		}
	}
}
