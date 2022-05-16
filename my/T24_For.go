package main

func main() {
	var sum int = 0
	for i := 1; i <= 5; i++ { // for的初始表达式只能用 :=
		sum += i
	}
	println(sum)
}
