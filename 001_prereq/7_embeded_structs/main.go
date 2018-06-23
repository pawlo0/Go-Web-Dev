package main


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
    
}