package main

import "testing"

func TestSearchDirections(t *testing.T) {
	got := part1("test_search_directions.txt")
	want := 8
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPadding(t *testing.T) {
	got := part1("test_padding.txt")
	want := 12
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMyTestInput(t *testing.T) {
	got := part1("test_input.txt")
	want := 18
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
