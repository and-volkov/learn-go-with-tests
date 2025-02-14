package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Andrey")

	got := buffer.String()
	want := "Hello, Andrey"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
