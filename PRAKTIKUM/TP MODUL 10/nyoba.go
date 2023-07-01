package main

import (
	"fmt"
)

func cek(a byte) int {
	var hasil int
	if int(a) >= int('a') && int(a) <= int('z') {
		for i := int('A'); i <= int('Z'); i++ {
			if int(a)-32 == i {
				hasil = i
			}
		}
	} else if int(a) >= int('A') && int(a) <= int('Z') {
		for i := int('a'); i <= int('z'); i++ {
			if int(a)+32 == i {
				hasil = i
			}
		}
	}
	return hasil
}

func main() {
	var (
		a, b           byte
		hasil1, hasil2 int
	)
	fmt.Scanf("%c", &a)
	fmt.Scanf("%c", &b)

	hasil1 = cek(a)
	hasil2 = cek(b)
	fmt.Print(hasil1, hasil2)
}

/*
package main

import (
	"fmt"
)

func main() {
	var (
		a     byte
		hasil uint8
	)

	fmt.Scanf("%c", &a)
	if int(a) >= 97 && int(a) <= 122 {
		for i := 65; i <= 90; i++ {
			if int(a)-32 == i {
				hasil = uint8(i)
				break
			}
		}
	} else if int(a) >= 65 && int(a) <= 90 {
		for i := 97; i <= 122; i++ {
			if int(a)+32 == i {
				hasil = uint8(i)
				break
			}
		}
	}
	fmt.Print(string(hasil))

}
*/
