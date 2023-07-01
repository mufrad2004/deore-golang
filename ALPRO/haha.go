package main

import "fmt"

type playlist struct {
	judul, penyanyi string
	menit, detik    int
	totalDetik      int
}

type tabPlaylist [1000]playlist

func isiPlaylist(t *tabPlaylist, n int) {
	var isi playlist
	for i := 0; i < n; i++ {
		fmt.Scan(&isi.judul, &isi.penyanyi, &isi.menit, &isi.detik)
		isi.totalDetik = isi.detik + (isi.menit * 60)
		t[i] = isi
	}
}

func terpendek(t tabPlaylist, n int) string {
	var (
		output string
		pendek playlist
	)

	pendek = t[0]
	for i := 0; i < n; i++ {
		if pendek.totalDetik > t[i].totalDetik {
			pendek = t[i]
		}
	}

	output = pendek.judul
	return output
}

func main() {
	var (
		n   int
		tab tabPlaylist
	)

	fmt.Scan(&n)
	isiPlaylist(&tab, n)
	fmt.Print("*", terpendek(tab, n))
}
