package main

import "fmt"

func pola(x int) {
	if x == 1 {
		fmt.Println("*")
	} else {
		pola(x - 1)
		for i := 1; i <= x; i++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func main() {
	var (
		x int
	)
	fmt.Scan(&x)
	pola(x)
}
