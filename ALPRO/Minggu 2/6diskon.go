package main

import "fmt"

func main() {
	var (
		member bool
		harga int
	) 
	
	fmt.Scan(&harga , &member)
	
	diskon(harga,member)
	
	

}


func diskon(harga int, member bool) {
	var total float64
	if (harga > 100000) && (member == true) {
		total = float64(harga) - (float64(harga) * 0.1)
	}else if (harga > 100000) && (member == false) {
		total = float64(harga) - (float64(harga) * 0.05)
	}else {
		total = float64(harga)
	}
	
	fmt.Println(int(total))
}