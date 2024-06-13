package main

import (
	"fmt"
	"math"
)

type Vector3D struct {
	x, y, z float64
}

type Vector2D struct {
	x, y float64
}

type Number float64

type Euclid interface {
	Norm() float64
}

func (v Vector3D) Norm() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector2D) Norm() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (n Number) Norm() float64 {
	if n > 0 {
		return float64(n)
	}
	return -float64(n)
}

func main() {
	var v Euclid
	v = Vector3D{1, 2, 3}
	fmt.Println(v.Norm())
	v = Vector2D{1, 2}
	fmt.Println(v.Norm())
	v = Number(-2.5)
	fmt.Println(v.Norm())
}
