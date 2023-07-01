package main

import "fmt"

func abc(x int) string {
	var (
		s string = ""
		a int
	)
	if x == 0 {
		return s
	}

	a = x % 2
	if a == 1 {
		s = "1" + abc(x/2)
	} else {
		s = "0" + abc(x/2)
	}
	return s
}

func main() {
	var (
		a int
	)
	fmt.Scan(&a)
	fmt.Print(abc(a))
}
