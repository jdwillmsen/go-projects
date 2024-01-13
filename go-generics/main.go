package main

import "fmt"

type Subtractable interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint32 | uint64
}

type Moveable[S Subtractable] interface {
	Move(S)
}

func Move[V Moveable[S], S Subtractable](v V, distance, meters S) S {
	v.Move(meters)
	return Subtract(distance, meters)
}

type Person[S Subtractable] struct {
	Name string
}

func (p Person[S]) Move(meters S) {
	fmt.Printf("%s moved %v meters\n", p.Name, meters)
}

type Car[S Subtractable] struct {
	Name string
}

func (c Car[S]) Move(meters S) {
	fmt.Printf("%s moved %v meters\n", c.Name, meters)
}

func main() {
	p := Person[float64]{Name: "Jake"}
	c := Car[int]{Name: "Corvette"}

	milesToDestination := 100

	distanceLeft := Move(c, milesToDestination, 95)

	fmt.Println("Remaining distance to destination: ", distanceLeft)

	newDistanceLeft := Move(p, float64(distanceLeft), 5)

	fmt.Println("Remaining distance to destination: ", newDistanceLeft)
}

// Subtract will subtract the second value from the first
func Subtract[V Subtractable](a, b V) V {
	return a - b
}
