package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
var i int
func KthNode( pRoot *TreeNode ,  k int ) *TreeNode {
	// write code here
	//return Pmid(pRoot,k,i)
	var stack []*TreeNode
	for pRoot

}

func Pmid(p *TreeNode,k,i int) *TreeNode {
	if p==nil{
		return p
	}
	if k==i{
		return p
	}
	i++
	Pmid(p.Left,k,i)
	//fmt.Println(p.Val)
	Pmid(p.Right,k,i)
	return nil
}
func main() {
	
}
