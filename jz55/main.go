package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//{1,2},{3,4,5}
//已上述链表为例，lows每次走一步，slow每次走两步，那么他们第一次相遇是在节点4
//然后重置lows到起始节点，两个指针各自都慢慢走,然后下次相遇就会在起点

func EntryNodeOfLoop(pHead *ListNode) *ListNode{
	lows:=pHead
	slow:=pHead
	number:=1
	for slow!=nil&&slow.Next!=nil{
		if number==1{
			slow=slow.Next.Next
			lows=lows.Next
		}else {
			slow=slow.Next
			lows=lows.Next
		}

		if lows.Val==slow.Val&&number!=0{
			lows=pHead
			number--
		}else if  lows.Val==slow.Val&&number==0{
			return lows
		}
	}
	return nil
}


func main() {
	
}
