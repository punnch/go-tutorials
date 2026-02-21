package company

type Employee struct {
	ID       int
	FullName string
	Position string
}

func NewEmployee(fullName, position string) Employee {
	return Employee{
		FullName: fullName,
		Position: position,
	}
}
