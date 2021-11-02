package avl

import (
	"math/rand"
	"reflect"
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

	TestAVL_IsBalanced(t)
}

func TestAVL_IsBalanced(t *testing.T) {
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
	if !tree.IsBalanced() {
		t.Errorf("Tree is not balanced")
	}
}

// func TestAVL_Delete(t *testing.T) {
// 	tree := NewWithIntComparator()
// 	rand.Seed(time.Now().UnixNano())
// 	keyVals := make(map[interface{}]int)
// 	size := 100
// 	for i := 0; i < size; i++ {
// 		key := rand.Int()
// 		got, err := tree.Insert(key, i)
// 		keyVals[got] = i
// 		if err != nil {
// 			t.Errorf("Insert() error = %v", err)
// 			return
// 		}
// 	}
//
// 	for key, _ := range keyVals {
// 		deletedKey, err := tree.Delete(key)
// 		if err != nil {
// 			t.Errorf("Delete() error = %v", err)
// 		}
// 		if !reflect.DeepEqual(deletedKey, key) {
// 			t.Errorf("Delete() got = %v, want %v", deletedKey, key)
// 		}
// 		if !tree.IsBalanced() {
// 			t.Errorf("Tree is not balanced after a deletion")
// 		}
// 	}
//
// 	if !tree.IsBalanced() && tree.Size() != 0 {
// 		t.Errorf("Failed deletion")
// 	}
//
// 	tree.InOrderTraversal()
// }

func TestAVL_Search(t *testing.T) {
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

func TestAVL_ReturnNodeValue(t *testing.T) {
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

func TestAVL_DepthFirstTraversal(t *testing.T) {
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

func TestAVL_InOrderTraversal(t *testing.T) {
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
