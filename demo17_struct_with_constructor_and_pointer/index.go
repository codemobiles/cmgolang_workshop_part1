package main

import (
	"demo15/employee"
)

func main() {
	e := employee.Init(
		"Lek",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()

	e = employee.Init(
		"Kan",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()
}
