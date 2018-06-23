package main

import "fmt"

// person 
type person struct {
    fName string
    lName string
    favFood []string
}

func main() {
    
    p1 := person{"John", "Smith", []string{"pizza", "lasagna", "pasta"}}
    
    for _, v := range p1.favFood {
        fmt.Println(v)
            
    }
}