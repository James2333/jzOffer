package main

import (
	"fmt"
	"strings"
)

func ReverseSentence( str string ) string {
	// write code here
	str = strings.TrimSpace(str)
	length:=len(str)
	l:=length-2
	r:=length-1
	strs:=""
	for l>=0{
		if string(str[l])==" "&&string(str[l+1])!=" "{
			strs=strs+" "+str[l+1:r+1]
			r=l
			l--
		}else if string(str[l])!=" "&&string(str[l+1])==" "{
			r=l
			l--
		}else {
			l--
		}
	}
	strs=strs+" "+str[0:r+1]
	strs=strs[1:]
	return strs
}

//func ReverseStr(str string)string  {
//	mid:=""
//	for i:=len(str)-1;i>=0;i--{
//		mid+=string(str[i])
//	}
//	return mid
//}

func main() {
	fmt.Println(ReverseSentence("  hello     world!  "))
	//fmt.Println(ReverseStr("zbc"))
}
