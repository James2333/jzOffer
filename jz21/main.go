package main

import "fmt"

func IsPopOrder( pushV []int ,  popV []int ) bool {
	// write code here
	st := make([]int,0)
	i:=0
	var l int
	for _,v:=range pushV{
		st=append(st,v)
		l=len(st)
		for popV[i]==st[l-1]{
			//出栈
			st=st[:len(st)-1]
			l=len(st)
			i++
			if i==len(popV)||l==0{
				break
			}
		}
	}
	//for i<len(popV){
	//	if st[len(st)-1]==popV[i] {
	//		st=st[:len(st)-1]
	//		i++
	//		continue
	//	}
	//	i++
	//}
	return len(st)==0
}


func main() {
	pushV:=[]int{1,2,3,4,5}
	popV:=[]int{4,5,3,2,1}
	fmt.Println(IsPopOrder(pushV,popV))
}
