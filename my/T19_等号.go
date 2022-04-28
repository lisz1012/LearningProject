package main

import "fmt"

func main() {
	var n1 int = 10
	fmt.Println(n1)
	var n2 int = (10+20)%3 + 3 - 7 // 等号又测得知运算清楚后再赋值给左侧
	fmt.Println(n2)

	var n3 int = 10
	n3 += 10
	fmt.Println(n3)
	n3 *= 3
	fmt.Println(n3)
	n3 /= 2
	fmt.Println(n3)
	n3 %= 7
	fmt.Println(n3)

	var a int = 3
	var b int = 4
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("a = %d  b = %d \n", a, b)

	//c := 1
	//d := 1.0
	//fmt.Println(c == d) // 类型不一样没法比
}
