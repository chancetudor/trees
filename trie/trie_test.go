package trie

import (
	"fmt"
	"testing"
)

func TestTrie_Insert(t1 *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{name: "hello", word: "hello", want: true},
		{name: "world", word: "world", want: true},
		{name: "hello2", word: "hello", want: false},
	}
	t := NewWithStringComparator()
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := t.Insert(tt.word); got != tt.want {
				t1.Errorf("Insert(%v) = %v, want %v", tt.word, got, tt.want)
			}
			fmt.Println(t.root.children)
		})
	}
}
