package main

import "fmt"

func findKthLargest(nums []int, k int) int {

	return QuickSort(nums,0,len(nums)-1,k)
}
func QuickSort(nums []int,l,r,k int) int{
	if l>=r{
		return 0
	}
	i,j:=l,r
	for i<j{
		for i<j&&nums[l] <=nums[j]{
			j--
		}
		for i<j&& nums[l] >= nums[i]{
			i++
		}
		nums[i],nums[j]=nums[j],nums[i]
		fmt.Println(i,j)
	}
	nums[j],nums[l]=nums[l],nums[j]
	if len(nums)-k==i{
		return nums[len(nums)-k]
	}
	QuickSort(nums,l,i-1,k)
	QuickSort(nums,i+1,r,k)
	return 0
}

func main() {
	zbc:=[]int{3,2,3,1,2,4,5,5,6}
	fmt.Println(findKthLargest(zbc,4))
}
