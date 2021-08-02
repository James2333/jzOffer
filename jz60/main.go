package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode)[][]int{
	if root==nil{
		return make([][]int,0)
	}
	queue:=make([]*TreeNode,0)
	queue=append(queue,root)
	depth:=make([][]int,0)
	for len(queue)>0{
		size := len(queue)
		sum:=make([]int,0)
		for i:=0;i<size;i++{
			v:=queue[0]
			sum=append(sum,v.Val)
			queue=queue[1:]

			if v.Left!=nil{
				queue=append(queue,v.Left)
			}
			if v.Right!=nil{
				queue=append(queue,v.Right)
			}
		}
		depth=append(depth,sum)
	}
	return depth
}


func main() {
	
}
