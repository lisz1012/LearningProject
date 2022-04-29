package main

import "fmt"

func main() {
	var count int = 200
	if count < 30 {
		fmt.Println("对不起，口罩存量不足")
	} else {
		fmt.Println("库存充足")
	}
}
