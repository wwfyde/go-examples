package main

import (
	"log"
	"math"
)

type Point struct {
	X, Y float64
}

type Circle struct {
	Point
	Radius float64
}

func (p *Circle) Area() (area float64) {
	area = math.Pi * float64(p.Radius) * float64(p.Radius)
	return

}

func main() {
	circleA := Circle{Radius: 2}

	log.Println(int(circleA.Area()))

}
