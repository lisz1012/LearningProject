package main

import "fmt"

func main() {
	//var score int = 11
	switch 5 > 10 {
	case true:
		fmt.Println("hahaha")
	case false:
		fmt.Println("heiheihei")
	}

	//var a int8 = 9
	//switch a {
	//case 128:
	//	fmt.Println("a")
	//}
	var a int8 = 12
	switch { // 怪异!Q！
	case a > 10:
		fmt.Println("A")
	case a > 11:
		fmt.Println("B") // 不会执行到，go case会自动省略break。
	}

	switch b := 12; { // ;不能省略，也感觉略怪异！！
	case b > 6:
		fmt.Println("b > 6")
		fallthrough // 只能穿透一层case！go特有的规定！
	case b > 0:
		fmt.Println("b > 0") // 不会执行到，除非上面有fallthrough，且只会穿透一层
	}
}
