package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	l := &ListNode{}
	l.Next = pHead
	r:=pHead

	for i:=0;i<k;i++{
		//先判断链表长度是否小于k
		if r==nil{
			return nil
		}
		r=r.Next
	}
	for r!=nil{
		r=r.Next
		l=l.Next
	}
	return l.Next
}

func main() {

}
