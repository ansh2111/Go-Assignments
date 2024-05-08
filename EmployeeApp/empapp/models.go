package empapp

type Employee struct{
	Id int `json:"id"`
	Name string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
	Salary float64 `json:"salary,omitempty"`
}

