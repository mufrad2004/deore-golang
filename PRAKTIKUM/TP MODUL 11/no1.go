package main

import (
	"fmt"
)

const NMAX int = 1000

type himpunan struct {
	info    [NMAX]string
	nElemen int
}

func createSet_1301223029(set *himpunan) {
	var (
		input string
		count int
	)
	for i := 0; i < NMAX; i++ {
		fmt.Scan(&input)
		for j := 0; j < len(set.info); j++ {
			if input == set.info[j] {
				count++
			}
		}
		if count > 0 {
			break
		}
		set.info[i] = input
		set.nElemen++
	}
}

func printSet_1301223029(set himpunan) {
	for i := 0; i < len(set.info); i++ {
		if set.info[i] != "" {
			fmt.Print(set.info[i], " ")
		}
	}
	fmt.Print("\n")
}

func isMember_1301223029(set himpunan, s string) bool {
	var (
		hasil int
	)
	for i := 0; i < len(set.info); i++ {
		if s == set.info[i] {
			hasil++
		}
	}

	if hasil > 0 {
		return true
	} else {
		return false
	}
}

func interserction_1301223029(set1, set2 himpunan, set3 *himpunan) {
	for i := 0; i < len(set1.info); i++ {
		for j := 0; j < len(set2.info); j++ {
			if set1.info[i] == set2.info[j] {
				duplikat := false
				for k := 0; k < len(set3.info); k++ {
					if set1.info[i] == set3.info[k] {
						duplikat = true
						break
					}
				}
				if !duplikat {
					*&set3.info[i] = set1.info[i]
				}
			}
		}
	}
}

func union_1301223029(set1, set2 himpunan, set3 *himpunan) {
	set3.nElemen = 0
	for i := 0; i < set1.nElemen; i++ {
		if !isMember_1301223029(*set3, set1.info[i]) {
			set3.info[set3.nElemen] = set1.info[i]
			set3.nElemen++
		}
	}
	for j := 0; j < set2.nElemen; j++ {
		if !isMember_1301223029(*set3, set2.info[j]) {
			set3.info[set3.nElemen] = set2.info[j]
			set3.nElemen++
		}
	}
}
func main() {
	var (
		A, B, C, D himpunan
	)
	createSet_1301223029(&A)
	createSet_1301223029(&B)
	fmt.Println("Irisan A dan B : ")
	interserction_1301223029(A, B, &C)
	for i := 0; i < len(C.info); i++ {
		if C.info[i] != "" {
			fmt.Print(C.info[i], " ")
		}
	}
	fmt.Print("\n\n")

	fmt.Println("Union A dan B : ")
	union_1301223029(A, B, &D)
	for i := 0; i < D.nElemen; i++ {
		fmt.Print(D.info[i], " ")
	}
}
