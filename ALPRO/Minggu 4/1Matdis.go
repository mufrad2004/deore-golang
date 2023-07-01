package main

import "fmt"

func mencariFaktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}

}

func permutation(n, r int) int {
	var (
		tuker int
		nr    int
	)
	nr = n - r

	if n < r {
		tuker = n
		n = r
		r = tuker
	}

	mencariFaktorial(n, &n)
	mencariFaktorial(nr, &nr)
	return n / nr

}

func combinatio(n, r int) int {
	var (
		nr, tuker int
	)
	nr = n - r
	if n < r {
		tuker = n
		n = r
		r = tuker
	}
	mencariFaktorial(n, &n)
	mencariFaktorial(r, &r)
	mencariFaktorial(nr, &nr)
	return n / (r * nr)
}

func main() {
	var (
		a, b, c, d int
	)

	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combinatio(a, c))
	fmt.Println(permutation(b, d), combinatio(b, d))
}
