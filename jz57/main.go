package main

import "fmt"

type TreeLinkNode struct {
    Val int
    Left *TreeLinkNode
    Right *TreeLinkNode
    Next *TreeLinkNode
}


func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	Pmid(pNode)
	return pNode
}

func Pmid(p *TreeLinkNode)  {
	if p==nil{
		return
	}

	Pmid(p.Left)
	fmt.Println(p.Val)
	Pmid(p.Right)
}
func main() {
	
}
