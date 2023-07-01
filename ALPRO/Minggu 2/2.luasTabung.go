package main

import "fmt"

func main(){
	var (
		r,t float64
	)
	
	fmt.Scan(&r ,&t)
	luas_tabung(r,t)
}


func luas_tabung(r,t float64) {
	var (
		pi float64 = 3.14
		luas float64
	)
	
	luas = 2* (pi * r) * (r + t)
	fmt.Println(luas)
	
}