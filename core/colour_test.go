package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewColour(t *testing.T) {

	inputs := [][3]int32{
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

func TestFaintColour(t *testing.T) {
	testCases := []struct {
		input  [3]int32
		output [3]int32
	}{
		{input: [3]int32{0, 0, 0}, output: [3]int32{128, 128, 128}},
		{input: [3]int32{1, 2, 3}, output: [3]int32{0, 1, 1}},
		{input: [3]int32{255, 255, 0}, output: [3]int32{127, 127, 0}},
		{input: [3]int32{0, 255, 0}, output: [3]int32{0, 127, 0}},
		{input: [3]int32{0, 0, 255}, output: [3]int32{0, 0, 127}},
		{input: [3]int32{128, 64, 32}, output: [3]int32{64, 32, 16}},
		{input: [3]int32{255, 255, 255}, output: [3]int32{127, 127, 127}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d,%d,%d", testCase.input[0], testCase.input[1], testCase.input[2]), func(t *testing.T) {
			colour := FaintColour(NewColour(testCase.input[0], testCase.input[1], testCase.input[2]))
			assert.Equalf(t, testCase.output[0], colour.Red(), "red")
			assert.Equalf(t, testCase.output[1], colour.Green(), "green")
			assert.Equalf(t, testCase.output[2], colour.Blue(), "blue")
		})
	}

}

func TestColourIsDefault(t *testing.T) {
	testCases := []struct {
		input  [3]int32
		output bool
	}{
		{input: [3]int32{0, 0, 0}, output: true},
		{input: [3]int32{1, 2, 3}, output: false},
		{input: [3]int32{255, 255, 0}, output: false},
		{input: [3]int32{0, 255, 0}, output: false},
		{input: [3]int32{0, 0, 255}, output: false},
		{input: [3]int32{128, 64, 32}, output: false},
		{input: [3]int32{255, 255, 255}, output: false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d,%d,%d", testCase.input[0], testCase.input[1], testCase.input[2]), func(t *testing.T) {
			result := NewColour(testCase.input[0], testCase.input[1], testCase.input[2]).IsDefault()
			assert.Equal(t, testCase.output, result)
		})
	}

}
