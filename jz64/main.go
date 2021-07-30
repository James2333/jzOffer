package main

/**
 *
 * @param num int整型一维数组
 * @param size int整型
 * @return int整型一维数组
 */
func maxInWindows( num []int ,  size int ) []int {
	// write code here
	if size==0{
		return nil
	}
	sum:=make([]int,0)
	var f func(num []int)
	f= func(num []int) {
		zbc:=0
		for _,v:=range num{
			if v>zbc{
				zbc=v
			}
		}
		sum=append(sum,zbc)
	}
	l:=0
	r:=size
	for r<=len(num){
		f(num[l:r])
		l++
		r++
	}
	return sum
}


func main() {
	
}
