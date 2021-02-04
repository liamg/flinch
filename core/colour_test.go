package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColourFunctions(t *testing.T) {

	inputs := [][3]uint8{
		{0, 0, 0},
		{1, 2, 3},
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
		{128, 64, 32},
		{255, 255, 255},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("%d,%d,%d", input[0], input[1], input[2]), func(t *testing.T) {
			colour := NewColour(input[0], input[1], input[2])
			assert.Equal(t, input[0], colour.Red())
			assert.Equal(t, input[1], colour.Green())
			assert.Equal(t, input[2], colour.Blue())
		})
	}

}
