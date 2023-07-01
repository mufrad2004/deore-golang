package main

import "fmt"

const nmax int = 513

type wisudawan struct {
	nama, nim string
	ipk       float64
	semester  int
}

type tabWisudawan struct {
	arrWisudawan [nmax]wisudawan
	n            int
}

func isiData(t *tabWisudawan) {
	var isi wisudawan
	fmt.Scanln(&isi.nama, &isi.nim, &isi.ipk, &isi.semester)

	for isi.nim != "000000" && t.n < nmax {
		t.arrWisudawan[t.n] = isi
		t.n++
		fmt.Scanln(&isi.nama, &isi.nim, &isi.ipk, &isi.semester)
	}
}

func tampilData(t tabWisudawan) {
	for i := 0; i < t.n; i++ {
		fmt.Println(t.arrWisudawan[i])
	}
}

func pisahkanData(t tabWisudawan, tabC, tabB *tabWisudawan) {
	for i := 0; i < t.n; i++ {
		if lulusCumLaude(t.arrWisudawan[i].ipk, t.arrWisudawan[i].semester) {
			tabC.arrWisudawan[i] = t.arrWisudawan[i]
		} else {
			tabB.arrWisudawan[i] = t.arrWisudawan[i]
		}
	}
}

func lulusCumLaude(ipk float64, semester int) bool {
	var hasil bool = false

	if ipk >= 3.5 && semester <= 8 {
		hasil = true
	}
	return hasil
}

func rataRata(t tabWisudawan) float64 {
	var (
		totalIPK float64
	)

	for i := 0; i < t.n; i++ {
		totalIPK += t.arrWisudawan[i].ipk
	}

	return totalIPK / float64(t.n)
}

func main() {
	var (
		t          tabWisudawan
		tabC, tabB tabWisudawan
	)
	isiData(&t)

	pisahkanData(t, &tabC, &tabB)

	// * mahasiswa cumlaude
	tampilData(tabC)
	// * mahasiswa non cumlaude
	tampilData(tabB)

	// * rataan mahasiswa cumlaude
	fmt.Println("Rata-rata IPK wisudawan cum laude : ", rataRata(tabC))
	// * rataan mahasiswa semua
	fmt.Println("Rata-rata IPK semua wisudawan : ", rataRata(t))

}
