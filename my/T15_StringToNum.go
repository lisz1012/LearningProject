package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string = "true"
	// 返回两个值，结果和err， _表示忽略错误
	b, _ := strconv.ParseBool(s1)
	fmt.Printf("b 的类型是  %T, b = %v \n", b, b)

	var s2 string = "19"
	var num1 int64
	num1, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Printf("num1 的类型是  %T, num1 = %v \n", num1, num1)

	var s3 string = "3.14"
	num2, _ := strconv.ParseFloat(s3, 64)
	fmt.Printf("num2 的类型是  %T, num2 = %v \n", num2, num2)

	var s4 string = "golang" // 无效数据按照默认值返回
	// 返回两个值，结果和err， _表示忽略错误
	b1, _ := strconv.ParseBool(s4)
	fmt.Printf("b1 的类型是  %T, b1 = %v \n", b1, b1)

	var s5 string = "golang" // 无效数据按照默认值返回, 要确定能转成有效的类型
	// 返回两个值，结果和err， _表示忽略错误
	b2, _ := strconv.ParseInt(s5, 10, 64)
	fmt.Printf("b2 的类型是  %T, b2 = %v \n", b2, b2)
}
