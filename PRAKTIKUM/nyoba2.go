package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 5}
	arr2 := []int{3, 4, 5, 6, 7}

	intersection := []int{}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				// Cek apakah nilai sudah ada di slice intersection
				isDuplicate := false
				for k := 0; k < len(intersection); k++ {
					if arr1[i] == intersection[k] {
						isDuplicate = true
						break
					}
				}
				// Jika nilai belum ada di slice intersection, tambahkan ke dalam slice
				if !isDuplicate {
					intersection = append(intersection, arr1[i])
				}
			}
		}
	}

	fmt.Println("Intersection of", arr1, "and", arr2, "is:", intersection)
}
