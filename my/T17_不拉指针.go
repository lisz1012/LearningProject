package main

import "fmt"

func main() {
	var n1 int = 10
	fmt.Println(n1)
	var ptr *int = &n1
	*ptr = 20
	fmt.Println(n1)

	var n2 int = 22
	ptr = &n2
	fmt.Println(*ptr)
}
