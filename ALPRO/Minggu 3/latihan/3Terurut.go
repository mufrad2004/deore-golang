package main

import "fmt"

func mengurutkan(a, b *int) {
	var (
		c int
	)
	if *a > *b {
		c = *a
		*a = *b
		*b = c
		// b - a
	}
}

func main() {
	var (
		a, b int
	)

	fmt.Scan(&a, &b)

	mengurutkan(&a, &b)
	fmt.Print(a, b)

}
