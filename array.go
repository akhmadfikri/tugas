package main

import "fmt"

func main() {
	//long declaration with value
	var salary_a = map[string] int {"umam": 1000, "iswanul": 2000, "aizat": 4000, "test": 0}
	fmt.Println(salary_a)

	//sort declaration
	salary_b := map [string]int{}
	fmt.Println(salary_b)

	//using make
	var salary_c = make(map[string]int)
	salary_c["doe"] = 7000 //assign value
	fmt.Println(salary_b)
}