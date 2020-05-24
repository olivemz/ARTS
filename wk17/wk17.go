// Algorythm

//25
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dumy, fakeHead := &ListNode{}, &ListNode{}
    dumy = head
    fakeHead.Next = head
    prev := fakeHead
    i := 0
    for dumy != nil {
        i = i + 1
        if (i % k == 0) {
            prev = reverseHelper(prev, dumy.Next)
            dumy = prev.Next
        } else {
            dumy = dumy.Next
        }
    }
    return fakeHead.Next
}

func reverseHelper (prev *ListNode, next *ListNode) *ListNode { 
    last := prev.Next
    cur := last.Next
    for cur != next {
        last.Next = cur.Next
        cur.Next = prev.Next
        prev.Next = cur
        cur = last.Next
    }
    return last
}

// 61,
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	i := -1
	dummy, fakeHead, pre, last := &ListNode{}, &ListNode{}, &ListNode{}, &ListNode{}
	dummy.Next = head
	fakeHead.Next = head

	for dummy != nil {
		dummy = dummy.Next
		i = i + 1
	}

	if head == nil || head.Next == nil {
		return head
	}

	rotate := k % i
	location := 0
	if rotate > 0 {
		location = i - rotate
	}

	if location == 0 {
		return head
	}

	j := -1
	for fakeHead != nil {
		if j == location-1 {
			pre = fakeHead
		}
		if j == i-1 {
			last = fakeHead
		}
		fakeHead = fakeHead.Next
		j = j + 1
	}

	newHead := pre.Next
	pre.Next = nil
	last.Next = head

	return newHead
}


// 82

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    temp, pre, newHead := &ListNode{}, &ListNode{} , &ListNode{}
    temp.Next = head
    newHead.Next = temp
    pre = temp
    
    if head == nil || head.Next == nil {
        return head
    }
    
    for head != nil {
        pre, head = removeDupe(pre, head)
    }
    
    return newHead.Next.Next
}

func removeDupe(pre *ListNode, cur *ListNode) (newPre *ListNode, newCur *ListNode){
    i := 0
    for cur != nil && cur.Next != nil && cur.Next.Val == cur.Val {
        cur = cur.Next
        i = i + 1
    }
    
    if i == 0 {
        pre = cur
    }
    
    if cur.Next != nil {
        cur = cur.Next
    } else {
        cur = nil
    }
    
     if i > 0 {
        pre.Next = cur
     }
    
    return pre, cur
}


// 83

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    newHead := head
    
    if head == nil || head.Next == nil {
        return head
    }
    
    for head != nil {
        dup := false
        if head != nil && head.Next != nil && head.Val == head.Next.Val {
            head.Next = head.Next.Next
            dup = true
        }
        if dup == false {
             head = head.Next
        }
    }
    
    return newHead
}