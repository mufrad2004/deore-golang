package main

import "fmt"

func pangkatA(i, n int) int {
	if i == n {
		return 2
	} else {
		return 2 * pangkatA(i+1, n)
	}
}

func pangkatB(n int) int {
	if n == 1 {
		return 2
	} else {
		return 2 * pangkatB(n-1)
	}
}
func main() {
	var (
		n int
		i int = 1
	)
	fmt.Scan(&n)
	fmt.Println("Menggunakan pangkatA : ", pangkatA(i, n))
	fmt.Println("Menggunakan pangkatB : ", pangkatB(n))

}
