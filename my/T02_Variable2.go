package main

import "fmt"

// 全剧变量
var n7 = 100
var n8 = 9.8

// 设计者认为以上的全局变量的写法太麻烦了
var (
	n9  = 500
	n10 = "netty"
)

func main() { //{} 里面的叫局部变量
	var num int = 8
	fmt.Println("num =", num)

	var num2 int
	fmt.Println("num2 =", num2) // 输出0. 默认值=0

	var num3 = 10
	fmt.Println("num3 = ", num3)

	var num4 = .9
	fmt.Println("num4 =", num4) // num4 = 0.9

	var str = "hahaha"
	fmt.Println(str)

	sex := "男"
	fmt.Println(sex)
	sex = "女"
	fmt.Println(sex)

	var sex1 = "haha"
	fmt.Println("sex1 =", sex1)

	var n1, n2, n3 int
	n1 = 2
	n2 = 3
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3) // n1 = 2 n2 = 3 n3 = 0

	var n4, name, n5 = 10, "Jack", 7.8
	fmt.Println("n4 =", n4, "name =", name, "n5 =", n5) // n4 = 10 name = Jack n5 = 7.8

	n6, height := 6.9, 10.1
	fmt.Println("n6 =", n6, "height =", height)

	fmt.Println("n7 =", n7, "n8 =", n8)
	fmt.Println("n9 =", n9, "n10 =", n10)
}
