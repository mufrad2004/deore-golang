package main


import "fmt"


func main() {
	var(
		angka1,angka2 int
	)
	fmt.Scan(&angka1,&angka2)
	fmt.Print(permutasi1(angka1), permutasi1(angka2), permutasi2(angka1,angka2))
}

func permutasi1(x int) int {
	/*Mengembalikan nilai permutasi bertipe integer*/
	var (
		total int = 1
	)
	for i := 1; i <= x ; i++ {
		total *= i
	}
	
	return total
}

func permutasi2(x,y int) int{
	/*Fungsi menghitung x permutasi y jika x >= y, dan y permutasi x jika x < y*/
	var (
		total int
	)
	
	if x >= y {
		total = permutasi1(x)/ permutasi1(x-y)
	}else {
		total = permutasi1(y)/ permutasi1(y-x)
	}
	
	return total
}