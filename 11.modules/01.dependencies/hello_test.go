package hello

import "testing"

func TestHello(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	if got := Quote(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHelloV3(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proberv(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
