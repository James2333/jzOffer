package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func KthNode( pRoot *TreeNode ,  k int ) *TreeNode {
	// write code here
	//return Pmid(pRoot,k,i)
	//var stack []*TreeNode
	//for pRoot !=nil||len(stack)>0{
	//	for pRoot!=nil{
	//		stack=append(stack,pRoot)
	//		pRoot=pRoot.Left
	//	}
	//	pRoot=stack[len(stack)-1]
	//	stack=stack[:len(stack)-1]
	//	k--
	//	if k==0{
	//		return pRoot
	//	}
	//	pRoot=pRoot.Right
	//}
	//return nil
	//res:=dfs(pRoot,k)
	var res *TreeNode
	var f func(root *TreeNode) *TreeNode
	f= func(root *TreeNode)*TreeNode {

		if root==nil{
			return nil
		}
		f(root.Left)
		if k==0{
			return res
		}
		k--
		if k==0{
			res=root
		}
		f(root.Right)
		return res
	}
	f(pRoot)
	return res
}
//leetcode
func kthLargest(root *TreeNode, k int) int {
	var res int
	var f func(root *TreeNode) int
	f= func(root *TreeNode) int {
		if root==nil{
			return res
		}
		f(root.Right)
		if k==0{
			return res
		}
		k--
		if k==0{
			res=root.Val
		}

		f(root.Left)
		return res
	}
	f(root)
	return res
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
