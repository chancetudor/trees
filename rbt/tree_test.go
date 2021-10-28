package rbt

import (
	"fmt"
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

func TestRBT_DepthFirstTraversal(t *testing.T) {
	tree := &RBT{
		root:       nil,
		comparator: utils.IntComparator,
		size:       0,
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println("DEPTH FIRST")
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
	tree.DepthFirstTraversal()
}

func TestRBT_InOrderTraversal(t *testing.T) {
	tree := &RBT{
		root:       nil,
		comparator: utils.IntComparator,
		size:       0,
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println("IN ORDER")
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
	tree.InOrderTraversal()
}

func TestRBT_IsBalanced(t *testing.T) {
	tree := NewWithIntComparator()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 7; i++ {
		key := rand.Int()
		_, err := tree.Insert(key, i)
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
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 17; i++ {
		key := rand.Int()
		_, err := tree.Insert(key, i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	height := tree.BlackHeight()
	fmt.Println(height)
}
