package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param A int整型一维数组
 * @return int整型一维数组
 */
func multiply( A []int ) []int {
	// write code here
	if len(A)==0{
		return nil
	}
	b:=make([]int,len(A))
	b[0]=1
	for i:=1;i<len(A);i++{
		b[i]=b[i-1]*A[i-1]
	}
	tmp:=1
	for j:=len(A)-2;j>=0;j--{
		tmp*=A[j+1]
		b[j]*=tmp
	}
	return b
}

func main() {
	zbc:=[]int{1,2,3,4,5}
	fmt.Println(multiply(zbc))
}
