package main


import "fmt"

func main(){
	var (
		jam, menit , detik int
	)
	
	fmt.Scan(&jam,&menit,&detik)
	
	konversiWaktu(jam,menit,detik)
}


func konversiWaktu(jam,menit,detik int){
	var (
		total int
	)
	
	jam = jam * 360
	menit = menit * 60
	total = jam + menit + detik
	fmt.Println("Hasil konversi = ",total," detik")
}