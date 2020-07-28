package employee

import (
	"fmt"
)

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

var employeeInstance *employee

func Init(
	firstname string,
	lastname string,
	totalLeaves int,
	leavesTaken int) *employee {
		
		if employeeInstance == nil{
			employeeInstance = &employee{
				FirstName:   firstname,
				LastName:    lastname,
				TotalLeaves: totalLeaves,
				LeavesTaken: leavesTaken}
		}

	return employeeInstance
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
