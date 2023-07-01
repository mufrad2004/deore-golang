package main

import "fmt"

const nmax int = 2023

type pokemon struct {
	nama                       string
	cp, hp, ivAtk, ivDef, ivHp int
}

type arrPokemon [nmax]pokemon

func inputPokemon_1301223029(t *arrPokemon, n *int) {
	var isi pokemon
	fmt.Scanln(&isi.nama, &isi.cp, &isi.hp, &isi.ivAtk, &isi.ivDef, &isi.ivHp)
	t[*n] = isi
	*n++
}

func printPokemon_1301223029(t arrPokemon, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(t[i])
	}
}

func mengurutkanPokemon_1301223029(t *arrPokemon, n *int, flag string) {
	var (
		i     int
		merah pokemon
	)

	if flag == "CP" {
		for pass := 1; pass < *n; pass++ {
			i = pass
			merah = t[i]
			for i > 0 && merah.cp > t[i-1].cp {
				t[i] = t[i-1]
				i--
			}
			t[i] = merah
		}
	} else if flag == "HP" {
		for pass := 1; pass < *n; pass++ {
			i = pass
			merah = t[i]
			for i > 0 && merah.hp > t[i-1].hp {
				t[i] = t[i-1]
				i--
			}
			t[i] = merah
		}
	} else if flag == "IV_atk" {
		for pass := 1; pass < *n; pass++ {
			i = pass
			merah = t[i]
			for i > 0 && merah.ivAtk > t[i-1].ivAtk {
				t[i] = t[i-1]
				i--
			}
			t[i] = merah
		}
	} else if flag == "IV_Def" {
		for pass := 1; pass < *n; pass++ {
			i = pass
			merah = t[i]
			for i > 0 && merah.ivDef > t[i-1].ivDef {
				t[i] = t[i-1]
				i--
			}
			t[i] = merah
		}
	} else {
		for pass := 1; pass < *n; pass++ {
			i = pass
			merah = t[i]
			for i > 0 && merah.ivHp > t[i-1].ivHp {
				t[i] = t[i-1]
				i--
			}
			t[i] = merah
		}
	}
}

func totaalIV_1301223029(t arrPokemon, n int, namaX string) float64 {
	var (
		total float64 = -9999
		j     int
	)
	for i := 0; i < n; i++ {
		if namaX == t[i].nama {
			j = i
			total = ((float64(t[j].ivAtk) + float64(t[j].ivDef) + float64(t[j].ivHp)) * 100.0) / 45.0
		}
	}
	return total
}

func main() {
	var nama, flag string
	var t arrPokemon
	var n, jml int

	fmt.Scanln(&jml)
	for i := 1; i <= jml; i++ {
		inputPokemon_1301223029(&t, &n)
	}
	fmt.Scanln(&nama, &flag)

	printPokemon_1301223029(t, n)
	mengurutkanPokemon_1301223029(&t, &n, flag)
	printPokemon_1301223029(t, n)
	fmt.Printf("Total IV %s adalah %.4f", nama, totaalIV_1301223029(t, n, nama))
}
