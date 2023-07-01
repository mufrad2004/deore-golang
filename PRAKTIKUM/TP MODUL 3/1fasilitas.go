package main

import "fmt"

func main() {
	var (
		total_belanja     int
		status_membership string
	)

	fmt.Scan(&total_belanja, &status_membership)
	if status_membership == "Yes" {
		if total_belanja > 100000 {
			fmt.Print("Anda dapat discount 5%, tambahan 10 poin dan free gift")
		} else if total_belanja > 75000 {
			fmt.Print("Anda dapat discount 5% , dan tambahan poin 5 ")
		} else {
			fmt.Print("Anda dapat tambahan 5 poin")
		}
	}

}
