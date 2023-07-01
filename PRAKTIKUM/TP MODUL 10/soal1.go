package main

import "fmt"

const NMAX int = 12345

type prodi struct {
	nama, akreditasi      string
	tahun, aktif, lulusan int
}

type fakultas struct {
	nama     string      // nama fakultas
	arrProdi [NMAX]prodi // array nyimpen list prodi
	N        int         // jumlah  prodi yang ada di arrProdi
}

func buatFakultas_1301223029(f *fakultas) {
	var namaF string
	fmt.Scan(&namaF)
	f.nama = namaF
}

func daftarProdi_1301223029(f *fakultas) {
	var (
		namaP  prodi
		indeks int
	)

	fmt.Scan(&namaP.nama, &namaP.akreditasi, &namaP.tahun, &namaP.aktif, &namaP.lulusan)
	indeks = cekProdi_1301223029(namaP.nama, *f)
	if indeks == 1 {
		f.arrProdi[indeks].aktif += namaP.aktif
		f.arrProdi[indeks].lulusan += namaP.lulusan
	} else {
		f.arrProdi[f.N] = namaP
		f.N += 1
	}
}

func cekProdi_1301223029(s string, f fakultas) int {
	for i := 0; i < f.N; i++ {
		if s == f.arrProdi[i].nama {
			return i
		}
	}
	return -1
}

func terbanyak_1301223029(f fakultas) prodi {
	var (
		max int = 0
	)

	for i := 1; i < f.N; i++ {
		if f.arrProdi[i].aktif+f.arrProdi[i].lulusan > f.arrProdi[max].aktif+f.arrProdi[max].lulusan {
			max = i
		}
	}
	return f.arrProdi[max]

}

func termuda_1301223029(f fakultas) int {
	var (
		min int = 0
	)

	for i := 1; i < f.N; i++ {
		if f.arrProdi[i].tahun > f.arrProdi[min].tahun {
			min = i
		}
	}
	return min
}

func prestasi_1301223029(f fakultas) string {
	var arrA = []string{"unggul", "baik", "cukup"}
	var (
		u, b, c int
		max     int
		hasil   string
	)
	for i := 0; i < f.N; i++ {
		if f.arrProdi[i].akreditasi == arrA[0] {
			u++
		}
		if f.arrProdi[i].akreditasi == arrA[1] {
			b++
		}
		if f.arrProdi[i].akreditasi == arrA[2] {
			c++
		}
	}

	max = u
	if max < b {
		max = b
	}
	if max < c {
		max = c
	}

	if max == u {
		hasil = arrA[0]
	} else if max == b {
		hasil = arrA[1]
	} else {
		hasil = arrA[2]
	}
	return hasil
}

func cetakProdi_1301223029(f fakultas, x string) {
	for i := 0; i < f.N; i++ {
		if f.arrProdi[i].akreditasi == x {
			fmt.Print(f.arrProdi[i].nama, " ")
		}
	}
	fmt.Print("\n")
}

func main() {
	var (
		fif     fakultas
		j, muda int
		Pro     prodi
		pres    string
	)

	buatFakultas_1301223029(&fif)
	for j = 1; j <= 10; j++ {
		daftarProdi_1301223029(&fif)
	}

	Pro = terbanyak_1301223029(fif)
	fmt.Print("Prodi ", Pro.nama, " memiliki mahasiswa dan lulusan terbanyak yaitu ", Pro.aktif+Pro.lulusan, "\n")
	stop := false
	for i := 0; i < fif.N && !stop; i++ {
		muda = termuda_1301223029(fif)
		j = fif.N - 1 - i
		if fif.arrProdi[muda].tahun == fif.arrProdi[j].tahun {
			stop = true
			fmt.Print("Prodi termuda adalah ", fif.arrProdi[i].nama, "\n")
		} else {
			stop = false
		}
	}
	pres = prestasi_1301223029(fif)
	fmt.Print("Akreditasi prodi terbanyak di Fakultas ", fif.nama, " adalah ", pres, " ")
	cetakProdi_1301223029(fif, pres)

}
