package main

import "fmt"

func main() {
	var age int = 18
	fmt.Println(&age)

	// 定义一个指针变量, 指向int类型的指针
	var ptr *int = &age
	fmt.Printf("ptr 的类型是 %T ptr = %v \n", ptr, ptr)
	age = *ptr
	fmt.Printf("age 的类型是 %T age = %v \n", age, age)
	fmt.Printf("&ptr 的类型是 %T &prt = %v \n", &ptr, &ptr)
	fmt.Printf("ptr 指向的数值的类型是 %T ptr 指向的数值是: %v \n", *ptr, *ptr)
}
