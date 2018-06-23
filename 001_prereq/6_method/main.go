package main

import "fmt"

// person
type person struct {
	fName   string
	lName   string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

func main() {

	p1 := person{"John", "Smith", []string{"pizza", "lasagna", "pasta"}}

	fmt.Println(p1.walk())

}
