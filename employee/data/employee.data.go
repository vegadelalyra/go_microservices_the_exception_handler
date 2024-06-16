package data

var Emp []Employee

func InitEmpData() {
	Emp = []Employee{
		{
			ID:         1,
			Name:       "Employee 1",
			Role:       []int{1, 2, 3},
			Department: "Accounts",
		},
		{
			ID:         2,
			Name:       "Employee 2",
			Role:       []int{4},
			Department: "Sales",
		},
		{
			ID:         3,
			Name:       "Employee 3",
			Role:       []int{1, 2, 4},
			Department: "Marketing",
		},
	}
}