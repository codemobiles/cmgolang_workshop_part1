package main

import (
	"fmt"
)

func main() {
	courses := []string {"Android", "iOS", "React"};

	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1, item)
	}

	for _, item := range courses {
		fmt.Printf("%s\n",item)
	}
}

