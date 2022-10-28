package problem_876

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	nodeList, nodesCount := getNodesAsList(head)

	if nodesCount%2 != 0 {
		nodesCount--
	}

	targetCount := nodesCount/2 + 1

	return nodeList[targetCount-1]
}

func middleNodeByPointer(head *ListNode) *ListNode {
	middle := head
	end := head

	for end != nil && end.Next != nil {
		middle = middle.Next
		end = end.Next.Next
	}

	return middle
}

func getNodesAsList(head *ListNode) ([]*ListNode, int) {
	var nodeList []*ListNode
	listNodesCount := 1

	node := head

	for {
		nodeList = append(nodeList, node)

		nextNode := node.Next
		if nextNode == nil {
			break
		}

		node = nextNode
		listNodesCount++
	}

	return nodeList, listNodesCount
}
