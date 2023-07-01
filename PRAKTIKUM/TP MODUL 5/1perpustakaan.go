package main

import "fmt"

func main() {
	var (
		tgl1, bln1, thn1 int
		tgl2, bln2, thn2 int
	)
	inputTglPinjam_1301223029(&tgl1, &bln1, &thn1)
	hitungTanggalKembali_1301223029(tgl1, bln1, thn1, &tgl2, &bln2, &thn2)
	if valid_1301223029(tgl1, bln1, thn1) {
		fmt.Print(tgl2, bln2, thn2)
	} else {
		fmt.Print("Input tidak valid")
	}

}

func inputTglPinjam_1301223029(tanggal, bulan, tahun *int) {
	var (
		t, g, y int
	)
	/*
		I.S memasukkan tanggal , bulan , tahun dalam int
		F.S tanggal, bulan, tahun telah memiliki nilai
	*/
	fmt.Scan(&t, &g, &y)
	*tanggal = t
	*bulan = g
	*tahun = y

}

func valid_1301223029(tanggal, bulan, tahun int) bool {
	//mengirimkan true jika tahun > 0 , bulan pada range 1..12, dan  0< tanggal 1< jumlah hari dalam bulan tsb
	var (
		arrBulan []int
	)

	arrBulan = append(arrBulan, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31)
	if (tahun > 0) && (bulan >= 1 && bulan <= 12) && (tanggal <= arrBulan[bulan-1]) {
		return true
	} else {
		return false
	}
}

func kabisat_1301223029(tahun int) bool {
	//Mengirimkan true jika tahun kabisat atau false jika bukan kabisat
	if tahun%400 == 0 {
		return true
	} else if tahun%100 == 0 {
		return false
	} else if tahun%4 == 0 {
		return true
	} else {
		return false
	}
}

func getJumlahHari_1301223029(bulan1, tahun1 int, jmlHari *int) {
	/*
		I.S Diketahui bulan dan tahun
		F.S jmlHari adalah jumlah hari dari bulan dan tahun tersebut
	*/
	var (
		arrBulan []int
	)

	arrBulan = append(arrBulan, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31)
	for i := 1; i <= len(arrBulan); i++ {
		if i == bulan1 {
			if kabisat_1301223029(tahun1) {
				*jmlHari = arrBulan[bulan1-1]
			} else {
				if bulan1 == 2 {
					*jmlHari = 28
				} else {
					*jmlHari = arrBulan[bulan1-1]
				}
			}
		}
	}
}

func hitungTanggalKembali_1301223029(tanggal1, bulan1, tahun1 int, tanggal2, bulan2, tahun2 *int) {
	var (
		jumlahHari int
	)
	/*
		I.S tanggal1,bulan1, tahun1 valid
		F.S tanggal2,bulan2,tahun2 adalah 3 har iberikutnya dari tanggal 1, bulan 1, tahun1
	*/
	*tanggal2 = tanggal1 + 3
	*bulan2 = bulan1
	*tahun2 = tahun1
	getJumlahHari_1301223029(bulan1, tahun1, &jumlahHari)
	if *tanggal2 > jumlahHari {
		*tanggal2 = *tanggal2 - jumlahHari
		*bulan2++
		if *bulan2 > 12 {
			*bulan2 = 1
			*tahun2++
		}
	}
}
