package main

import "fmt"

type himpunan [50]int

func isi(t *himpunan, n *int) {
	*n = 0
	var masuk int
	fmt.Scan(&masuk)
	for masuk != 0 {
		t[*n] = masuk
		*n++
		fmt.Scan(&masuk)
	}
}

func cek(t1, t2 himpunan) {
	var n int
	var (
		hasil []int
	)
	if len(t1) < len(t2) {
		n = len(t1)
	} else {
		n = len(t2)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if t1[i] == t2[j] {
				hasil = append(hasil, t1[i])
			}
		}
	}

	for i := 0; i < len(hasil); i++ {
		if hasil[i] != 0 {
			fmt.Print(hasil[i], " ")
		}
	}
}

func main() {
	var (
		a          int
		tab1, tab2 himpunan
	)

	isi(&tab1, &a)
	isi(&tab2, &a)
	// fmt.Println(tab1)
	// fmt.Println(tab2)
	cek(tab1, tab2)

}
