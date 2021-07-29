package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pre int整型一维数组
 * @param vin int整型一维数组
 * @return TreeNode类
 */
func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
	// write code here
	if len(pre)==0{
		return nil
	}
	root:=&TreeNode{pre[0],nil,nil}
	i:=0
	for ;i<len(vin);i++{
		if vin[i]==pre[0]{
			break
		}
	}
	//遍历找到根节点   及per[0]==vin[i]
	//然后分别递归 ，传入前序遍历和中序遍历
	root.Left=reConstructBinaryTree(pre[1:len(vin[:i])+1],vin[:i])
	root.Right=reConstructBinaryTree(pre[len(vin[:i])+1:],vin[i+1:])
	return root
}



func main() {
	
}
