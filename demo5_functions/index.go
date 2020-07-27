package main

import (
	"fmt"
)

func main() {
	fn1();
	fn2("Good Morning")
	fn3("Good Morning", 1)
	const a,b = 2,3
	fmt.Printf("%d+%d=%d\n",a,b, sum(a,b))

	x,y := swap(a,b)
	fmt.Printf("%d,%d => %d,%d\n",a,b, x,y)

	_x,_y,title := swapV2(10,20)
	fmt.Printf("%d,%d => %d,%d,%s\n",10,20, _x,_y, title)
}

func fn1(){
	fmt.Println("CodeMobiles");
}

func fn2(msg string){
	fmt.Println(msg);
}

func fn3(title string, version int){
	fmt.Print(title);
	fmt.Println(version);
}

func sum(a int, b int) int{
	return a + b;
}

func swap(a int, b int) (int, int){
	return b,a
}

func swapV2(a, b int) (x,y int, title string){
	y = a
	x = b
	title = "Result"
	return 
}


