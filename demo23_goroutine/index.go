package main

import (
	"time"
	"fmt"
)

func run1()  {
	for i := 0; i < 100; i++ {
		fmt.Println("Run1 something")
	}	
}

func run2()  {
	for i := 0; i < 100; i++ {
		fmt.Println("Run2 something")
	}
}

func main() {
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}