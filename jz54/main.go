package main

import "fmt"

var hash =make(map[byte]int)
var stack []byte
func Insert(ch byte) {
	//把输入的字节流先判断是否在hash里面，不在则添加到hash里面去，在则
	if _,ok:=hash[ch];ok{
		hash[ch]++
	}else {
		hash[ch]=1
		stack=append(stack,ch)
	}
}

func FirstAppearingOnce() byte {
	for _,v:=range stack{
		if hash[v]==1{
			return v
		}
	}
	return '#'
}


func firstUniqChar(s string) byte {
	Hash:=make(map[byte]int)
	var Stack []byte
	str:=[]byte(s)
	for _,v:=range str{
		if _, ok := Hash[v];ok {
			Hash[v]++
			//Stack=append(Stack,v)
		} else {
			Hash[v]=1
			Stack=append(Stack,v)
		}
	}
	for _,v:=range Stack{
		if Hash[v]==1{
			return v
		}
	}
	return ' '
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
