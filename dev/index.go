package main

// right click to see more go quick menu
// cmd+shift+p
// ctrl + spacebar
import (
	"demo5/helper"
	"fmt"
)

func fnFor(){
	for i := 0; i < 10; i++ {
		fmt.Println("ok")
	}
}

func fnWhile(){
	
	i := 0
	for i < 20 {
		i++;
		fmt.Printf("ok %d\n", i)
	}
}

func fnRange(){
	a := []string{"angular", "nodejs", "golang"}
	for index, s := range a{
		fmt.Println(index,s)
	}
}

func main() {
	fmt.Println("CodeMobiles");
	helper.Test();
	helper.Count++
	helper.Test();
	helper.Count++
	helper.Test();
	fnFor();
	fnWhile();
	fnRange();
}