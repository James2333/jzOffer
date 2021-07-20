package main


func IsContinuous( numbers []int ) bool {
	// write code here
	max,min:=0,14
	set:=make(map[int]int)
	for _,v:=range numbers{
		if v==0{
			continue
		}
		if _,ok:=set[v];ok{
			return false
		}else {
			set[v]=1
		}
		if v>max{
			max=v
		}
		if v<min{
			min=v
		}
	}
	return max-min<5&&max-min>=0
}
func main() {
	
}
