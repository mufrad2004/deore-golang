package main

import (
	"fmt"
)

func fx(x int) int {
	return x * x
}

func gx(x int) int {
	return x - 2
}

func hx(x int) int {
	return x + 1
}

func main() {
	var (
		a, b, c int
	)
	fmt.Scan(&a, &b, &c)
	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
}
