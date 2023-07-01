package main

import "fmt"

func faktor(i, x int) {
	if i == x {
		fmt.Print(i)
	} else {
		if x%i == 0 {
			fmt.Print(i, " ")
		}
		faktor(i+1, x)
	}

}

func main() {
	var (
		x int
	)

	fmt.Scan(&x)

	faktor(1, x)
}
