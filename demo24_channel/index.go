package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1 // send
	fmt.Println("step1")
	fmt.Println(<-ch)


	ch <- 2 // send
	fmt.Println("step2")
	fmt.Println(<-ch)
	
	time.Sleep(1*time.Second)

}