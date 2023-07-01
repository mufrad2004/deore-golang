package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		modalAwal                            float64
		tUang, tPengeluaran, tPemasukan, toP float64
	)

	fmt.Scan(&modalAwal)
	tUang = modalAwal
	membeliKain(modalAwal, &tUang, &tPengeluaran)
	toP += tPengeluaran
	membeliBenangJahit(modalAwal, &tUang, &tPengeluaran)
	toP += tPengeluaran
	for i := 1; i <= 2; i++ {
		membuatPolaBaju(modalAwal, &tUang, &tPengeluaran)
		toP += tPengeluaran
	}
	for i := 1; i <= 4; i++ {
		menjahitBaju(modalAwal, &tUang, &tPengeluaran)
		toP += tPengeluaran
	}
	for i := 1; i <= 2; i++ {
		mengemasBaju(modalAwal, &tUang, &tPengeluaran)
		toP += tPengeluaran
	}
	mendistirbusikan(modalAwal, &tUang, &tPemasukan, &tPengeluaran)
	toP += tPengeluaran
	fmt.Println(math.Round(toP), tPemasukan, math.Round(tUang))

	// fmt.Print(tPengeluaran, tPemasukan, tUang)

}

func membeliKain(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = (uangAwal / 4)
	*tUang = *tUang - *tPengeluaran
}

func membeliBenangJahit(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = (uangAwal / 20)
	*tUang = *tUang - *tPengeluaran
}

func membuatPolaBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = (uangAwal / 200)
	*tUang = *tUang - *tPengeluaran
}

func menjahitBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = (uangAwal / 200)
	*tUang = *tUang - *tPengeluaran
}

func mengemasBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = ((uangAwal * 3) / 200)
	*tUang = *tUang - *tPengeluaran
}

func mendistirbusikan(uangAwal float64, tUang, tPemasukan, tPengeluaran *float64) {
	*tPengeluaran = ((uangAwal * 3) / 200)
	*tPemasukan = uangAwal / 2
	*tUang = *tUang + *tPemasukan - *tPengeluaran
}
