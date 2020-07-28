package main

import (
	"demo15/employee"
)

func main() {
	e := employee.New(
		"Sam",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()
}
