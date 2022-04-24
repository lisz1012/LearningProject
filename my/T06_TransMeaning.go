package main

import "fmt"

func main() {
	fmt.Println("aaa\nbbb")
	fmt.Println("aaa\bbbb")   // aabbb
	fmt.Println("aaa\b")      // aaa
	fmt.Println("aaa\rbbb")   // bbb
	fmt.Println("aaaaa\rbbb") // bbbaa \r是移动光标并覆盖当前位置上的字符
	fmt.Println("aaaaaaaaaaaa")
	fmt.Println("aaaaa\tbbbbb")    // aaaaa   bbbbb
	fmt.Println("aaaa\tbbbbb")     // aaaa    bbbbb 每8个字符位算一个制表符
	fmt.Println("aaaaaaaa\tbbbbb") // aaaaaaaa        bbbbb 每8个字符位算一个制表符
	fmt.Println("\"Hello\"")

}
