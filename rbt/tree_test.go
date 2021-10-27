package rbt

import (
	"github.com/emirpasic/gods/utils"
	"reflect"
	"testing"
)

func TestRBT_Insert(t *testing.T) {
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
				root:       NewNode(1, "1", BLACK),
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
				root:       NewNode("test 3", "1", BLACK),
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
				root:       NewNode('3', "1", BLACK),
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
				root:       NewNode('4', "1", BLACK),
				comparator: utils.RuneComparator,
				size:       1,
			},
			args:    args{key: '4', value: 1},
			wantErr: true,
		},
		{
			name: "Test duplicate insertion with int32",
			fields: fields{
				root:       NewNode(int32(4), "1", BLACK),
				comparator: utils.Int32Comparator,
				size:       1,
			},
			args:    args{key: int32(4), value: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &RBT{
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
