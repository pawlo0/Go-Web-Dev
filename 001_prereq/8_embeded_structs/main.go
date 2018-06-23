package main

import "fmt"

// vehicle
type vehicle struct {
	doors int
	color string
}

// truck
type truck struct {
	vehicle
	fourWheel bool
}

// sedan
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	colorado := truck{vehicle{2, "blue"}, true}
	cruze := sedan{vehicle{4, "white"}, false}

	fmt.Println(colorado)
	fmt.Println(colorado.color)

	fmt.Println(cruze)
	fmt.Println(cruze.color)

}
