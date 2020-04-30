package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("liukun01")
	want := "Hello, world liukun01"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
