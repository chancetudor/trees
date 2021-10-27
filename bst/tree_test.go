package bst

import (
	"github.com/emirpasic/gods/utils"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestBST_InsertThenUpdate(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "test 1",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			args: args{
				key:   1,
				value: "2",
			},
			want:    "2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			keyVals := make(map[int]int)
			for i := 0; i < 100; i++ {
				rand.Seed(time.Now().UnixNano())
				k := rand.Int()
				v := rand.Int()
				keyVals[k] = v
				_, err := tree.Insert(k, v)
				if err != nil {
					t.Error(err)
				}
			}
			for keys, _ := range keyVals {
				newValue, err := tree.Update(keys, "new value")
				if err != nil {
					t.Errorf(err.Error())
				}
				if newValue != "new value" {
					t.Errorf("Delete() = %v, want %v", newValue, "new value")
				}
			}
		})
	}
}

func TestBST_InsertThenSize(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test 1",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			want: 101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			keyVals := make(map[int]int)
			for i := 0; i < 100; i++ {
				rand.Seed(time.Now().UnixNano())
				k := rand.Int()
				v := rand.Int()
				keyVals[k] = v
				_, err := tree.Insert(k, v)
				if err != nil {
					t.Error(err)
				}
			}
			if got := tree.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_InsertThenSearch(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test 1",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			keyVals := make(map[int]int)
			for i := 0; i < 100; i++ {
				rand.Seed(time.Now().UnixNano())
				k := rand.Int()
				v := rand.Int()
				keyVals[k] = v
				_, err := tree.Insert(k, v)
				if err != nil {
					t.Error(err)
				}
			}
			for keys, _ := range keyVals {
				found := tree.Search(keys)
				if found != true {
					t.Errorf("Delete() = %v, want %v", found, true)
				}
			}
		})
	}
}

func TestBST_InsertThenDelete(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test 1",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			keyVals := make(map[int]int)
			for i := 0; i < 100; i++ {
				rand.Seed(time.Now().UnixNano())
				k := rand.Int()
				v := rand.Int()
				keyVals[k] = v
				_, err := tree.Insert(k, v)
				if err != nil {
					t.Error(err)
				}
			}
			for keys, _ := range keyVals {
				key, err := tree.Delete(keys)
				if err != nil {
					t.Error(err)
				}
				if key != keys {
					t.Errorf("Delete() = %v, want %v", key, keys)
				}
			}
		})
	}
}

func TestBST_InsertThenClear(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test 1",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			for i := 0; i < 100; i++ {
				rand.Seed(time.Now().UnixNano())
				k := rand.Int()
				v := rand.Int()
				_, err := tree.Insert(k, v)
				if err != nil {
					t.Error(err)
				}
			}
			tree.Clear()
			if tree.root != nil {
				t.Error("Root is not nil")
			}
			if tree.size != 0 {
				t.Error("Size is not 0")
			}
		})
	}
}

func TestBST_InsertThenDFS(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test 1",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			for i := 0; i < 100; i++ {
				rand.Seed(time.Now().UnixNano())
				k := rand.Int()
				v := rand.Int()
				_, err := tree.Insert(k, v)
				if err != nil {
					t.Error(err)
				}
			}
			tree.InOrderTraversal()
		})
	}
}

func TestBST_Clear(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test 1",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			tree.Clear()
			if tree.root != nil {
				t.Error("Root is not nil")
			}
			if tree.size != 0 {
				t.Error("Size is not 0")
			}
		})
	}
}

func TestBST_Delete(t *testing.T) {
	type fields struct {
		root         *Node
		comparator   utils.Comparator
		size         int
		originalSize int
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test deleting with int comparator",
			fields: fields{
				root:         NewNode(1, "1"),
				comparator:   utils.IntComparator,
				size:         1,
				originalSize: 1,
			},
			args:    args{key: 1},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test deleting with string comparator",
			fields: fields{
				root:         NewNode("test 2", "1"),
				comparator:   utils.StringComparator,
				size:         1,
				originalSize: 1,
			},
			args:    args{key: "test 2"},
			want:    "test 2",
			wantErr: false,
		},
		{
			name: "Test deleting with rune comparator",
			fields: fields{
				root:         NewNode('1', "1"),
				comparator:   utils.RuneComparator,
				size:         1,
				originalSize: 1,
			},
			args:    args{key: '1'},
			want:    '1',
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			got, err := tree.Delete(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
			if tree.size != (tt.fields.originalSize - 1) {
				t.Errorf("Size after delete = %v, want %v", tree.size, tt.fields.originalSize)
			}
		})
	}
}

