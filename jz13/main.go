package main

import (
	"fmt"
	"log"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型一维数组
 * @return int整型一维数组
 */
//func reOrderArray( array []int ) []int {
//	// write code here
//	length:=len(array)
//	l:=0
//	r:=0
//	//l先遍历找到偶数，然后r再从l的地方开始遍历，找到奇数，然后调换位置。直到l遍历完。
//	for ;l<length;l++{
//		if array[l]%2==0{
//			for r=l;r<length;r++{
//				if array[r]%2==1{
//					exchange(array,l,r)
//					break
//				}
//			}
//		}
//	}
//	return array
//}

func exchange(nums []int) []int {
	length:=len(nums)
	l:=0
	r:=0
	//l先遍历找到偶数，然后r再从l的地方开始遍历，找到奇数，然后调换位置。直到l遍历完。
	for ;l<length;l++{
		if nums[l]%2==0{
			for r=l;r<length;r++{
				if nums[r]%2==1{
					nums[l],nums[r]=nums[r],nums[l]
					break
				}
			}
		}
	}
	return nums
}

func reOrderArray( array []int ) []int {
	// write code here
	l:=make([]int,0)
	r:=make([]int,0)
	for _,v:=range array{
		if v%2==0{
			r=append(r,v)
		}else {
			l=append(l,v)
		}
	}
	l=append(l,r...)
	return l
}

func exChange(array []int,l,r int)  {
	array[l],array[r]=array[r],array[l]
	fmt.Println(array)
}
func main() {
	arr:=[]int{2,4,6,5,7}
	zbc:=reOrderArray(arr)
	log.Println(zbc)
}
