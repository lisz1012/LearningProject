package main

import "fmt"

func main() {
	// 类型转换
	var n1 int = 100
	var n2 float64 = float64(n1) // n1本身类型不变
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Printf("n1 type: %T\n", n1)
	fmt.Printf("n2 type: %T\n\n", n2)

	// 编译不会出错，但是只截取了第8位
	var n3 int64 = 888888
	var n4 int8 = int8(n3)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Printf("n3 type: %T\n", n3)
	fmt.Printf("n4 type: %T\n\n", n4)

	var n5 int32 = 12
	var n6 int64 = int64(n5) + 30 // n5必须得转一下，否则两数之和是int32类型
	println(n6)

	var n7 int64 = 12
	var n8 int8 = int8(n7) + 127 // -117
	//var n9 int8 = int8(n7) + 128 编译都不通过
	fmt.Println(n8)
	fmt.Println(n9)
}
