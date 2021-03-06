package main


type ListNode struct{
   Val int
   Next *ListNode
 }


/**
 *
 * @param pHead1 ListNodeš▒╗
 * @param pHead2 ListNodeš▒╗
 * @return ListNodeš▒╗
 */
func Merge( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	if pHead1.Val>pHead2.Val{
		pHead2.Next=Merge(pHead1,pHead2.Next)
		return pHead2
	}else {
		pHead1.Next=Merge(pHead1.Next,pHead2)
		return pHead1
	}
}

func main() {
	
}
