package main

import "fmt"

type AsDos_T struct {
	kode   string
	nama   string
	nim    int
	mk     string
	jumlah int
}

type TabAsDos_T [2500]AsDos_T

func BacaAsDos_1301223029(tabel *TabAsDos_T, n *int) {
	var (
		kode_sementara string
	)
	*n = 0
	fmt.Print("Masukkan kode, nama, nim, mk, jumlah : \n")
	fmt.Scan(&kode_sementara)
	for kode_sementara != "XXX" && *n < 2500 {
		tabel[*n].kode = kode_sementara
		fmt.Scan(&tabel[*n].nama, &tabel[*n].nim, &tabel[*n].mk, &tabel[*n].jumlah)
		*n++
		fmt.Scan(&kode_sementara)
	}

}

func CetakAsDos_1301223029(tabel TabAsDos_T, n int, mk string) {
	for i := 0; i < n; i++ {
		if tabel[i].mk == mk {
			fmt.Print("\nNama : ", tabel[i].nama, "\nkode : ", tabel[i].kode, "\nmk : ", tabel[i].mk, "\n")
		}
	}
}

func main() {
	var (
		data   TabAsDos_T
		n      int
		matkul string
	)
	BacaAsDos_1301223029(&data, &n)
	fmt.Print("\nMasukkan matakuliah yang ingin dicari : ")
	fmt.Scan(&matkul)

	fmt.Print("\nBerikut asisten dosen yang mengajar matakuliah ", matkul, " : ")
	CetakAsDos_1301223029(data, n, matkul)
}
