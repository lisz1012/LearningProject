package main

import "fmt"

// float32/64可以用小数点也可以用科学计数法，e可以大小写
func main() {
	var num1 float32 = 3.14
	fmt.Println(num1)
	num1 = -3.14
	fmt.Println(num1)
	num1 = 314e+2
	fmt.Println(num1)
	num1 = 314e+2
	fmt.Println(num1)
	num1 = 314e-2
	fmt.Println(num1)

	var num2 float64 = 3.14
	fmt.Println(num2)
	num2 = -3.14
	fmt.Println(num2)
	num2 = 314e+2
	fmt.Println(num2)
	num2 = 314e+2
	fmt.Println(num2)
	num2 = 314e-2
	fmt.Println(num2)

	var num3 float32 = 256.000000916
	fmt.Println(num3)
	var num4 float64 = 256.000000916 // 建议使用，精度比较高
	fmt.Println(num4)
	var num5 float64 = 256.00000000000000000000000000000000000000000916 // 即使是float64，也精确不到这么多小数位
	fmt.Println(num5)

	var num6 = 3.17
	fmt.Printf("num6 的类型是：%T", num6) //  默认小数类型就是float64
}
