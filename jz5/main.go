package main
var stack1 [] int
var stack2 [] int

func Push(node int) {
	stack1=append(stack1,node)
}

func Pop() int{
	if len(stack1)==0{
		return 0
	}
	defer func() {
		stack1=stack1[1:]
	}()
	return stack1[0]
}
func main() {
	
}
