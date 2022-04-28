package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int = 18
	var s1 string = strconv.FormatInt(int64(n1), 10)
	fmt.Printf("s1 的类型是  %T, s1 = %q \n", s1, s1)

	var n2 float64 = 6.29
	// 'f' s是十进制，9表示小数点之后9位，第四个参数表示这个小数是64位浮点型
	var s2 string = strconv.FormatFloat(n2, 'f', 9, 64)
	fmt.Printf("s2 的类型是  %T, s2 = %q \n", s2, s2)

	var n3 bool = true
	s3 := strconv.FormatBool(n3)
	fmt.Printf("s3 的类型是  %T, s3 = %q \n", s3, s3)
}
