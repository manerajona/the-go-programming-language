package hello

import "testing"

func TestQuote(t *testing.T) {
	expected := "I can eat glass and it doesn't hurt me."
	if actual := Quote(); actual != expected {
		t.Errorf("Got '%q' but want '%q'", actual, expected)
	}
}

func TestProverb(t *testing.T) {
	expected := "Concurrency is not parallelism."
	if actual := Proverb(); actual != expected {
		t.Errorf("Got '%q' but want '%q'", actual, expected)
	}
}
