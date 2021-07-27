package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param n int整型
 * @return int整型
 */
func NumberOf1( n int ) int {
	// write code here
	cnt:=0
	for n!=0{
		n=n&int(int32(n)-1)
		cnt++
	}
	return cnt
}


func main() {
	fmt.Println(NumberOf1(10))
}
