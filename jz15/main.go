package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	//1）nex = cur->next, 保存作用
	//2）cur->next = pre 未反转链表的第一个节点的下个指针指向已反转链表的最后一个节点
	//3）pre = cur， cur = nex; 指针后移，操作下一个未反转链表的第一个节点
	var pre *ListNode
	cur := new(ListNode)
	nex := new(ListNode)
	cur = pHead
	for cur != nil {
		nex = cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}
func main() {

}
