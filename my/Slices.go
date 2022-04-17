package main

import "fmt"

func main() {
	var x []int
	fmt.Printf("%p\n", x)
	x = append(x, 1)
	fmt.Printf("%p\n", x)
	x = append(x, 2)
	fmt.Printf("%p\n", x)
	x = append(x, 3)
	fmt.Printf("%p\n", x)
}
