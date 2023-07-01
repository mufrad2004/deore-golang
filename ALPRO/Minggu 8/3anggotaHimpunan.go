package main

import "fmt"

type himpunan [100]int

func isiTabel(t *himpunan, n int) {
	var isi int
	fmt.Scan(&isi)
	for isi != 0 {
		t[n] = isi
		n++
		fmt.Scan(&isi)
	}

}

func cekTabel(t himpunan, n int) bool {
	var (
		count int
		hasil int
	)
	count = t[0]
	for i := 1; i < n; i++ {
		if count == t[i] {
			hasil++
		}
	}
	if hasil > 0 {
		return false
	} else {
		return true
	}

}

func main() {
	var (
		n   int = 0
		tab himpunan
	)

	isiTabel(&tab, n)
	fmt.Print(cekTabel(tab, n))
}
