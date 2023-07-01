package main

import "fmt"

func bilangan(x int) {
	if x == 1 {
		fmt.Print(x, " ")
	} else {
		bilangan(x - 1)
		if x%2 == 1 {
			fmt.Print(x, " ")
		}
	}
}

func main() {
	var (
		a int
	)

	fmt.Scan(&a)
	bilangan(a)
}
