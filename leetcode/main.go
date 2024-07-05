package main

import "linkedlist"

func main() {
	head := linkedlist.ListNode{
		Val: 0,
		Next: &linkedlist.ListNode{
			Val: 3,
			Next: &linkedlist.ListNode{
				Val: 1,
				Next: &linkedlist.ListNode{
					Val: 0,
					Next: &linkedlist.ListNode{
						Val: 4,
						Next: &linkedlist.ListNode{
							Val: 5,
							Next: &linkedlist.ListNode{
								Val: 2,
								Next: &linkedlist.ListNode{
									Val:  0,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	linkedlist.MergeNodes(&head)
}

// do it the other way
