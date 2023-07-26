package math

import "testing"

func TestAddition(t *testing.T) {
	want := 12

	ans := Addition(5, 7)

	if ans != want {
		t.Fatalf("got %d, wanted %d", ans, want)
	}

}

func TestSubtraction(t *testing.T) {
	want := 0

	ans := Subtraction(10, 5)

	if ans != want {
		t.Fatalf("got %d, wanted %d", ans, want)
	}
}

func TestMultiplication(t *testing.T) {
	want := 25

	ans := Multiplication(5, 5)

	if ans != want {
		t.Fatalf("got %d, wanted %d", ans, want)
	}
}

func TestDivision(t *testing.T) {
	want := 1.0

	ans := Division(5, 5)

	if ans != want {
		t.Fatalf("got %f, wanted %f", ans, want)
	}
}
