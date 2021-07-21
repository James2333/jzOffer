package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @param n int整型
 * @return string字符串
 */
func LeftRotateString( str string ,  n int ) string {
	// write code here
	if len(str)<n{
		return ""
	}
	rstr:=[]rune(str)
	mid := rstr[0:n]
	rstr=rstr[n:]
	rstr=append(rstr,mid...)
	return string(rstr)
}

func main() {
	fmt.Println(LeftRotateString("abcdefg",3))
}
