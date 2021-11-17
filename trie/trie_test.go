package trie

import (
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
		{name: "eggs", word: "eggs", want: true},
		{name: "bacon", word: "bacon", want: true},
		{name: "pancakes", word: "pancakes", want: true},
	}
	t := NewWithStringComparator()
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := t.Insert(tt.word); got != tt.want {
				t1.Errorf("Insert(%v) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestTrie_Delete(t1 *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{name: "success", word: "hello", want: true},
		{name: "success", word: "world", want: true},
		{name: "word not existent", word: "eggs", want: false},
	}
	t := NewWithStringComparator()
	_ = t.Insert("hello")
	_ = t.Insert("world")
	_ = t.Insert("bacon")
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := t.Delete(tt.word); got != tt.want {
				t1.Errorf("Delete(%v) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestTrie_IsPrefix(t1 *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{name: "prefix of 'hello'", word: "h", want: true},
		{name: "prefix of 'world'", word: "wor", want: true},
		{name: "word not existent", word: "egg", want: false},
	}
	t := NewWithStringComparator()
	_ = t.Insert("hello")
	_ = t.Insert("world")
	_ = t.Insert("bacon")
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := t.IsPrefix(tt.word); got != tt.want {
				t1.Errorf("IsPrefix(%v) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
