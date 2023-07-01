package main

import "fmt"

type student struct {
	name                            string
	math, indo, eng, sains, average float64
}

type arrStudent [2048]student

func entryStudent(T *arrStudent, n *int) {
	var (
		isi student
	)

	fmt.Scan(&isi.name)
	for isi.name != "stop" {
		fmt.Scan(&isi.math)
		fmt.Scan(&isi.indo)
		fmt.Scan(&isi.eng)
		fmt.Scan(&isi.sains)
		T[*n] = isi
		*n++
		fmt.Scan(&isi.name)
	}
}

func calculateAverage(T *arrStudent, n int) {
	var (
		rataan float64
	)
	for i := 0; i < n; i++ {
		rataan = (T[i].math + T[i].indo + T[i].eng + T[i].sains) / 4
		T[i].average = rataan
	}
}

func max(T arrStudent, n int, flag string) int {
	var (
		max   student
		hasil int
	)
	max = T[0]

	if flag == "math" {
		for i := 0; i < n; i++ {
			if max.math < T[i].math {
				max = T[i]
				hasil = i
			}
		}
	} else if flag == "ind" {
		for i := 0; i < n; i++ {
			if max.indo < T[i].indo {
				max = T[i]
				hasil = i
			}
		}
	} else if flag == "eng" {
		for i := 0; i < n; i++ {
			if max.eng < T[i].eng {
				max = T[i]
				hasil = i
			}
		}
	} else if flag == "sains" {
		for i := 0; i < n; i++ {
			if max.sains < T[i].sains {
				max = T[i]
				hasil = i
			}
		}
	}
	return hasil
}

func rangking(T *arrStudent, n int) {
	var (
		pass int
		idx  int
		temp student
	)

	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i := pass; i < n; i++ {
			if T[idx].average < T[i].average {
				idx = i
			}
		}
		temp = T[idx]
		T[idx] = T[pass-1]
		T[pass-1] = temp
	}
}

func printTop3(T arrStudent, n int) {
	rangking(&T, n)
	var (
		rank int = 1
	)
	for i := 0; i <= 2; i++ {
		fmt.Println("Rangking ", rank, " : ", T[i].name, " rata-rata ", T[i].average)
		rank++
	}
}

func printMax(T arrStudent, n int) {
	var (
		math, ind, eng, sains int
	)

	math = max(T, n, "math")
	fmt.Println("Matematika tertinggi diraih oleh ", T[math].name)
	ind = max(T, n, "ind")
	fmt.Println("Bahasa Indonesia tertinggi diraih oleh ", T[ind].name)
	eng = max(T, n, "eng")
	fmt.Println("Bahasa Inggris tertinggi diraih oleh ", T[eng].name)
	sains = max(T, n, "sains")
	fmt.Println("Pelajaran IPA/IPS tertinggi diraih oleh ", T[sains].name)

}

func main() {
	var (
		t arrStudent
		n int
		// index int
	)

	entryStudent(&t, &n)
	calculateAverage(&t, n)
	printTop3(t, n)
	printMax(t, n)
}
