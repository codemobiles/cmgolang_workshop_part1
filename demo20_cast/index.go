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
	r := s.(rectangle); // cast from shape to rectangle
	fmt.Printf("Rect  width: %f, height: %f\n", r.width, r.height)
	break
  case "circle":
	c := s.(circle); // cast from shape to circle
	fmt.Printf("Circle radius: %f\n", c.radius)
  }
}

func castToRectangle(s shape)  {
	r, ok := s.(rectangle);
	if !ok {
		fmt.Println("Casting Error")
	}else{
		fmt.Printf("Casting to Rectangle Successfuly w: %f, h: %f", r.width, r.height)
	}
	
}

func main() {
	r := rectangle{width: 10, height: 10}
	c := circle{radius: 10}		

	fmt.Printf("Rectangle: %f\n", getArea(r))
	fmt.Printf("Circle: %f\n", getArea(c))

	showInfo(r)
	showInfo(c)

	castToRectangle(c) // error
	castToRectangle(r) // no error
}