package main

import "fmt"

func pola_1301223029(n int, s *string) {
	if n == 1 {
		*s = "*"
		fmt.Print(*s, "\n")
	} else {
		pola_1301223029(n-1, &*s)
		for i := 1; i <= n; i++ {
			*s = "*"
			fmt.Print(*s)
		}
		fmt.Print("\n")
	}
}

func main() {
	var (
		angka int
		s     string = ""
	)
	fmt.Scan(&angka)
	pola_1301223029(angka, &s)

}
