package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ygor")
	want := "Hello, Ygor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
