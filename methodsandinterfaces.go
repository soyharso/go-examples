package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// No se pueden declarar con el mismo nombre métdodos del apuntador a la estructura ni de la propia estructura, pero aún así existen y son diferentes 
// los metódos de la estructura y los métodos del apuntador a la estructura

/*func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}*/

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	fmt.Println(a)

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	fmt.Printf("Type of f: %T\n", f)

	v := Vertex{3, 4}
	//fmt.Println(v.Yo())
	fmt.Println(v.Abs())
	fmt.Printf("Type of v: %T\n", v)

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
	fmt.Printf("Type of a: %T\n", a)

	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
	fmt.Printf("Type of a: %T\n", a)

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v
	fmt.Println(a.Abs())
	fmt.Printf("Type of a: %T\n", a)

	//fmt.Println(a.Abs())
}
