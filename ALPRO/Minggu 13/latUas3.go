package main

import "fmt"

const nmax int = 10013

type account struct {
	nama    string
	accId   int
	balance float64
}

type accounts struct {
	accList [nmax]account
	n       int
}

func loadAccounts(t *accounts) {
	var (
		isi account
	)
	fmt.Scanln(&isi.nama, &isi.accId, &isi.balance)
	for isi.nama != "-1" && isi.accId != -1 && isi.balance != -1 {
		t.accList[t.n] = isi
		t.n++
		fmt.Scanln(&isi.nama, &isi.accId, &isi.balance)
	}
}

func sortAccounts(t *accounts) {
	var (
		idx  int
		temp account
	)
	for pass := 1; pass < t.n; pass++ {
		idx = pass - 1
		for i := pass; i < t.n; i++ {
			if t.accList[idx].nama > t.accList[i].nama {
				idx = i
			}
		}
		temp = t.accList[pass-1]
		t.accList[pass-1] = t.accList[idx]
		t.accList[idx] = temp
	}
}

func displayAccountByName(t accounts, name string) {
	var (
		kiri, kanan, tengah int
		hasil               bool = false
	)
	sortAccounts(&t)
	kiri = 0
	kanan = t.n
	for kiri <= kanan && !hasil {
		tengah = (kiri + kanan) / 2
		if name < t.accList[tengah].nama {
			kanan = tengah - 1
		} else if name > t.accList[tengah].nama {
			kiri = tengah + 1
		} else if name == t.accList[tengah].nama {
			hasil = true
			fmt.Println(t.accList[tengah])
		} else {
			fmt.Println("NO ACCOUNT")
		}
	}
}

func main() {
	var (
		t    accounts
		nama string
	)

	loadAccounts(&t)
	sortAccounts(&t)
	fmt.Scan(&nama)
	for nama != "DONE" {
		displayAccountByName(t, nama)
		fmt.Scan(&nama)
	}
}
