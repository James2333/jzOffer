package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeDepth( pRoot *TreeNode ) int {
	// write code here
	if pRoot==nil{
		return 0
	}
	return max(TreeDepth(pRoot.Left),TreeDepth(pRoot.Right))+1
}



func max(l,r int) int {
	if l>r{return l}
	return r
}
func main() {
	
}
