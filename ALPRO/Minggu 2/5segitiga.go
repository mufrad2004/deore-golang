package main

import "fmt"

func main() {
	var (
		s1, s2, s3 int
	)

	fmt.Scan(&s1, &s2, &s3)
	segitiga(s1, s2, s3)
}

func segitiga(s1, s2, s3 int) {

	if ((s1 * s2) > s3) && ((s1 * s3) > s2) && ((s2 * s3) > s1) {
		fmt.Println("segitiga")
	} else {
		fmt.Println("bukan segitiga")
	}
}
