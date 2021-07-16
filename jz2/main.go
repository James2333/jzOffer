package main

import "fmt"

func replaceSpace( s string ) string {
	// write code here
	runes:=make([]rune,0)
	//fmt.Println(runes)
	//return ""
	for _,st:=range s{
		if st==' '{
			runes=append(runes,'%','2','0')
		}else {
			runes=append(runes,st)
		}
	}
	return string(runes)
}
func main() {
	fmt.Println(replaceSpace("zbc  zbb 卧槽"))
}
