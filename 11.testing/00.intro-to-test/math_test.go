package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

// Basic test
func TestAdder(t *testing.T) {
	if gotResult := Adder(4, 7, 1, 0); gotResult != 12 {
		t.Errorf("Adder() = %v, want 12", gotResult)
	}
}

// Table test
func TestAdderUsingTableTest(t *testing.T) {
	type args []int

	tests := []struct {
		args       args
		wantResult int
	}{
		{[]int{4, 7, 1, 0}, 12},
		{[]int{2, 2}, 4},
		{[]int{1, 1, 3, 3, 12}, 20},
		{[]int{21, -21}, 0},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestAdderUsingTableTest[%v]", i), func(t *testing.T) {
			if gotResult := Adder(tt.args...); gotResult != tt.wantResult {
				t.Errorf("Adder() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// Blackbox
func TestAdderBlackbox(t *testing.T) {

	f := func(x, y int) bool {
		return Adder(x, y) == x+y
	}

	err := quick.Check(f, nil)
	if err != nil {
		t.Fatal(err)
	}
}

// Benchmarking (test performance)
func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adder(4, 7)
	}
}

// Provide examples for godoc
func ExampleAdder() {
	fmt.Println(Adder(4, 7))
	// Output:
	// 11
}

func ExampleAdder_multiple() {
	fmt.Println(Adder(3, 6, 7, 4, 61))
	// Output:
	// 81
}
