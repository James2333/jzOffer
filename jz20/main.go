package main

type stack []int

var stack1 []int = []int{}
var stack2 []int = []int{}

func Push(node int) {
	// write code here
	stack1 = append(stack1, node)
	if len(stack2) == 0 || node <= stack2[len(stack2)-1] {
		stack2 = append(stack2, node)
	} else {
		stack2 = append(stack2, stack2[len(stack2)-1])
	}
}
func Pop() {
	// write code here
	stack1 = stack1[:len(stack1)-1]
	stack2 = stack2[:len(stack2)-1]
}
func Top() int {
	// write code here
	return stack1[len(stack1)-1]
}
func Min() int {
	// write code here
	return stack2[len(stack2)-1]
}
func main() {

}
