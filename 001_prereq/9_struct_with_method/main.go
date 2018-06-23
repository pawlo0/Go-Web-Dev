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

func (t truck) transportationDevice() string {
	return fmt.Sprintln("I am a", t.color, "truck with", t.doors, "doors and I am durable.")
}

// sedan
type sedan struct {
	vehicle
	luxury bool
}

func (s sedan) transportationDevice() string {
	if s.luxury {
		return fmt.Sprintln("I am a sedan with", s.doors, "doors and I am luxurious.")
	}
	return fmt.Sprintln("The", s.color, s.doors, "doors sedan is crossing the city")
}

func main() {
	t1 := truck{vehicle{2, "blue"}, true}
	s1 := sedan{vehicle{4, "white"}, false}

	fmt.Println(t1.transportationDevice())

	fmt.Println(s1.transportationDevice())

}
