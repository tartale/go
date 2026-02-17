package constraintz

import "testing"

func add[N Number](a, b N) N {
	return a + b
}

func TestNumberConstraint_AllowsIntsAndFloats(t *testing.T) {
	if got := add(1, 2); got != 3 {
		t.Fatalf("expected 3, got %v", got)
	}
	if got := add(1.0, 2.5); got != 3.5 {
		t.Fatalf("expected 3.5, got %v", got)
	}
}


