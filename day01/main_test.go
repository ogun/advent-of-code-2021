package main

import "testing"

func TestPart1(t *testing.T) {
	want := 1711
	value := Part1()

	if value != want {
		t.Errorf("	() = %v, want %v", value, want)
	}
}

func TestPart2(t *testing.T) {
	want := 1743
	value := Part2()

	if value != want {
		t.Errorf("Part2() = %v, want %v", value, want)
	}
}
