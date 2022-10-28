package problem_876

import (
	"reflect"
	"testing"
)

type args struct {
	head *ListNode
}

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test case with ([1]=>[2]=>[3]=>[4]=>[5]) list input",
			args: args{
				head: makeListNode(1, 5),
			},
			want: makeListNode(3, 3),
		},
		{
			name: "Test case with ([1]=>[2]=>[3]=>[4]=>[5]) list input",
			args: args{
				head: makeListNode(1, 6),
			},
			want: makeListNode(4, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, wantList %v", got, tt.want)
			}
		})
	}
}

func Test_middleNodeByPointer(t *testing.T) {
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Test case with ([1]=>[2]=>[3]=>[4]=>[5]) list input",
			args: args{
				head: makeListNode(1, 5),
			},
			want: makeListNode(3, 3),
		},
		{
			name: "Test case with ([1]=>[2]=>[3]=>[4]=>[5]) list input",
			args: args{
				head: makeListNode(1, 6),
			},
			want: makeListNode(4, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNodeByPointer(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNodeByPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNodes(t *testing.T) {
	tests := []struct {
		name      string
		args      args
		wantList  []*ListNode
		wantCount int
	}{
		{
			name: "Test case with [[1]=>[2]=>[3]] input",
			args: args{
				head: makeListNode(1, 3),
			},
			wantList:  makeListOfNodes(1, 3),
			wantCount: 3,
		},
		{
			name: "Test case with [[3]=>[4]=>[5]] input",
			args: args{
				head: makeListNode(3, 3),
			},
			wantList:  makeListOfNodes(3, 3),
			wantCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getNodesAsList(tt.args.head)
			if !reflect.DeepEqual(got, tt.wantList) {
				t.Errorf("getNodesAsList() got = %v, wantList %v", got, tt.wantList)
			}
			if got1 != tt.wantCount {
				t.Errorf("getNodesAsList() got1 = %v, wantList %v", got1, tt.wantCount)
			}
		})
	}
}

func makeListNode(number int, size int) *ListNode {
	if size <= 1 {
		return &ListNode{
			Val:  number,
			Next: nil,
		}
	}

	return &ListNode{
		Val:  number,
		Next: makeListNode(number+1, size-1),
	}
}

func makeListOfNodes(number int, size int) []*ListNode {
	nodeList := make([]*ListNode, size)

	for i := 0; i < size; i++ {
		nodeList[i] = makeListNode(number+i, size-i)
	}

	return nodeList
}
