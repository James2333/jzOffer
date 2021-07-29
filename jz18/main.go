package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pRoot TreeNode类
 * @return TreeNode类
 */
func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot==nil{
		return pRoot
	}
	updown(pRoot)
	return pRoot
}

func updown(p *TreeNode){
	if p==nil{
		return
	}
	p.Left,p.Right=p.Right,p.Left
	updown(p.Left)
	updown(p.Right)
}
func main() {

}
