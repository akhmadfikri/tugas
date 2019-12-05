package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}
func main() {
	var Person0 Person
	Person0.FirstName = "Muchson"
	Person0.LastName = "Ibi"
	Person0.Age = 27
	fmt.Println(Person0.FirstName, Person0.LastName, Person0.Age)
}