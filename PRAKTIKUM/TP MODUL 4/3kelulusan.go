package main

import "fmt"

func cumlaude_1301223029(ipk float64, masaStudi int, publikasi bool) bool {
	// Mengembalikan true apablia memenuhi kriteria "Cumlaude", dan false apabila sebaliknya
	var (
		hasil bool
	)
	hasil = (ipk >= 3.51 && ipk <= 4.00) && (masaStudi <= 8) && (publikasi == true)
	return hasil
}

func sangatMEmuaskan_1301223029(ipk float64, masaStudi int, publikasi bool) bool {
	// Mengembalikan true apablia memenuhi kriteria "Sangat Memuaskan", dan false apabila sebaliknya
	var (
		hasil bool
	)
	hasil = (ipk >= 2.76 && ipk <= 3.5) && (masaStudi <= 14) && (publikasi == true || publikasi == false) || (cumlaude_1301223029(ipk, masaStudi, publikasi) == false && memuaskan_1301223029(ipk, masaStudi, publikasi) == false)
	return hasil
}

func memuaskan_1301223029(ipk float64, masaStudi int, publikasi bool) bool {
	// Mengembalikan true apablia memenuhi kriteria "Memuaskan", dan false apabila sebaliknya
	var (
		hasil bool
	)
	hasil = (ipk >= 2.00 && ipk <= 2.75) && (masaStudi <= 14) && (publikasi == true || publikasi == false)
	return hasil
}

func main() {
	var (
		ipk                  float64
		masaStudi            int
		publikasi, kelulusan bool
		pesan                string
	)

	fmt.Scan(&ipk, &masaStudi, &publikasi)

	kelulusan = cumlaude_1301223029(ipk, masaStudi, publikasi)
	if kelulusan == true {
		pesan = "Cumlaude"
	} else {
		kelulusan = sangatMEmuaskan_1301223029(ipk, masaStudi, publikasi)
		if kelulusan == true {
			pesan = "Sangat Memuaskan"
		} else {
			kelulusan = memuaskan_1301223029(ipk, masaStudi, publikasi)
			if kelulusan == true {
				pesan = "Memuaskan"
			}
		}
	}

	fmt.Print(pesan)

}
