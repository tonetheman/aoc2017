package main

import "testing"

// FUCK GO testing package

func TestPart1(t *testing.T) {
	if part1(12) != 3 {
		t.Error("part1(12) failed")
	}

}