func TestBST_Insert(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test insertion with int comparator",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			args:    args{key: 2, value: 1},
			want:    2,
			wantErr: false,
		},
		{
			name: "Test insertion with string comparator",
			fields: fields{
				root:       NewNode("test 3", "1"),
				comparator: utils.StringComparator,
				size:       1,
			},
			args:    args{key: "new string key", value: 1},
			want:    "new string key",
			wantErr: false,
		},
		{
			name: "Test insertion with rune comparator",
			fields: fields{
				root:       NewNode('3', "1"),
				comparator: utils.RuneComparator,
				size:       1,
			},
			args:    args{key: '5', value: 1},
			want:    '5',
			wantErr: false,
		},
		{
			name: "Test duplicate insertion",
			fields: fields{
				root:       NewNode('4', "1"),
				comparator: utils.RuneComparator,
				size:       1,
			},
			args:    args{key: '4', value: 1},
			wantErr: true,
		},
		{
			name: "Test duplicate insertion with int32",
			fields: fields{
				root:       NewNode(int32(4), "1"),
				comparator: utils.Int32Comparator,
				size:       1,
			},
			args:    args{key: int32(4), value: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			got, err := tree.Insert(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_IsEmpty(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Empty tree",
			fields: fields{
				root:       nil,
				comparator: utils.IntComparator,
				size:       0,
			},
			want: true,
		},
		{
			name: "Non-empty tree",
			fields: fields{
				root:       NewNode(1, 2),
				comparator: utils.IntComparator,
				size:       1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			if got := tree.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_ReturnNodeValue(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test insertion with int comparator",
			fields: fields{
				root:       NewNode(1, 1),
				comparator: utils.IntComparator,
				size:       1,
			},
			args:    args{key: 1},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test insertion with string comparator",
			fields: fields{
				root:       NewNode("test 3", "test 3"),
				comparator: utils.StringComparator,
				size:       1,
			},
			args:    args{key: "test 3"},
			want:    "test 3",
			wantErr: false,
		},
		{
			name: "Test insertion with string comparator",
			fields: fields{
				root:       NewNode("test 4", "test 3"),
				comparator: utils.StringComparator,
				size:       1,
			},
			args:    args{key: "new string key"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			got, err := tree.ReturnNodeValue(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnNodeValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReturnNodeValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Root(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name: "test 1",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			want: NewNode(1, "1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			if got := tree.Root(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Search(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Found key",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			args: args{key: 1},
			want: true,
		},
		{
			name: "Not found key",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			args: args{key: 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			if got := tree.Search(tt.args.key); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Size(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		tree   *BST
	}{
		{
			name: "Test with just root",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			want: 1,
		},
		{
			name: "Test with insertion",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			if tt.name == "Test with insertion" {
				_, err := tree.Insert(2, "new insertion")
				if err != nil {
					t.Error(err)
				}
			}
			if got := tree.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Update(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Test update with int comparator",
			fields: fields{
				root:       NewNode(1, 1),
				comparator: utils.IntComparator,
				size:       1,
			},
			args:    args{key: 1, value: 2},
			want:    2,
			wantErr: false,
		},
		{
			name: "Test failed update with int comparator",
			fields: fields{
				root:       NewNode(1, 1),
				comparator: utils.IntComparator,
				size:       1,
			},
			args:    args{key: 2, value: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			got, err := tree.Update(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_findNode(t *testing.T) {
	type fields struct {
		root       *Node
		comparator utils.Comparator
		size       int
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Node
		wantErr bool
	}{
		{
			name: "Test finding node -- no error",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			args:    args{key: 1},
			want:    NewNode(1, "1"),
			wantErr: false,
		},
		{
			name: "Test deleting with int comparator",
			fields: fields{
				root:       NewNode(1, "1"),
				comparator: utils.IntComparator,
				size:       1,
			},
			args:    args{key: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BST{
				root:       tt.fields.root,
				comparator: tt.fields.comparator,
				size:       tt.fields.size,
			}
			got, err := tree.findNode(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("findNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestBST_pruneLeaf(t *testing.T) {
// 	type fields struct {
// 		root       *Node
// 		comparator utils.Comparator
// 		size       int
// 	}
// 	type args struct {
// 		toDelete *Node
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tree := &BST{
// 				root:       tt.fields.root,
// 				comparator: tt.fields.comparator,
// 				size:       tt.fields.size,
// 			}
// 		})
// 	}
// }
//
// func TestBST_replaceSubTree(t *testing.T) {
// 	type fields struct {
// 		root       *Node
// 		comparator utils.Comparator
// 		size       int
// 	}
// 	type args struct {
// 		toDelete        *Node
// 		replacementNode *Node
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tree := &BST{
// 				root:       tt.fields.root,
// 				comparator: tt.fields.comparator,
// 				size:       tt.fields.size,
// 			}
// 		})
// 	}
// }
//
// func TestBST_setRoot(t *testing.T) {
// 	type fields struct {
// 		root       *Node
// 		comparator utils.Comparator
// 		size       int
// 	}
// 	type args struct {
// 		newRoot *Node
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tree := &BST{
// 				root:       tt.fields.root,
// 				comparator: tt.fields.comparator,
// 				size:       tt.fields.size,
// 			}
// 		})
// 	}
// }
//
// func TestBST_setSize(t *testing.T) {
// 	type fields struct {
// 		root       *Node
// 		comparator utils.Comparator
// 		size       int
// 	}
// 	type args struct {
// 		newSize int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tree := &BST{
// 				root:       tt.fields.root,
// 				comparator: tt.fields.comparator,
// 				size:       tt.fields.size,
// 			}
// 		})
// 	}
// }

// func TestNewNode(t *testing.T) {
// 	type args struct {
// 		k interface{}
// 		v interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *Node
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewNode(tt.args.k, tt.args.v); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewNode() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNewWith(t *testing.T) {
// 	type args struct {
// 		comparator utils.Comparator
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *BST
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewWith(tt.args.comparator); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewWith() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNewWithIntComparator(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want *BST
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewWithIntComparator(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewWithIntComparator() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNewWithStringComparator(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want *BST
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewWithStringComparator(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewWithStringComparator() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_clear(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 		})
// 	}
// }
//
// func TestNode_getParent(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   *Node
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.getParent(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("getParent() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_isLeaf(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.isLeaf(); got != tt.want {
// 				t.Errorf("isLeaf() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_isRoot(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.isRoot(); got != tt.want {
// 				t.Errorf("isRoot() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_key(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   interface{}
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.key(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("key() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_leftChild(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   *Node
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.leftChild(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("leftChild() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_predecessor(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   *Node
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.predecessor(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("predecessor() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_rightChild(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   *Node
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.rightChild(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("rightChild() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_setKey(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	type args struct {
// 		key interface{}
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 		})
// 	}
// }
//
// func TestNode_setLeftChild(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	type args struct {
// 		left *Node
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 		})
// 	}
// }
//
// func TestNode_setParent(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	type args struct {
// 		parent *Node
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 		})
// 	}
// }
//
// func TestNode_setRightChild(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	type args struct {
// 		right *Node
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 		})
// 	}
// }
//
// func TestNode_setValue(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	type args struct {
// 		value interface{}
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 		})
// 	}
// }
//
// func TestNode_subtreeMax(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	type args struct {
// 		child *Node
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   *Node
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.subtreeMax(tt.args.child); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("subtreeMax() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_subtreeMin(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	type args struct {
// 		child *Node
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   *Node
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.subtreeMin(tt.args.child); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("subtreeMin() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_successor(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   *Node
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.successor(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("successor() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestNode_value(t *testing.T) {
// 	type fields struct {
// 		left   *Node
// 		right  *Node
// 		parent *Node
// 		Data   *NodeData
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   interface{}
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &Node{
// 				left:   tt.fields.left,
// 				right:  tt.fields.right,
// 				parent: tt.fields.parent,
// 				Data:   tt.fields.Data,
// 			}
// 			if got := node.value(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("value() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
