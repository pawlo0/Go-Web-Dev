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

func (t truck) transportationDevice() string {
	return fmt.Sprintln("I am a", t.color, "truck with", t.doors, "doors and I am durable.")
}

func (s sedan) transportationDevice() string {
	if s.luxury {
		return fmt.Sprintln("I am a sedan with", s.doors, "doors and I am luxurious.")
	}
	return fmt.Sprintln("The", s.color, s.doors, "doors sedan is crossing the city")
}

// transportation
type transportation interface {
	transportationDevice() string
}

func report(r transportation) {
	fmt.Println(r.transportationDevice())
}

func main() {
	t1 := truck{vehicle{2, "blue"}, true}
	s1 := sedan{vehicle{4, "white"}, false}

	report(t1)
	report(s1)

}
