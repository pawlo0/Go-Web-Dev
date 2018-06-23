package main

import "fmt"

func main() {
    m := map[string]int{
        "John": 23,
        "Paul": 25,
        "Marie": 21,
    }
    
    fmt.Println(m)
    
    for name := range m {
        fmt.Println(name)
    }
    
    for name, age := range m {
        fmt.Println(name, age)
    }   
}
