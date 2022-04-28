package main

import "fmt"

func main() {
	fmt.Println("姓名：")
	var s string
	fmt.Scanln(&s)

	fmt.Println("年龄")
	var n int
	fmt.Scanln(&n)

	fmt.Println("分数")
	var f float64
	fmt.Scanln(&f)

	fmt.Println("VIP?")
	var b bool
	fmt.Scanln(&b)

	fmt.Printf("姓名： %v  年龄： %v 分数： %v  VIP： %v", s, n, f, b)
}
