package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)
	if (x%2 == 0 && x%3 == 0) || (x%3 == 0 && x%5 == 0) {
		fmt.Print("bilangan kelipatan 2 dan 3, atau kelipatan 3 dan 5")
	} else {
		fmt.Print("BUKAN kelipatan 2 dan 3, juga BUKAN kelipatan 3 dan 5")
	}
}
