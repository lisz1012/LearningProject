package main

import "fmt"

func main() {
	var score int = 11
	switch score / 10 {
	case 9, 10: // 可以写两个，逗号隔开，自己研究出来的^_^
		fmt.Println("A") // 无需break
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	case 5, 4, 3, 2, 1, 0:
		fmt.Println("F")
	default:
		fmt.Println("成绩非法") // 可以放在case前面
	}
}
