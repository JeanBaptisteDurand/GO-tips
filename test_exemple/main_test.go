package main

import (
	"reflect"
	"testing"
)

// Test for Player struct equality
func TestEqualPlayers(t *testing.T) {
	expected := Player{
		name: "Mark",
		hp:   100,
	}
	have := Player{
		name: "Lo",
		hp:   100,
	}

	// Compare struct values
	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %+v but got %+v", expected, have)
	}
}

// Test for calculateValues function
func TestCalculateValues(t *testing.T) {
	var (
		expected = 10
		a        = 5
		b        = 5
	)

	have := calculateValues(a, b)
	if have != expected {
		t.Errorf("expected %d but got %d", expected, have)
	}
}
