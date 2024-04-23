package main

type EmployeeBuilder struct {
	Name     string
	Position string
	Age      int
	Salary   int
	Address  string
}

func NewEmployeeBuilder() *EmployeeBuilder {

	return &EmployeeBuilder{}
}

func (eb *EmployeeBuilder) setName(name string) {
	eb.Name = name
}

func (eb *EmployeeBuilder) setPosition(position string) {
	eb.Position = position
}

func (eb *EmployeeBuilder) setAge(age int) {
	eb.Age = age
}

func (eb *EmployeeBuilder) setSalary(salary int) {
	eb.Salary = salary
}

func (eb *EmployeeBuilder) setAddress(address string) {
	eb.Address = address
}

func (eb *EmployeeBuilder) build() *Employee {

	return &Employee{
		Name:     eb.Name,
		Position: eb.Position,
		Age:      eb.Age,
		Address:  eb.Address,
		Salary:   eb.Salary,
	}

}
