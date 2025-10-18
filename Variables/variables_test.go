package main

import "testing"

func TestVariables(t *testing.T) {
	got := add(1, 2) // add 1 and 2 to equal 3
	want := 3        // expect 3

	if got != want {
		t.Errorf("%q != %q", got, want)
	}
}
