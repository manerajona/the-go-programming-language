package godoc

import (
	"fmt"
	"log"
	"testing"
)

func TestGetOne(t *testing.T) {
	expect := Usr{
		ID:       42,
		Username: "mrobot",
	}
	users = []Usr{expect}

	got, err := getOne(expect.ID)

	if err != nil {
		t.Fatal(err)
	}
	if got != expect {
		t.Errorf("did not get expected user. Got %+v, expected %+v", got, expect)
	}
}

func ExampleGetOne() {
	users = []Usr{
		{ID: 1, Username: "mrobot"},
	}
	u, err := getOne(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(u.ID, u.Username)

	// Output:
	// 1 mrobot
}
