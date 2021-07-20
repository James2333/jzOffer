package main

import "fmt"

func StrToInt( str string ) int {
	// write code here
	if len(str)==0||str=="0"{
		return 0
	}
	hash:=make(map[int32]int)
	rstr:=[]rune(str)
	//fmt.Println(rstr)
	a:=0
	num:=0
	hash[43]=0
	hash[45]=-1
	for i:=int32(48);i<=57;i++ {
		hash[i]=a
			a++
	}
	for _,v:=range rstr{
		if value,ok:=hash[v];ok{
			if v==43||v==45{
				continue
			}
				num=num*10+value
		}else {
			return 0
		}
	}
	if rstr[0]==45{num*=-1}

	return num
}

func main() {
	zbc:=StrToInt("-123")
	fmt.Println(zbc)
}
