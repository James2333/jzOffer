package main

func FirstNotRepeatingChar( str string ) int {
	// write code here
	hashT:=[256]int{}
	runeStr:=[]rune(str)
	for k,_:=range runeStr{
		hashT[runeStr[k]]++
	}
	for j := 0; j < len(str); j++ {
		if hashT[str[j]] == 1 {
			return j
		}
	}
	return -1
}

func main() {
	
}
