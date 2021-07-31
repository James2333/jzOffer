package main

import "fmt"

/**
 *
 * @param data int整型一维数组
 * @param k int整型
 * @return int整型
 */
func GetNumberOfK( data []int ,  k int ) int {
	// write code here
	mid:=search(data,k)
	if mid==-1{
		return 0
	}
	l,r:=mid-1,mid+1
	sum:=1
		for l>=0&&data[l]==k{
			sum++
			l--
		}
		for r<=len(data)-1&&data[r]==k{
			sum++
			r++
		}

	return sum
}

func search(num []int,k int)int{
	if len(num) < 0 {
		return -1
	}
	l,r:=0,len(num)-1
	for l<=r{
		mid:=(l+r)/2
		if num[mid]==k{
			return mid
		}
		if num[mid]>k{
			r=mid-1
		}
		if num[mid]<k{
			l=mid+1
		}
	}
	return -1
}

func main() {
	zbc:=[]int{1,3,3,3,3,4,5}
	fmt.Println(GetNumberOfK(zbc,2))
}
