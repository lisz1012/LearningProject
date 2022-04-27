package main

import "fmt"

func main() {
	var s1 string = "你好，全面拥抱go语言"
	fmt.Println(s1)
	var s2 string = "abc"
	s2 = "def"
	fmt.Println(s2)
	// 无特殊字符，用双引号没啥问题.如果又很多需要转义的，用`` 反引号就比较方便了, 前提是不能有反引号
	var s3 string = `abc"ascad`
	fmt.Println(s3)
	var s4 string = `aaabbb`
	fmt.Println(s4)
	// 字符串的拼接
	var s5 string = "abc" + "def"
	fmt.Println(s5)
	s5 += "hijk"
	fmt.Println(s5)
	// 当一个字符串过长的时候
	var s6 string = "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" + // + 留在上面一个行
		"abc" + "def" + "abc" + "def" + "abc" + "def"
	fmt.Println(s6)
}
