package main

func Find( target int ,  array [][]int ) bool {
	// write code here
	if len(array)==0{return false}
	for i:=0;i<len(array);i++{
		l:=0
		r:=len(array[i])-1
		for l<=r{
			mid:=l+(r-l)/2
			if array[i][mid]<target{
				l=mid+1
			}else if array[i][mid]>target{
				r=mid-1
			}else {
				return true
			}
		}
		//for j:=0;j<len(array[i]);j++{
		//	if array[i][j]==target{return true}
		//}
	}
	return false
}
