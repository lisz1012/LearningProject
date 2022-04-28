package main

import "fmt"

func main() {
	var n1 int = 19
	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("s1的类型是： %T, s1 = %q\n", s1, s1)

	var n2 float32 = 4.78
	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Printf("s2的类型是： %T, s2 = %q\n", s2, s2)

	var n3 bool = false
	var s3 string = fmt.Sprintf("%t", n3)
	fmt.Printf("s3的类型是： %T, s3 = %q\n", s3, s3)

	var n4 byte = 'c'
	var s4 string = fmt.Sprintf("%c", n4)
	fmt.Printf("s3的类型是： %T, s4 = %q\n", s4, s4)
}
