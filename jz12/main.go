package main


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param base double浮点型
 * @param exponent int整型
 * @return double浮点型
 */
func Power( base float64 ,  exponent int ) float64 {
	if exponent<0{
		base=1/base
		exponent*=-1
	}
	result,x:=1.0,base
	for ; exponent!=0; exponent >>=1 {
		if exponent & 0x01 !=0 {
			result= result*x
		}
		//         x*=base
		x*=x
	}
	return result

}


func main() {
	
}
