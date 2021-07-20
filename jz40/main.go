package main


func FindNumsAppearOnce( array []int ) []int {
	// write code here
	nums:=make([]int,0)
	hash:=make(map[int]int)
	for _,v:=range array{
		if _,ok:=hash[v];ok{
			hash[v]++
		}else{
			hash[v]=1
		}
	}
	for k,v:=range hash{
		if v!=1{
			continue
		}else {
			nums=append(nums,k)
		}
	}
	if nums[0]>nums[1]{
		nums[0],nums[1]=nums[1],nums[0]
	}
	return nums
}

func main() {
	
}
