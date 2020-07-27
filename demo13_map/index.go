package main

import (
	"fmt"
)

func main() {
   var numbers = map[string]int{"one":1, "two":2, "three":3}
   fmt.Println("",numbers["two"])

   var colors = make(map[string]string)
   colors["red"] = "#f00";
   colors["green"] = "#0f0";
   colors["blue"] = "#00f";
   fmt.Println("",colors)
   fmt.Println("",colors["green"])
}