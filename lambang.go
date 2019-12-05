package main

import "fmt"

func main() {
	sentence := "hello"

	var _ int = 0

	for _, char := range sentence {
	fmt.Println (string(char))
}
}
