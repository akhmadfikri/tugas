package main

import "fmt"

func main() {
	var bil int
	var i int
fmt.Println("masukkan bilangan")
fmt.Scanln(&bil)
for i=1; i<=bil;i++{
	if bil % i == 0 {
		fmt.Println(i)
	}
}
}
