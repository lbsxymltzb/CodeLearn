package list

/*
描述
假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。
给定两个这种链表，请生成代表两个整数相加值的结果链表。

数据范围：0 \le n,m \le 10000000≤n,m≤1000000，链表任意值 0 \le val \le 90≤val≤9
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)

例如：链表 1 为 9->3->7，链表 2 为 6->3，最后生成新的结果链表为 1->0->0->0。

示例1
输入：
[9,3,7],[6,3]
复制
返回值：
{1,0,0,0}
复制
说明：
如题面解释
示例2
输入：
[0],[6,3]
复制
返回值：
{6,3}
*/

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head1 ListNode类
 * @param head2 ListNode类
 * @return ListNode类
 */
//func addInList( head1 *ListNode ,  head2 *ListNode ) *ListNode {
//	// write code here
//}

func AddInList(head1 *ListNode ,  head2 *ListNode) *ListNode {
	var s1 []*ListNode
	var s2 []*ListNode
	p1 := head1
	p2 := head2
	for {
		if p1 == nil && p2 == nil {
			break
		}
		if p1 != nil {
			s1 = append(s1, p1)
			p1 = p1.Next
		}
		if p2 != nil {
			s2 = append(s2, p2)
			p2 = p2.Next
		}
	}
	// add
	src := s1
	dst := s2
	if len(dst) < len(src) {
		t := src
		src = dst
		dst = t
	}
	pre := 0
	var top_dst *ListNode
	var top_src *ListNode
	for {
		if len(dst) > 0 {
			top_dst = dst[len(dst) - 1]
			dst = dst[:len(dst) - 1] //pop
		} else {
			break
		}
		num_add := 0
		if len(src) > 0 {
			top_src = src[len(src) - 1]
			src = src[:len(src) - 1] //pop
			num_add = top_src.Val
		} else {
			num_add = 0
		}
		add_rst :=top_dst.Val + num_add + pre
		if add_rst >= 10 {
			pre = 1
			top_dst.Val = add_rst - 10
		} else {
			pre = 0
			top_dst.Val = add_rst
		}
	}
	if pre > 0 {
		// new node
		n := &ListNode{Val: pre, Next: nil}
		n.Next = top_dst
		return n
	} else {
		return top_dst
	}

	return nil
}