package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID string
	Person
}

func (employee Employee) printInfo() {
	fmt.Printf("员工姓名: %s,年龄: %d, 员工id: %s\n", employee.Name, employee.Age, employee.EmployeeID)
}

func main() {
	em := Employee{
		Person: Person{
			Name: "xxx",
			Age:  10,
		},
		EmployeeID: "12234",
	}
	em.printInfo()
}
