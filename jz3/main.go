package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	return process(head)
}
func process(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	res := process(head.Next)
	res = append(res, head.Val)
	return res
}
func main() {

}
