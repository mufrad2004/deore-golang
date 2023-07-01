package main

import "fmt"

const nmax int = 1024

type angka struct {
	angka int
}

type tabAngka struct {
	arrAngka [nmax]angka
	n        int
}

func isiAngka(t *tabAngka) {
	var isi angka
	fmt.Scan(&isi.angka)
	for isi.angka != -1 {
		t.arrAngka[t.n] = isi
		t.n++
		fmt.Scan(&isi.angka)
	}
}

func printData(t tabAngka) {
	for i := 0; i < t.n; i++ {
		fmt.Println(t.arrAngka[i])
	}
}

func cekAngka(t tabAngka, target int) int {
	var hasil int = -1
	for i := 0; i < t.n; i++ {
		if target == t.arrAngka[i].angka {
			hasil = i
		}
	}

	return hasil
}

func hapusData(t *tabAngka, target int) {
	var index = cekAngka(*t, target)

	if index == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		for index < t.n {
			t.arrAngka[index] = t.arrAngka[index+1]
			index++
		}
		t.n--
	}
}

func main() {
	var (
		t      tabAngka
		target int
	)
	isiAngka(&t)
	printData(t)
	fmt.Print("\nMau apus angka brp ? ")
	fmt.Scan(&target)
	hapusData(&t, target)
	fmt.Println("\nSetelah dihapus ")
	printData(t)
}
