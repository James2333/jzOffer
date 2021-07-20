package main


func Add( a int ,  b int ) int {
	// write code here
	for a != 0 {
		temp := a ^ b
		a = (a & b) << 1
		b = temp
	}
	return b
}

func main() {
	
}
