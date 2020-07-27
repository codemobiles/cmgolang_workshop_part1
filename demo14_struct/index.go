package main

import (
	"fmt"
)

func main() {
   var p1 product;
   p1.name = "Arduino"
   p1.price = 100
   p1.stock = 20
   show(p1)
   update(&p1)
   show(p1)

   // p1 = p1.clear()
   p1 = p1.setDiscount(1)
   show(p1)
}

type product struct {
   name string;
   price int;
   stock int;
}

func (p product) clear() product{
   p.price = 0 
   p.stock = 0
   return p
}

func (p product) setDiscount(d int) product {
   p.price = p.price - d
   return p
}


func show(p product){
   fmt.Println(p)
}

func update(p *product)  {
   p.price = p.price + 100
   p.stock = 10   
}