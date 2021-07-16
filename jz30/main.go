package main

import "math"

func FindGreatestSumOfSubArray( array []int ) int {
	// write code here
	//去从这个数组头开始遍历，相加，如果array[i-1]<0则直接舍弃掉
	//
	mid:=array[0]
	for k,_ :=range array{
		if k==0{
			continue
		}
		array[k]+=int(math.Max(float64(array[k-1]),float64(0)))
		mid=int(math.Max(float64(array[k]),float64(mid)))
	}
	return mid
}
func main() {
	
}
