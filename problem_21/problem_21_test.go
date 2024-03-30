package problem_21

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Example 2",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "Example 3",
			args: args{
				list1: nil,
				list2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
		{
			name: "List 1 not empty",
			args: args{
				list1: &ListNode{
					Val:  0,
					Next: nil,
				},
				list2: nil,
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendNodes(t *testing.T) {
	type args struct {
		list     LinkedList
		listNode *ListNode
	}
	tests := []struct {
		name string
		args args
		want LinkedList
	}{
		{
			name: "Two elements",
			args: args{
				list: LinkedList{
					head:   nil,
					length: 0,
				},
				listNode: &ListNode{
					Val: 10,
					Next: &ListNode{
						Val:  20,
						Next: nil,
					},
				},
			},
			want: LinkedList{
				head: &ListNode{
					Val: 10,
					Next: &ListNode{
						Val:  20,
						Next: nil,
					},
				},
				length: 2,
			},
		},
		{
			name: "One element",
			args: args{
				list: LinkedList{
					head:   nil,
					length: 0,
				},
				listNode: &ListNode{
					Val:  10,
					Next: nil,
				},
			},
			want: LinkedList{
				head: &ListNode{
					Val:  10,
					Next: nil,
				},
				length: 1,
			},
		},
		{
			name: "Without elements",
			args: args{
				list: LinkedList{
					head: &ListNode{
						Val:  10,
						Next: nil,
					},
					length: 1,
				},
				listNode: nil,
			},
			want: LinkedList{
				head: &ListNode{
					Val:  10,
					Next: nil,
				},
				length: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendNodes(tt.args.list, tt.args.listNode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendNode(t *testing.T) {
	var someValue = 20

	type args struct {
		list  LinkedList
		value int
	}
	tests := []struct {
		name string
		args args
		want LinkedList
	}{
		{
			name: "Empty list",
			args: args{
				list:  LinkedList{},
				value: someValue,
			},
			want: LinkedList{
				head: &ListNode{
					Val:  someValue,
					Next: nil,
				},
				length: 1,
			},
		},
		{
			name: "One value",
			args: args{
				list: LinkedList{
					head: &ListNode{
						Val:  10,
						Next: nil,
					},
					length: 1,
				},
				value: someValue,
			},
			want: LinkedList{
				head: &ListNode{
					Val: 10,
					Next: &ListNode{
						Val:  someValue,
						Next: nil,
					},
				},
				length: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendNode(tt.args.list, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
