package problem_21

type LinkedList struct {
	head   *ListNode
	length int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list := LinkedList{
		head:   nil,
		length: 0,
	}

	var node1 = list1
	var node2 = list2

	for {
		if node1 == nil && node2 == nil {
			break
		}

		if node1 != nil && node2 == nil {
			list = appendNodes(list, node1)

			break
		}

		if node1 == nil && node2 != nil {
			list = appendNodes(list, node2)

			break
		}

		if node1.Val <= node2.Val {
			list = appendNode(list, node1.Val)
			node1 = node1.Next
		} else {
			list = appendNode(list, node2.Val)
			node2 = node2.Next
		}
	}

	return list.head
}

func appendNodes(list LinkedList, listNode *ListNode) LinkedList {
	node := listNode

	for {
		if node == nil {
			break
		}

		list = appendNode(list, node.Val)
		node = node.Next
	}

	return list
}

func appendNode(list LinkedList, value int) LinkedList {
	node := &ListNode{
		Val:  value,
		Next: nil,
	}

	if list.head == nil {
		list.head = node
	} else {
		listNode := list.head

		for listNode.Next != nil {
			listNode = listNode.Next
		}

		listNode.Next = node
	}
	list.length += 1

	return list
}
