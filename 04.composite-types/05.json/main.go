package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

type employee struct {
	Person   person `json:"Person"`
	Position string `json:"Position"`
	Salary   int    `json:"Salary"`
}

func main() {

	ceo := employee{
		Person: person{
			FirstName: "Steven",
			LastName:  "King",
			Age:       50,
		},
		Position: "CEO",
		Salary:   1_000_000.,
	}

	cto := employee{
		Person: person{
			FirstName: "Kyle",
			LastName:  "Cross",
			Age:       35,
		},
		Position: "CTO",
		Salary:   100_000.,
	}

	employees := []employee{ceo, cto}

	fmt.Println("employees:", employees)

	// Marshal
	jsonList, _ := json.Marshal(employees)
	// NOTE: Ignoring errors
	fmt.Println(string(jsonList))

	// Unmarshal
	s := `[{"Person":{"FirstName":"Steven","LastName":"King","Age":50},"Position":"CEO","Salary":1000000},{"Person":{"FirstName":"Kyle","LastName":"Cross","Age":35},"Position":"CTO","Salary":100000}]`
	bs := []byte(s)

	var emps []employee
	json.Unmarshal(bs, &emps)
	fmt.Println("emps:", emps)
}
