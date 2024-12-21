package main

import "fmt"

// ListNode definition
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd removes the n-th node from the end of the list.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a dummy node to handle edge cases like removing the head
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// Move the fast pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// Move both fast and slow pointers until fast reaches the end
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the n-th node from the end
	slow.Next = slow.Next.Next

	// Return the head of the updated list
	return dummy.Next
}

// Helper function to print the list
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

// Helper function to create a list from a slice
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func main() {
	// Example usage:
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Println("Original list:")
	printList(head)

	// Remove the 2nd node from the end (node with value 4)
	head = removeNthFromEnd(head, 2)

	// Print the updated list
	fmt.Println("Updated list after removal:")
	printList(head)
}
