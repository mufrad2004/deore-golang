package main

import "fmt"

const NMAX int = 1000

type wisudawan struct {
	nama, nim      string
	eprt, semester int
	ipk            float64
}

type tabWisudawan [NMAX]wisudawan

func isiData(t *tabWisudawan, N int) {
	var (
		input wisudawan
	)
	fmt.Scan(&input.nim)
	for input.nim != "none" && N <= NMAX {
		fmt.Scan(&input.nama, &input.eprt, &input.semester, &input.ipk)
		t[N] = input
		N++
		fmt.Scan(&input.nim)
	}
}

func ipkTerendah(t *tabWisudawan, N int) float64 {
	var (
		min wisudawan
		i   int = 1
	)

	min = t[0]
	for i < N {
		if min.ipk > t[i].ipk {
			min = t[i]
		}
		i++
	}
	return min.ipk
}

func main() {

}
