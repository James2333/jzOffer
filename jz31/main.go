package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param n int整型
 * @return int整型
 */
func NumberOf1Between1AndN_Solution(n int) int {
	// write code here
	low := 0
	cur := n % 10
	high := n / 10
	digit := 1
	res := 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += high*digit + digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}

func main() {
	fmt.Println(NumberOf1Between1AndN_Solution(13))
}
