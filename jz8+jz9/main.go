package main

import "fmt"

func jumpFloor( number int ) int {
	// write code here
	//a,b:=1,1
	//for i:=1;i<number;i++{
	//	a,b=b,a+b
	//}
	//return b
	return jump(number)
}
//f(n)=2f(n-1)
//f(n)=f(1)<<n

func jump(number int) int{
	//sum:=0
	if number<1{
		return number
	} else {
		return 1<<(number-1)
	}
	//return sum
	return -1
}
func main() {
	fmt.Println(jump(3))
}
