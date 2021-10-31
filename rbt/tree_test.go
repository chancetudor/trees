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
	size := 10000
	for i := 0; i < size; i++ {
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

	if tree.Size() != size {
		t.Errorf("Tree size = %v, want %v", tree.Size(), 100)
	}
}

func TestRBT_Delete(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	keyVals := make(map[interface{}]int)
	size := 10000
	for i := 0; i < size; i++ {
		key := rand.Int()
		got, err := tree.Insert(key, i)
		keyVals[got] = i
		if err != nil {
			t.Errorf("Insert() error = %v", err)
			return
		}
	}

	for key, _ := range keyVals {
		deletedKey, err := tree.Delete(key)
		if err != nil {
			t.Errorf("Delete() error = %v", err)
		}
		if !reflect.DeepEqual(deletedKey, key) {
			t.Errorf("Delete() got = %v, want %v", deletedKey, key)
		}
		if !tree.IsBalanced() {
			t.Errorf("Tree is not balanced after a deletion")
		}
	}

	if !tree.IsBalanced() && tree.Size() != 0 {
		t.Errorf("Failed deletion")
	}

	tree.InOrderTraversal()
}

func TestRBT_Search(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	keyVals := make(map[interface{}]int)
	size := 10000
	for i := 0; i < size; i++ {
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
	size := 10000
	for i := 0; i < size; i++ {
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
	size := 10000
	for i := 0; i < size; i++ {
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
	size := 10000
	for i := 0; i < size; i++ {
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
	size := 10000
	for i := 0; i < size; i++ {
		key := rand.Int()
		_, err := tree.Insert(key, i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	if !tree.IsBalanced() && tree.Size() != size {
		t.Errorf("Tree is not balanced")
	}
}

func TestRBT_BlackHeight(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	size := 10000
	for i := 0; i < size; i++ {
		key := rand.Int()
		_, err := tree.Insert(key, i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	height := tree.BlackHeight()
	fmt.Println("Black height = " + strconv.Itoa(height))
}
