package main

import "fmt"

func Fibonacci( n int ) int {
	// write code here
	a:=1
	b:=1
	if n<2{
		return n
	}else {
		for n>2{
			a,b=b,a+b
			n--
		}
	}
	return b
}

func main() {
	fmt.Println(Fibonacci(2))
}
