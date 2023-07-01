package main

import "fmt"

// buatkan kodingan mengurutkan ipk mahasiswa secara ascending

const NMAX = 100

type mahasiswa struct {
	nama, nim string
	ipk       float64
}

type tabMhs [NMAX]mahasiswa

func isiTabel(T *tabMhs, n *int) {
	var (
		isi mahasiswa
		i   int = 0
	)

	fmt.Scan(&isi.nama, &isi.nim, &isi.ipk)

	for isi.nama != "X" && isi.nim != "X" && isi.ipk != 0.0 {
		*n++
		T[i] = isi
		i++
		fmt.Scan(&isi.nama, &isi.nim, &isi.ipk)
	}

}

func ascendingIPK(T *tabMhs, n *int) {
	var (
		pass, idx int
		temp      mahasiswa
	)

	for pass = 1; pass < *n; pass++ {
		// orang pertama yang megang buku
		idx = pass - 1
		// Bandingkan sama setiap orang mulai dari orang ke 2
		for i := pass; i < *n; i++ {
			// Jika orang yang megang buku lebih besar , maka kasih bukunya ke orang ke i
			if T[idx].ipk > T[i].ipk {
				idx = i
			}
		}
		temp = T[idx]
		T[idx] = T[pass-1]
		T[pass-1] = temp
	}
}

func main() {
	var (
		TabelM tabMhs
		n      int = 0
	)

	isiTabel(&TabelM, &n)
	ascendingIPK(&TabelM, &n)

	for i := 0; i < n; i++ {
		fmt.Print(TabelM[i], " ")
	}
}
