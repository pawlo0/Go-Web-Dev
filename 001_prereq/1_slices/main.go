package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7}

	fmt.Println(a)

	for i := range a {
		fmt.Println(i)
	}

	for i, num := range a {
		fmt.Println(i, num)
	}

}
