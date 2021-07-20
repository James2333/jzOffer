package main

import "fmt"

var s int
func solution(n int)bool{
	s+=n
	return n>0&&solution(n-1)
}
func Sum_Solution( n int ) int {
	// write code here
	solution(n)
	return s
}

func main() {
	fmt.Println(Sum_Solution(5))
}
