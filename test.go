package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testSorting(t *testing.T) {

	io := &IntOrdering[myInt]{}

	intTests := []struct {
		name  string
		o     Ordering[myInt]
		ords  []myInt
		wants []myInt
	}{
		{name: "my-ints", o: io, ords: []myInt{1, 3, 6, 2, 80, 34, 56}, wants: []myInt{1, 2, 3, 634, 56, 80}},
	}

	for _, tt := range intTests {
		t.Run(t.Name(), func(t *testing.T) {
			sorted := Sorted(tt.o, tt.ords)
			assert.Equal(t, sorted, tt.wants, "error")
		})
	}

}
