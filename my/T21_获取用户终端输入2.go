package main

import "fmt"

func main() {
	fmt.Println("请录入学生的姓名，年龄，成绩，是否是VIP，使用空格分隔")
	var name string
	var age int
	var score float64
	var isVIP bool
	fmt.Scanf("%s %d %f %t", &name, &age, &score, &isVIP)
	fmt.Printf("姓名：%v   年龄：%v   分数：%v   VIP： %v \n", name, age, score, isVIP)
}
