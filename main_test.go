package main

import (
	"reflect"
	"testing"
)

func TestBuildNextGenGrid(t *testing.T) {
	initialGrid := Grid{
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, true, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}

	expectedNextGenGrid := Grid{
		{false, false, true, false, false},
		{false, false, true, true, false},
		{false, true, false, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	nextGenGrid := NextGeneration(initialGrid)

	if !reflect.DeepEqual(nextGenGrid, expectedNextGenGrid) {
		t.Errorf("Next generation grid is incorrect.\nGot:\n%v\nExpected:\n%v", nextGenGrid, expectedNextGenGrid)
	}
}
