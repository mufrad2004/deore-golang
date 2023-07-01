package main

import "fmt"

type data struct {
	nama string
}

type dataUnik struct {
	nama string
	jml  int
}

type tabData struct {
	arrData [1000]data
	n       int
}

type tabDataUnik struct {
	arrData [1000]dataUnik
	n       int
}

func isiData(t *tabData) {
	var isi data
	fmt.Scan(&isi.nama)
	for isi.nama != "-1" {
		t.arrData[t.n] = isi
		t.n++
		fmt.Scan(&isi.nama)
	}
}

func cekDataUnik(t tabDataUnik, target string) int {
	var hasil int = -1
	for i := 0; i < t.n; i++ {
		if target == t.arrData[i].nama {
			hasil = i
		}
	}
	return hasil
}

func isiDataUnik(t *tabDataUnik, t2 tabData) {
	for i := 0; i < t2.n; i++ {
		if t.n == 0 {
			t.arrData[t.n].nama = t2.arrData[i].nama
			t.arrData[t.n].jml++
			t.n++
		} else {
			if cekDataUnik(*t, t2.arrData[i].nama) == -1 {
				t.arrData[t.n].nama = t2.arrData[i].nama
				t.arrData[t.n].jml++
				t.n++
			} else {
				t.arrData[cekDataUnik(*t, t2.arrData[i].nama)].jml++
			}
		}
	}

}

func printData(t tabDataUnik) {
	for i := 0; i < t.n; i++ {
		fmt.Println(t.arrData[i])
	}
}

func main() {
	var t tabData
	var t2 tabDataUnik
	isiData(&t)

	isiDataUnik(&t2, t)
	fmt.Println("Hasil : ")
	printData(t2)

}
