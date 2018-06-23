package main

import "fmt"

type gator int

type flamingo bool

// greeting
func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

// greeting
func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

// swampCreature
type swampCreature interface {
	greeting()
}

func bayou(s swampCreature) {
	s.greeting()

}

func main() {
	var g1 gator
	var f1 flamingo

	bayou(f1)
	bayou(g1)

}
