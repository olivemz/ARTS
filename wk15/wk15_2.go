package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := &ListNode{}
	var returnVal = newList
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Next = l1
			l1 = l1.Next
		} else {
			newList.Next = l2
			l2 = l2.Next
		}
		newList = newList.Next
	}
	if l1 != nil {
		newList.Next = l1
	} else if l2 != nil {
		newList.Next = l2
	}

	return returnVal.Next
}
