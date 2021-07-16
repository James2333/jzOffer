package main



func MoreThanHalfNum_Solution( numbers []int ) int {
	// write code here
	//设定一个中间值，如果等于这个值则count++，否贼count--
	count:=0
	number:=numbers[0]
	for _,v:=range numbers{
		if number==v{
			count++
		}else {
			count--
		}
		if count<0 {
			count=1
			number=v
		}
	}
	return number
}
func main() {
	
}
