package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param array int整型一维数组
 * @param sum int整型
 * @return int整型一维数组
 */
func FindNumbersWithSum( array []int ,  sum int ) []int {
	l:=0
	r:=len(array)-1
	for r>l{
		s:=array[l]+array[r]
		if s<sum{
			l++
		}else if s>sum{
			r--
		}else {
			return []int{array[l],array[r]}
		}
	}
	return nil
}

func main() {
	zbc:=[]int{10,18,25,33,36,50,50,52,57,74}
	fmt.Println(FindNumbersWithSum(zbc,126))
}
