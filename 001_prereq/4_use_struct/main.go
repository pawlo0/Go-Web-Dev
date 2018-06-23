package main

import "fmt"

// person first and last name
type person struct {
    fName string
    lName string
}

func main() {
    
    p1 := person{"John", "Smith"}
    
    fmt.Println(p1)
    fmt.Println(p1.fName)
        
        
}