package main

import "fmt"

func main() {
	var n1 int = 10
	fmt.Println(n1)
	var n2 int = 3 + 12
	fmt.Println(n2)
	var s1 string = "abc" + "def"
	fmt.Println(s1)
	// var s2 string = s1 + 10 没有这么写的，报错

	// 除号
	fmt.Println(10 / 3)   // 两个int / 结果为int
	fmt.Println(10.0 / 3) // 只要一个为浮点型结果就为浮点类型
	f1 := 10.0 / 3
	fmt.Printf("f1 type: %T \n", f1) // float64

	// %
	fmt.Println(10 % 3) // 10-10/3*3
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)

	var a int = 10
	a++
	fmt.Println(a)
	// fmt.Println(++a) 错误！
	// fmt.Println(a++) 错误！
	a--
	fmt.Println(a)
}
