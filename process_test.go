package main

import "testing"

func TestProcess(t *testing.T) {
	input, expected := 3, 6
	if output := Process(input); output != expected {
		t.Errorf("expected %d, got %d", expected, output)
	}
}
