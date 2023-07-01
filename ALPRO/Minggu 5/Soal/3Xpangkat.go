package main

import "fmt"

func pangkat(x, y int) int {
	var (
		hasil int = 1
	)
	if y == 0 {
		return hasil
	} else {
		hasil *= x * pangkat(x, y-1)
		return hasil
	}
}

func main() {
	var (
		x, y int
	)

	fmt.Scan(&x, &y)
	fmt.Print(pangkat(x, y))
}
