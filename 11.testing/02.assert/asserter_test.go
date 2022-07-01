package asserter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	empty       = ""
	zero        = 0
	one         = 1
	emailRegexp = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

func TestAssertActual(t *testing.T) {
	type args struct {
		a      ActualAssertion
		actual interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"assertActual_NotEmpty_success", args{assert.NotEmpty, "hello"}, false},
		{"assertActual_NotEmpty_fail", args{assert.NotEmpty, empty}, true},
		{"assertActual_Empty_success", args{assert.Empty, empty}, false},
		{"assertActual_Empty_fail", args{assert.Empty, "hello"}, true},
		{"assertActual_NotNil_success", args{assert.NotNil, empty}, false},
		{"assertActual_NotNil_fail", args{assert.NotNil, nil}, true},
		{"assertActual_Nil_success", args{assert.Nil, nil}, false},
		{"assertActual_Nil_fail", args{assert.Nil, empty}, true},
		{"assertActual_NotZero_success", args{assert.NotZero, one}, false},
		{"assertActual_NotZero_fail", args{assert.NotZero, zero}, true},
		{"assertActual_Zero_success", args{assert.Zero, zero}, false},
		{"assertActual_Zero_fail", args{assert.Zero, one}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AssertActual(tt.args.a, tt.args.actual); (err != nil) != tt.wantErr {
				t.Errorf("AssertActual() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAssertExpectedAndActual(t *testing.T) {
	type args struct {
		a        ExpectedAndActualAssertion
		expected interface{}
		actual   interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"assertExpectedAndActual_Equal_success", args{assert.Equal, "hello", "hello"}, false},
		{"assertExpectedAndActual_Equal_fail", args{assert.Equal, "hello", "Hello"}, true},
		{"assertExpectedAndActual_NotEqual_success", args{assert.NotEqual, "hello", "Hello"}, false},
		{"assertExpectedAndActual_NotEqual_fail", args{assert.NotEqual, "hello", "hello"}, true},
		{"assertExpectedAndActual_Contains_success", args{assert.Contains, "hello world", "hello"}, false},
		{"assertExpectedAndActual_Contains_fail", args{assert.Contains, "hello world", "hi"}, true},
		{"assertExpectedAndActual_Greater_success", args{assert.Greater, one, zero}, false},
		{"assertExpectedAndActual_Greater_fail", args{assert.Greater, zero, one}, true},
		{"assertExpectedAndActual_Less_success", args{assert.Less, zero, one}, false},
		{"assertExpectedAndActual_Less_fail", args{assert.Less, one, zero}, true},
		{"assertExpectedAndActual_Regexp_success", args{assert.Regexp, emailRegexp, "jdoe@example.com"}, false},
		{"assertExpectedAndActual_Regexp_fail", args{assert.Regexp, emailRegexp, "jdoeexample.com"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AssertExpectedAndActual(tt.args.a, tt.args.expected, tt.args.actual, "Expected %v, but was %v", tt.args.expected, tt.args.actual)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssertExpectedAndActual() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
