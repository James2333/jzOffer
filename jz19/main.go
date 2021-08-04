package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param matrix int整型二维数组
 * @return int整型一维数组
 */
var sum = make([]int, 0)

func printMatrix(matrix [][]int) []int {
	// write code here

	length := len(matrix)
	if length==1{
		return matrix[0]
	}
	i,j:=0,length
	for i<j{
		Mat(matrix, j, i)
		i++
		j--
	}

	return sum
}
func Mat(matrix [][]int, length, size int) {
	i, j := size, size
	for ; j < length; j++ {
		sum = append(sum, matrix[i][j])
	}
	i += 1
	j -= 1
	for ; i < length; i++ {
		sum = append(sum, matrix[i][j])
	}
	i -= 1
	j -= 1
	for ; j >= size; j-- {
		sum = append(sum, matrix[i][j])
	}
	i -= 1
	j += 1
	for ; i > size; i-- {
		sum = append(sum, matrix[i][j])
	}
	i += 1
}

func main() {
	zbc := [][]int{{1,2,3,4,5},{6,7,8,9,10},{11,12,13,14,15},{16,17,18,19,20},{21,22,23,24,25}}
	fmt.Println(printMatrix(zbc))
}
