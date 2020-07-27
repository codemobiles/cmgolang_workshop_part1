package main

import (
	"fmt"
)

func main() {
	fn1();
	fn2("Good Morning")
	fn3("Good Morning", 1)
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


