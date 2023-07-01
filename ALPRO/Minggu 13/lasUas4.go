package main

import "fmt"

const b int = 6
const k int = 13

type typeOne [k]int
type typeTwo [b,b]byte

func inputData(t typeOne) {
	var x int = 1
	for x <= k {
		fmt.Scan(&t[x])
		x++
	}
}

func inputMat(t typeTwo) {
	var i, j int
	i = 1
	for i <= b {
		j = 1
		for j <= b {
			fmt.Scan(&t[i, j])
		}
	}
}

func main() {

}
