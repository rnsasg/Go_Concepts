package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	p, m := middleNode(head)
	p.Next = reverse(m)
	middle := p.Next
	for middle != nil {
		if head.Val != middle.Val {
			return false
		}
		middle = middle.Next
		head = head.Next
	}
	return true
}

func middleNode(head *ListNode) (*ListNode, *ListNode) {

	if head == nil || head.Next == nil {
		return nil, head
	}

	var slow, fast *ListNode = head, head
	var prev *ListNode = nil

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		}
	}

	cur := head
	len := 0
	for cur != nil {
		len++
		cur = cur.Next
	}
	if len%2 == 1 {
		return slow, slow.Next
	}
	return prev, slow
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode = nil

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func main() {
	// var t1, t2, t3 *ListNode = nil, nil, nil
	t3 := &ListNode{Val: 1, Next: nil}
	t2 := &ListNode{Val: 0, Next: t3}
	t1 := &ListNode{Val: 1, Next: t2}
	head := t1

	p, m := middleNode(head)
	fmt.Print(p.Val)
	fmt.Print(m.Val)
	//isPalindrome(head)
}
