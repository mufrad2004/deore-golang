package main

import "fmt"

func sum_verA(i, n int) int {
	if i == n {
		return i
	} else {
		return i + sum_verA(i+1, n)
	}
}

func sum_verB(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + sum_verB((n - 1))
	}

}

func main() {
	var (
		angka1, angka2 int
	)
	fmt.Scan(&angka1, &angka2)
	fmt.Println(sum_verA(angka1, angka2))
	fmt.Println(sum_verB(angka1))
	fmt.Println(sum_verB(angka2))
}
