package rbt

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestRBT_Insert(t *testing.T) {
	tree := NewWithIntComparator()
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

	if tree.Size() != 100 {
		t.Errorf("Tree size = %v, want %v", tree.Size(), 100)
	}
}

func TestRBT_Delete(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	keyVals := make(map[interface{}]int)
	for i := 0; i < 100; i++ {
		// key := rand.Int()
		got, err := tree.Insert(i, i)
		keyVals[got] = i
		if err != nil {
			t.Errorf("Insert() error = %v", err)
			return
		}
	}

	if tree.Size() != 100 {
		t.Errorf("Tree does not have correct size")
		return
	}

	// tree.InOrderTraversal()

	// for key, _ := range keyVals {
	// 	deletedKey, err := tree.Delete(key)
	// 	if err != nil {
	// 		t.Errorf("Delete() error = %v", err)
	// 		if !tree.Search(key) {
	// 			t.Errorf("Search() error = %v", err)
	// 			return
	// 		}
	// 	}
	// 	if !reflect.DeepEqual(deletedKey, key) {
	// 		t.Errorf("Delete() got = %v, want %v", deletedKey, key)
	// 	}
	// 	if !tree.IsBalanced() {
	// 		t.Errorf("Tree is not balanced after a deletion")
	// 	}
	// }

	if tree.Root() == nil {
		t.Errorf("Root is nil?")
	}

	for i := 0; i < 100; i++ {
		deletedKey, err := tree.Delete(i)
		if err != nil {
			t.Errorf("Delete() error = %v", err)
			if !tree.Search(i) {
				t.Errorf("Search() error = %v", err)
				//return
			}
		}
		if !reflect.DeepEqual(deletedKey, i) {
			t.Errorf("Delete() got = %v, want %v", deletedKey, i)
		}
		if !tree.IsBalanced() {
			t.Errorf("Tree is not balanced after a deletion")
		}
	}

	if !tree.IsBalanced() {
		t.Errorf("Tree is not balanced after ALL deletions")
	}

	tree.InOrderTraversal()
}

func TestRBT_Search(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	keyVals := make(map[interface{}]int)
	for i := 0; i < 100; i++ {
		key := rand.Int()
		got, err := tree.Insert(key, i)
		keyVals[got] = i
		if err != nil {
			t.Errorf("Insert() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, key) {
			t.Errorf("Insert() got = %v, want %v", got, key)
		}
	}

	for key, _ := range keyVals {
		if !tree.Search(key) {
			t.Errorf("Cannot find key %d", key)
		}
	}
}

func TestRBT_ReturnNodeValue(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	keyVals := make(map[interface{}]int)
	for i := 0; i < 100; i++ {
		key := rand.Int()
		got, err := tree.Insert(key, i)
		keyVals[got] = i
		if err != nil {
			t.Errorf("Insert() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, key) {
			t.Errorf("Insert() got = %v, want %v", got, key)
		}
	}

	for key, _ := range keyVals {
		val, _ := tree.ReturnNodeValue(key)
		if keyVals[key] != val.(int) {
			t.Errorf("Cannot find key %d", key)
		}
	}
}

func TestRBT_DepthFirstTraversal(t *testing.T) {
	tree := NewWithIntComparator()
	for i := 0; i < 100; i++ {
		key := rand.Int()
		_, err := tree.Insert(key, i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	if !tree.IsBalanced() && tree.Size() != 100 {
		t.Errorf("Tree is not balanced")
	}
	tree.DepthFirstTraversal()
}

func TestRBT_InOrderTraversal(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		key := rand.Int()
		_, err := tree.Insert(key, i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	if !tree.IsBalanced() && tree.Size() != 100 {
		t.Errorf("Tree is not balanced")
	}
	tree.InOrderTraversal()
}

func TestRBT_IsBalanced(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		key := rand.Int()
		_, err := tree.Insert(key, i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	if !tree.IsBalanced() && tree.Size() != 100 {
		t.Errorf("Tree is not balanced")
	}
}

func TestRBT_BlackHeight(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		key := rand.Int()
		_, err := tree.Insert(key, i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	height := tree.BlackHeight()
	fmt.Println("Black height = " + strconv.Itoa(height))
}
