package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param sum int整型
 * @return int整型二维数组
 */
func FindContinuousSequence( sum int ) [][]int {
	// write code here
	l:=1
	r:=2
	s:=3
	sums:=make([][]int,0)
	for l<=sum/2{
		if s==sum{
			sums=append(sums,NewContainer(l,r))
			s-=l
			l++
		}else if s>=sum{
			s-=l
			l++
		}else{
			r++
			s+=r
		}

	}
	return sums
}
func NewContainer(l,r int)  []int{
	sum:=make([]int,0)
	for ;l<=r;l++{
		sum=append(sum,l)
	}
	return sum
}
func main() {
	fmt.Println(FindContinuousSequence(9))
}
