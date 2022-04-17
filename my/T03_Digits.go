package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int8 = 23
	fmt.Println("num1 =", num1)

	var num2 int16 = 23333
	fmt.Println("num2 =", num2)

	var num3 uint8 = 255
	fmt.Println("num3 =", num3)

	var num4 = 28
	fmt.Printf("num4的类型是：%T", num4)
	fmt.Println(" num4 =", num4)
	fmt.Println("num4 占用的字节数：", unsafe.Sizeof(num4))

	var age byte = 33
	fmt.Println("age = ", age)
}
