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
		deletedKey, err := tree.Delete(key)
		if err != nil {
			t.Errorf("Delete() error = %v", err)
		}
		if !reflect.DeepEqual(deletedKey, key) {
			t.Errorf("Delete() got = %v, want %v", deletedKey, key)
		}
	}

	if !tree.IsBalanced() {
		t.Errorf("Tree is not balanced after all deletions")
	}
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
		if !tree.Search(key) {
			t.Errorf("Cannot find key %d", key)
		}
	}
}

func TestRBT_DepthFirstTraversal(t *testing.T) {
	tree := NewWithIntComparator()
	for i := 0; i < 100; i++ {
		// key := rand.Int()
		got, err := tree.Insert(i, i)
		if err != nil {
			t.Errorf("Insert() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, i) {
			t.Errorf("Insert() got = %v, want %v", got, i)
		}
	}
	tree.DepthFirstTraversal()
}

func TestRBT_InOrderTraversal(t *testing.T) {
	tree := NewWithIntComparator()
	for i := 0; i < 100; i++ {
		// key := rand.Int() % 10
		_, err := tree.Insert(i, i)
		if err != nil {
			t.Errorf("Insert() error = %v", err)
		}
	}
	tree.InOrderTraversal()
}

func TestRBT_IsBalanced(t *testing.T) {
	tree := NewWithIntComparator()
	for i := 0; i < 100; i++ {
		_, err := tree.Insert(i, i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	if !tree.IsBalanced() {
		t.Errorf("tree is not balanced")
	}
}

func TestRBT_BlackHeight(t *testing.T) {
	tree := NewWithIntComparator()
	for i := 0; i < 100; i++ {
		_, err := tree.Insert(i, i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	height := tree.BlackHeight()
	fmt.Println("Black height = " + strconv.Itoa(height))
	tree.InOrderTraversal()
}
