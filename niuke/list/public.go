package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Create(in []int) (*ListNode, *ListNode) {
	if len(in) < 1 {
		return nil, nil
	}
	var h *ListNode
	var p *ListNode

	for _, v := range in {
		n := &ListNode{Val: v, Next: nil}
		if h == nil {
			h = n
			p = n
			continue
		}
		p.Next = n
		p = n
	}
	return h, p
}

func (l *ListNode) Show() {
	if l == nil {
		return
	}
	p := l
	for {
		if p.Next == nil {
			// last node
			fmt.Printf("%d ", p.Val)
			break
		}
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Printf("\n")
}
