package main

import "fmt"

func f(n,m int)int  {
	if (n == 1) {return 0}
	x := f(n-1, m)
	return (x+m) % n
}
func LastRemaining_Solution( n int ,  m int ) int {
	// write code here
	if n==0{
		return -1
	}
	return f(n,m)
	//nums:=make([]bool,n)
	//mid:=0
	//for  {
	//	for k,v:=range nums{
	//		if v==false {
	//			if mid == m-1 {
	//				nums[k] = true
	//				mid = 0
	//				n--
	//				if n == 0 {
	//					return k
	//				}
	//			} else {
	//				mid++
	//				continue
	//			}
	//		}
	//	}
	//}
	//return -1
}


func main() {
	fmt.Println(LastRemaining_Solution(5,3))
}
