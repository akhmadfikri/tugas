package main

import  ("fmt")

var (
	alas int
)

func main() {
	//input
	fmt.Print("masukkan alas: ")
	fmt.Scan(&alas)
	fmt.Print("masukkan tinggi: ")
	fmt.Scan(&tinggi)
	luas = alas * tinggi / 2

	//output
	fmt.Println(luas)
}
