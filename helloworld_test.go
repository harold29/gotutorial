package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	got := Hello("Harold")
	want := "Hello, Harold"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
