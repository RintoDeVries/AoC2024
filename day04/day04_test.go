package main

import "testing"

func TestSearchDirections(t *testing.T) {
	unpadded_input, err := parseInput("test_search_directions.txt")
	if err != nil {
		panic(err)
	}

	got := part1(unpadded_input)
	want := 8
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPaddingPt1(t *testing.T) {
	unpadded_input, err := parseInput("test_padding.txt")
	if err != nil {
		panic(err)
	}
	got := part1(unpadded_input)
	want := 12
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func TestPaddingPt2(t *testing.T) {
	unpadded_input, err := parseInput("test_padding_part2.txt")
	if err != nil {
		panic(err)
	}
	got := part2(unpadded_input)
	want := 4
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMyTestInput(t *testing.T) {
	unpadded_input, err := parseInput("test_input.txt")
	if err != nil {
		panic(err)
	}
	got := part1(unpadded_input)
	want := 18
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMyTestInputPt2(t *testing.T) {
	unpadded_input, err := parseInput("test_input_part2.txt")
	if err != nil {
		panic(err)
	}
	got := part2(unpadded_input)
	want := 9
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
