package main

// right click to see more go quick menu
// cmd+shift+p
// ctrl + spacebar
import (
	"demo5/helper"
	"fmt"
)

func main() {
	fmt.Println("CodeMobiles");
	helper.Test();
	helper.Count++
	helper.Test();
	helper.Count++
	helper.Test();
}