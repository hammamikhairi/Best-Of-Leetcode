package main

// my solution for https://leetcode.com/problems/merge-k-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func MinNdPos(list *[]*ListNode) (*ListNode, int) {
	var min *ListNode = (*list)[0]
	var minIndex int

	for ind, node := range *list {
		if node.Val < min.Val {
			min = node
			minIndex = ind
		}
	}
	return min, minIndex
}

func mergeKLists(list []*ListNode) *ListNode {

	head := &ListNode{0, nil}
	tmp := head

	for true {
		if len(list) == 0 {
			break
		}

		curr, pos := MinNdPos(&list)

		if list[pos].Next == nil {
			list = append(list[:pos], list[pos+1:]...)
		} else {
			list[pos] = list[pos].Next
		}
		tmp.Next = curr
		tmp = tmp.Next
	}
	tmp.Next = nil

	return head.Next
}
