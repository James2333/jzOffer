package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//用广度优先遍历，把左右节点以此添加到队列里面，然后再依次出队
//把出队的值记录下来。

func maxDepth(root *TreeNode)[]int{
	if root==nil{
		return make([]int,0)
	}
	queue:=make([]*TreeNode,0)
	queue=append(queue,root)
	depth:=make([]int,0)
	for len(queue)>0{
		size := len(queue)
		for i:=0;i<size;i++{
			v:=queue[0]
			depth=append(depth,v.Val)
			queue=queue[1:]

			if v.Left!=nil{
				queue=append(queue,v.Left)
			}
			if v.Right!=nil{
				queue=append(queue,v.Right)
			}
		}
	}
	return depth
}


func main() {
	
}
