package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// https://leetcode.com/problems/merge-nodes-in-between-zeros/
func MergeNodes(head *ListNode) *ListNode {
	curr := head.Next
	prev := head
	for curr.Next != nil {
		if curr.Val == 0 {
			prev.Next = curr
			prev = curr
		}
		prev.Val += curr.Val
		curr = curr.Next
	}
	return head
}

func MergeNodesV1(head *ListNode) *ListNode {
	temp := head
	newHead := &ListNode{}
	writeHead := newHead
	for temp.Next != nil {
		val := 0
		node := temp.Next
		for node.Val != 0 {
			val += node.Val
			node = node.Next
		}
		writeHead.Next = &ListNode{
			Val: val,
		}
		writeHead = writeHead.Next
		temp.Next = node.Next
	}
	return newHead.Next
}
