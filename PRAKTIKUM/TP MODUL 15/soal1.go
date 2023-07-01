package main

import "fmt"

const nmax int = 1024

type negara struct {
	nama  string
	tahun int
	gpd   int
}
type tabNegara struct {
	arrNegara [nmax]negara
	n         int
}

type negara2 struct {
	nama                      string
	gpd2021, gpd2022, gpd2023 int
}

type tabNegara2 struct {
	arrNegara [nmax]negara2
	n         int
}

func isiNegara(t *tabNegara) {
	var isi negara
	fmt.Scanln(&isi.nama, &isi.tahun, &isi.gpd)
	for isi.nama == "Brunei" || isi.nama == "Indonesia" || isi.nama == "Singapore" || isi.nama == "Laos" || isi.nama == "Philippines" || isi.nama == "Thailand" || isi.nama == "Vietnam" || isi.nama == "Myanmar" || isi.nama == "Cambodia" || isi.nama == "Malaysia" {
		t.arrNegara[t.n] = isi
		t.n++
		fmt.Scanln(&isi.nama, &isi.tahun, &isi.gpd)
	}
}

func cekNegara2(t tabNegara2, negara string) bool {
	var hasil bool = false
	for i := 0; i < t.n; i++ {
		if negara == t.arrNegara[i].nama {
			hasil = true
		}
	}
	return hasil
}

func isiNamaNegara2(t tabNegara, t2 *tabNegara2) {
	var isi negara
	for i := 0; i < t.n; i++ {
		isi = t.arrNegara[i]
		if cekNegara2(*t2, isi.nama) {
		} else {
			t2.arrNegara[t2.n].nama = isi.nama
			t2.n++
		}
	}
}
func isiGPDNegara2(t tabNegara, t2 *tabNegara2) {
	for i := 0; i < t2.n; i++ {
		isi := t2.arrNegara[i]
		for j := 0; j < t.n; j++ {
			if isi.nama == t.arrNegara[j].nama {
				if t.arrNegara[j].tahun == 2023 {
					t2.arrNegara[i].gpd2023 = t.arrNegara[j].gpd
				} else if t.arrNegara[j].tahun == 2022 {
					t2.arrNegara[i].gpd2022 = t.arrNegara[j].gpd
				} else {
					t2.arrNegara[i].gpd2021 = t.arrNegara[j].gpd
				}
			}
		}
	}
}

func sortingGPD(t2 *tabNegara2, tahun int) {
	var (
		idx  int
		temp negara2
		pass int
	)

	for pass = 1; pass < t2.n; pass++ {
		idx = pass - 1
		for i := pass; i < t2.n; i++ {
			if tahun == 2021 {
				if t2.arrNegara[idx].gpd2021 < t2.arrNegara[i].gpd2021 {
					idx = i
				}
			} else if tahun == 2022 {
				if t2.arrNegara[idx].gpd2022 < t2.arrNegara[i].gpd2022 {
					idx = i
				}
			} else {
				if t2.arrNegara[idx].gpd2023 < t2.arrNegara[i].gpd2023 {
					idx = i
				}
			}
		}
		temp = t2.arrNegara[pass-1]
		t2.arrNegara[pass-1] = t2.arrNegara[idx]
		t2.arrNegara[idx] = temp
	}
}

func tampilkanNegara2(t tabNegara2) {
	for i := 0; i < t.n; i++ {
		fmt.Println(t.arrNegara[i])
	}
}

func main() {
	var (
		t          tabNegara
		t2         tabNegara2
		tahunAkhir int
	)
	isiNegara(&t)
	fmt.Scan(&tahunAkhir)
	fmt.Print("\n\n")
	isiNamaNegara2(t, &t2)
	isiGPDNegara2(t, &t2)
	sortingGPD(&t2, tahunAkhir)
	tampilkanNegara2(t2)
}

/* input 1
Brunei 2023 889
Singapore 2021 676
Cambodia 2021 881
Myanmar 2021 950
Singapore 2023 712
Philippines 2021 152
Malaysia 2022 282
Philippines 2023 120
Indonesia 2023 62
Laos 2022 497
Vietnam 2021 29
Cambodia 2023 277
Brunei 2022 469
Myanmar 2023 480
Thailand 2023 393
Thailand 2021 361
Laos 2021 108
Laos 2023 31
Malaysia 2023 968
Indonesia 2021 543
Vietnam 2023 590
Cambodia 2022 293
Brunei 2021 333
Singapore 2022 82
Myanmar 2022 354
Malaysia 2021 476
Thailand 2022 803
Philippines 2022 913
Vietnam 2022 711
Indonesia 2022 876
Spain 2023
*/

/*input 2

Vietnam 2022 615
Malaysia 2022 163
Indonesia 2022 884
Philippines 2023 37
Indonesia 2023 1893
Philippines 2021 286
Laos 2021 98
Myanmar 2022 135
Singapore 2021 357
Vietnam 2021 638
Thailand 2021 724
Cambodia 2021 415
Malaysia 2021 176
Brunei 2023 80
Thailand 2022 575
Malaysia 2023 649
Laos 2022 740
Vietnam 2023 49
Thailand 2023 870
Japan 2021
*/

/*
{Brunei 2023 889}
{Singapore 2023 712}
{Philippines 2023 120}
{Indonesia 2023 62}
{Cambodia 2023 277}
{Myanmar 2023 480}
{Thailand 2023 393}
{Laos 2023 31}
{Malaysia 2023 968}
{Vietnam 2023 590}
{Brunei 2022 469}
{Malaysia 2022 282}
{Laos 2022 497}
{Cambodia 2022 293}
{Singapore 2022 82}
{Myanmar 2022 354}
{Thailand 2022 803}
{Philippines 2022 913}
{Vietnam 2022 711}
{Indonesia 2022 876}
{Vietnam 2021 29}
{Philippines 2021 152}
{Brunei 2021 333}
{Singapore 2021 676}
{Thailand 2021 361}
{Malaysia 2021 476}
{Laos 2021 108}
{Cambodia 2021 881}
{Myanmar 2021 950}
{Indonesia 2021 543}
*/
