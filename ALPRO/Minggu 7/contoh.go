package main

import "fmt"

func main() {
	var (
		days [7]string
		i    int
	)

	days[0] = "senin"
	fmt.Scan(&days[1])
	for i = 2; i <= 6; i++ {
		fmt.Scan(&days[i])
	}

	fmt.Println(days[1], days[5])
}
