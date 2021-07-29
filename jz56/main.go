package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplication(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil || pHead.Next == nil{
		return pHead
	}
	p := pHead
	if p.Val!=p.Next.Val{
		p.Next=deleteDuplication(pHead.Next)
		return p
	}else {
		mid:=p.Val
		cur:=p.Next
		for cur.Next!=nil&&cur.Val==mid{
			cur=cur.Next
		}
		if cur==nil{
			return nil
		}
		return deleteDuplication(cur)
	}
}

func main() {

}
