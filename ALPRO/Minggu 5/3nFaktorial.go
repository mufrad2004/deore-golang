package main

import "fmt"

func faktorialA(n int) int {
	if n == 1 || n == 0 {
		return 1
	} else {
		return n * faktorialA(n-1)
	}
}

func faktorialB(i, n int) int {
	if i == n {
		return i
	} else {
		return i * faktorialB(i+1, n)
	}

}

func main() {
	var (
		i int = 1
		n int
	)

	fmt.Scan(&n)

	fmt.Println("Menggunakan fakorialA : ", faktorialA(n))
	fmt.Println("Menggunakan fakorialB : ", faktorialB(i, n))

}
