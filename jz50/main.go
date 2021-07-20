package main



func duplicate( numbers []int ) int {
	// write code here
	if len(numbers)==0{
		return -1
	}
	hash:=make(map[int]struct{})
	for i:=0;i<len(numbers);i++{
		if _,ok:=hash[numbers[i]];ok{
			return numbers[i]
		}else {hash[numbers[i]]= struct{}{}}
	}
	return -1
}


func main() {
	
}
