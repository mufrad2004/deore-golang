package main

import "fmt"

func procB(b, c *int) {
	var (
		a int
	)

	*b = *b + *c
	*c = a + *b + *c
}

func main() {
	var (
		a int = 10
	)

	procB(&a, &a)
	fmt.Print(a)
}
