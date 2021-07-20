package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	//把head1短的遍历完之后遍历另一个，长的遍历完遍历短的，最终会在交汇点相遇。
	p1, p2 := pHead1, pHead2
	if p1 == nil || p2 == nil {
		return nil
	}
	for p1 != p2 {
		if p1 == nil {
			p1 = pHead2
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = pHead1
		} else {
			p2 = p2.Next
		}
	}
	return p2
}
func main() {

}
