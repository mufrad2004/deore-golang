package main

import (
	"fmt"
	"math"
)

func membeliK(uangA float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = uangA / 4
	*tUang = *tUang - *tPengeluaran
}

func membeliBJ(uangA float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = uangA / 20
	*tUang = *tUang - *tPengeluaran
}

func membuatPB(uangA float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = uangA / 200
	*tUang = *tUang - *tPengeluaran
}

func menjahitB(uangA float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = uangA / 200
	*tUang = *tUang - *tPengeluaran
}

func mengemasB(uangA float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = uangA * 3 / 200
	*tUang = *tUang - *tPengeluaran
}

func distribusi(uangA float64, tPengeluaran, tUang, tPemasukan *float64) {
	*tPengeluaran = uangA * 3 / 200
	*tPemasukan = uangA / 2
	*tUang = *tUang + *tPemasukan - *tPengeluaran
}

func main() {
	var modal, pengeluaran, pemasukan, uang, totp float64
	fmt.Scan(&modal)
	uang = modal
	membeliK(modal, &uang, &pengeluaran)
	totp = totp + pengeluaran
	membeliBJ(modal, &uang, &pengeluaran)
	totp = totp + pengeluaran
	for i := 1; i <= 2; i++ {
		membuatPB(modal, &uang, &pengeluaran)
		totp = totp + pengeluaran
	}
	for i := 1; i <= 4; i++ {
		menjahitB(modal, &uang, &pengeluaran)
		totp = totp + pengeluaran
	}
	for i := 1; i <= 2; i++ {
		mengemasB(modal, &uang, &pengeluaran)
		totp = totp + pengeluaran
	}
	distribusi(modal, &pengeluaran, &uang, &pemasukan)
	totp = totp + pengeluaran
	fmt.Println(math.Round(totp), pemasukan, math.Round(uang))

}
