package emp

type employee struct {
	FirstName string
	LastName  string
	Age       int
}

//Ctor creates New Emp
func New(firstName string, lastName string, age int) employee {
	return employee{FirstName: firstName, LastName: lastName, Age: age}
}

// Ctor creates old emp
func OldEmp() employee {
	return employee{Age: 100}
}
