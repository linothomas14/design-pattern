package main

import "fmt"

func main() {

	bldr := NewEmployeeBuilder()

	bldr.setName("Joe")
	bldr.setPosition("Frontend Engineer")

	employee1 := bldr.build()

	fmt.Println(employee1)

	bldr2 := NewEmployeeBuilder()

	bldr2.setName("Yulyano")
	bldr2.setPosition("Backend Engineer")
	bldr2.setAddress("Bogor")
	bldr2.setSalary(99999999)
	bldr2.setAge(22)

	employee2 := bldr2.build()

	fmt.Println(employee2)
}
