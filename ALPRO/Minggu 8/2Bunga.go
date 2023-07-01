package main

import (
	"fmt"
)

const NMAX int = 100

type tabBunga [NMAX]string

func cari(t tabBunga, n int, bunga string) int {
	var found int
	for i := 0; i < n; i++ {
		if bunga == t[i] {
			found = i
		} else {
			found = -1
		}
	}
	return found
}

func rename(t *tabBunga, n int, X string) {
	var (
		found int
	)

	found = cari(*t, n, X)
	for i := 0; i < n; i++ {
		if found == -1 {
			fmt.Print("Bunga tidak ditemukan")
		} else {
			fmt.Scan(&t[found])
		}
	}
}

func delete(t *tabBunga, n int, X string) {
	for i := 0; i < n; i++ {
		if t[i] != X {
			fmt.Print("Bunga tidak ditemukan")
		} else {
			t[i] = t[i+1]
			n--
		}
	}
}

func main() {
	var (
		a     string
		bunga string
		n     int
		tab   tabBunga
	)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		tab[i] = a
	}
	fmt.Scan(&bunga)

	rename(&tab, n, bunga)
	delete(&tab, n, bunga)
	for i := 0; i < n; i++ {
		fmt.Println(tab[i])
	}

}
