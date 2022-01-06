package list

import (
	"reflect"
	"testing"
)

func TestAddInList(t *testing.T) {
	// case 1
	H1, _ := Create([]int{5,9,7,5,7,1,2,6,4,2,7,8,9,6,1,6,6,1,1,4,2,9,5,5,0,4,6,3,0,4,3,5,6,7,0,5,5,4,4,0})
	H2, _ := Create([]int{1,3,2,5,0,6,0,2,1,4,3,9,3,0,9,9,0,3,1,6,5,7,8,6,2,3,8,5,0,9,7,9,4,5,9,9,4,9,3,6})
	Rst, _ := Create([]int{7,3,0,0,7,7,2,8,5,7,1,8,2,7,1,5,6,4,3,0,8,7,4,1,2,8,4,8,1,4,1,5,1,3,0,5,0,3,7,6})
	type args struct {
		head1 *ListNode
		head2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{head1:H1, head2: H2},
			want: Rst,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddInList(tt.args.head1, tt.args.head2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddInList() = %v, want %v", got, tt.want)
			}
		})
	}
}
