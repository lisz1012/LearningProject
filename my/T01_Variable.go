package main

import "fmt"

func main() {
	// 变量的声明
	var age int
	//  变量的赋值
	age = 10
	// 变量的使用：只声明和赋值而不是用变量是不允许的
	fmt.Println("age =", age)

	// 变量的声明和赋值合成一句
	var age2 = 19
	fmt.Println("age2 = ", age2)

	var num float32 = 12.56
	fmt.Println("num =", num)
}
