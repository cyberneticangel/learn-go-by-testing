/*

---------------------------------------
Writing Tests in Go
---------------------------------------

- File needs to be named foo_test.go
- Test Function must start with the word "Test"
- The Test function takes one argument only t *testing.T
- To use the testing type, "testing" needs to be imported

*/

package main

import "testing"

func TestHello(t *testing.T) {
	got := greeting("World")
	want := "Hello World"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
