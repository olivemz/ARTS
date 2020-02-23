package main

import "fmt"

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0 // length of linked list
	dumyPointer := &ListNode{0, head}
	for dumyPointer.Next != nil {
		i++
		dumyPointer = dumyPointer.Next
	}

	if i == 1 {
		return nil
	}
	modifiedNode := i - n
	fmt.Println("modify cation node", modifiedNode)

	modified := head
	prev := head
	j := 0
	for modified != nil {
		if j == modifiedNode {
			fmt.Println("matched current value", modified.Val)
			if j == 0 {
				return head.Next
			} else {
				prev.Next = modified.Next
			}
		}
		prev = modified
		modified = modified.Next
		j++
	}
	return head
}
func main() {

	second := &ListNode{2, nil}
	head := &ListNode{1, second}
	removeNthFromEnd(head, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// TBC...
