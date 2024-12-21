package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev := (*ListNode)(nil)
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	first, second := head, prev
	for second.Next != nil {
		temp1 := first.Next
		temp2 := second.Next
		first.Next = second
		second.Next = temp1
		first = temp1
		second = temp2
	}

}
