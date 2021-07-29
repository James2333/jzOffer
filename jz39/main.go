package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	if pRoot==nil{
		return true
	}
	return abs(getMaxDepth(pRoot.Left)-getMaxDepth(pRoot.Right))<=1
}

func getMaxDepth(pRoot *TreeNode) int {
	var deepth int
	if pRoot == nil {
		deepth = 0
		return deepth
	}
	templ, tempr := getMaxDepth(pRoot.Left), getMaxDepth(pRoot.Right)
	return 1 + max(templ, tempr)
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func main() {

}
