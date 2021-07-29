package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */
func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here

	//l := 0
	r := len(input) - 1
	QuickSort(input,0,r)
	return input[0:k]
}

func QuickSort(input []int,l,r int)  {
	if l>=r{
		return
	}
	i, j := l, r
	for i < j {
		for i < j && input[j] >= input[l] {
			j--
		}
		for i < j && input[i] <= input[l] {
			i++
		}
		input[i], input[j] = input[j], input[i]
	}
	input[l], input[i] = input[i], input[l]
	QuickSort(input,l,i-1)
	QuickSort(input,i+1,r)
}
func main() {
	zbc:=[]int{4,5,1,6,2,7,3,8}
	fmt.Println(GetLeastNumbers_Solution(zbc,4))
}
