package main

import "fmt"

func minNumberInRotateArray( rotateArray []int ) int {
	if len(rotateArray)==0{
		return 0
	}
	minVal := rotateArray[0]
	return getMinValue(minVal, rotateArray[1:])

}

func getMinValue(val int, array []int)int{
	if len(array)==0{
		return val
	}
	index:=len(array)/2
	middleVal := array[index]
	if val<middleVal{
		return getMinValue(val, array[index+1:])
	}
	return getMinValue(middleVal, array[:index])

}
func binnarySearch(array []int,search int,l, r int) int {
	if l > r {
		return -1
	}
	mid :=(l+r)/2
	if search==array[mid]{
		return mid
	}
	if array[mid]<search{
		return binnarySearch(array,search,mid+1,r)
	}
	if array[mid]>search{
		return binnarySearch(array,search,l,mid-1)
	}
	return -1
}
func search(nums []int, target int) int {
	return binnarySearch(nums, target, 0, len(nums))
}
func main() {
	var sortedArray []int = []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	fmt.Println(minNumberInRotateArray(sortedArray))
	fmt.Println(search(sortedArray,11))
}
