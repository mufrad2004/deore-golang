package main

import "fmt"

type tabGol [1000]tim
type tim struct {
	tim1, tim2     int
	total1, total2 int
}

func inputData(t *tabGol, n int) {
	var (
		x tim
	)
	for i := 0; i < n; i++ {
		fmt.Scan(&x.tim1, &x.tim2)
		t[i] = x
	}
}

func rataan(t tabGol, n int, tim1, tim2 *float64) {
	var (
		x tim
	)
	for i := 0; i < n; i++ {
		x.total1 += t[i].tim1
		x.total2 += t[i].tim2
	}

	*tim1 = float64(x.total1) / float64(n)
	*tim2 = float64(x.total2) / float64(n)
}

func main() {
	var (
		a            int
		b            tabGol
		rata1, rata2 float64
	)
	fmt.Scan(&a)
	inputData(&b, a)
	rataan(b, a, &rata1, &rata2)
	fmt.Print(rata1, rata2)

}
