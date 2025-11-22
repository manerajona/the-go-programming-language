package architecture

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func TestPersonService_Get_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAcc := NewMockAccessor(ctrl)
	mockAcc.EXPECT().Retrieve(1).Return(Person{First: "Alice"})

	ps := NewPersonService(mockAcc)
	p, err := ps.Get(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.First != "Alice" {
		t.Fatalf("expected Alice, got %q", p.First)
	}
}

func TestPersonService_Get_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAcc := NewMockAccessor(ctrl)
	mockAcc.EXPECT().Retrieve(2).Return(Person{})

	ps := NewPersonService(mockAcc)
	_, err := ps.Get(2)
	if err == nil {
		t.Fatalf("expected error for missing person, got nil")
	}
}

func TestPutAndGetFunctions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAcc := NewMockAccessor(ctrl)

	// For Put we expect Save to be called with these args
	mockAcc.EXPECT().Save(3, Person{First: "Bob"}).Times(1)

	Put(mockAcc, 3, Person{First: "Bob"})

	// For Get free function we expect Retrieve to be called and return a Person
	mockAcc.EXPECT().Retrieve(3).Return(Person{First: "Bob"}).Times(1)
	got := Get(mockAcc, 3)
	if got.First != "Bob" {
		t.Fatalf("expected Bob, got %q", got.First)
	}
}
