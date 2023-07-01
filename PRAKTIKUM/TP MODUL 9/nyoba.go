package main

import (
	"fmt"
	"strconv"
)

/*
str := "123"
num, err := strconv.Atoi(str)

	if err != nil {
	    fmt.Println("Error:", err)
	    return
	}

fmt.Println(num)
*/
func main() {
	var (
		a     byte
		hasil int
		angka string
	)

	fmt.Scanf("%c", &a)
	for a != '.' {
		angka = string(a)
		if angka[0] >= 48 && angka[0] <= 57 {
			num, err := strconv.Atoi(angka)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			hasil += num
		}
		fmt.Scanf("%c", &a)
	}
	fmt.Print(hasil)
}
