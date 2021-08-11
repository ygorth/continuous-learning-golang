package main

import "testing"

// subtests around Hello()
func TestHello(t *testing.T) {
	t.Run("saying hello to orcs", func(t *testing.T) {
		got := Hello("Ygor")
		want := "Hello, Ygor"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, Ygor' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Ygor"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
