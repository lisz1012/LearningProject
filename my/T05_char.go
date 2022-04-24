package main

import "fmt"

func main() {
	var c1 byte = 'a'
	fmt.Println(c1)

	var c2 byte = 'b'
	fmt.Println(c2)

	var c3 byte = '('
	fmt.Println(c3)

	var c4 byte = 7
	fmt.Println(c4)
	// 字符类型本质上就是整数，可以参与运算，输出字符的时候，输出了对应的码值
	// 字母数字标点等，底层是按照ascii码存储的
	fmt.Println(c3 + 60)

	// 汉字底层是Unicode的码值。'哈'对应的是21704，可以用int，而不是byte
	var c5 int = '哈'
	fmt.Println(c5)

	// 总结golang的自负对应的是用的是utf-8编码。Unicode是对应的字符集，utf-8是其中的一种编码方案而已
	// 格式化打印可以打印出对应的字符
	fmt.Printf("%c", c5)
}
