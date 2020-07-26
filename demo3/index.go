package main

import (
	"time"
	"fmt"
)

var count int = 0;

func main(){
	fmt.Println("Begin");

	// Explicit Declaration
	var tmp1 int = 0;
	tmp1 = 10;
	var tmp2 string = "hello";
	var tmp3 bool = true;
	
	/*
	const tmp4 int = 0;
	tmp4 = 10;
	*/

	// Implicit Declration
	// var tmp5 int = 0
	tmp5 := 0
	tmp6 := "codemobiles"


	fmt.Println(tmp1);
	fmt.Println(tmp2);
	fmt.Println(tmp3);
	fmt.Println(tmp5);
	fmt.Println(tmp6);

	count++
	fmt.Println(count);
	run();
	run();
	run();
}

func run()  {
	count++
	fmt.Println(count);
	// fmt.Println(tmp6);
}