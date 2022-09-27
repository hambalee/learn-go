package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {

	employee1 := employee{
		employeeID:   "101",
		employeeName: "abc",
		phone:        "123",
	}
	fmt.Println("employee1 =", employee1) // employee1 = {101 abc 123}

	employeeList := [3]employee{}
	employeeList[0] = employee{
		employeeID:   "101",
		employeeName: "abc",
		phone:        "123",
	}
	employeeList[1] = employee{
		employeeID:   "102",
		employeeName: "def",
		phone:        "456",
	}
	employeeList[2] = employee{
		employeeID:   "103",
		employeeName: "ghi",
		phone:        "789",
	}

	fmt.Println("employeeList =", employeeList) // employeeList = [{101 abc 123} {102 def 456} {103 ghi 789}]

	//struct with slice
	employeeList2 := []employee{}
	employee2 := employee{
		employeeID:   "101",
		employeeName: "abc",
		phone:        "123",
	}
	employee3 := employee{
		employeeID:   "102",
		employeeName: "def",
		phone:        "456",
	}
	employee4 := employee{
		employeeID:   "103",
		employeeName: "ghi",
		phone:        "789",
	}

	employeeList2 = append(employeeList2, employee2)
	employeeList2 = append(employeeList2, employee3)
	employeeList2 = append(employeeList2, employee4)
	fmt.Println("employeeList2 =", employeeList2) // employeeList2 = [{101 abc 123} {102 def 456} {103 ghi 789}]
}
