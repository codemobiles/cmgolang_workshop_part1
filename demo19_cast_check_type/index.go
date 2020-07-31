package main

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64{
	return s.area()
}

func showInfo(s shape){
  t := reflect.TypeOf(s).Name()
  switch t {
  case "rectangle":
	r := s.(rectangle);
	fmt.Printf("Rect  width: %f, height: %f\n", r.width, r.height)
	break
  case "circle":
	c := s.(circle);
	fmt.Printf("Circle radius: %f\n", c.radius)
  }

}

func main() {
	r := rectangle{width: 10, height: 10}
	c := circle{radius: 10}		

	fmt.Printf("Rectangle: %f\n", getArea(r))
	fmt.Printf("Circle: %f\n", getArea(c))

	showInfo(r)
	showInfo(c)
}