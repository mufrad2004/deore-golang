package main

import "fmt"

func main(){
	var (
		n int
		celcius float64
	)
	
	fmt.Scan(&n)
	
	for i:= 1 ; i<= n ; i++ {
		fmt.Scan(&celcius)
		temperatur(celcius)
	}
}

func temperatur(celcius float64){
	var (
		fahrenheit float64
	)
	
	fahrenheit = ((9.0/5.0) * celcius) + 32.0
	fmt.Println(celcius, "Celcius = ", fahrenheit , "Fahrenheit" )
}