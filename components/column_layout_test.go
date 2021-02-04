package components

import (
	"testing"

	"github.com/liamg/flinch/core"
	"github.com/stretchr/testify/assert"
)

func TestColumnSizing(t *testing.T) {

	canvasWidth, canvasHeight := 100, 50

	tests := []struct {
		name           string
		justify        core.Justification
		componentSizes [][2]int
		expectedSizes  [][2]int
	}{
		{
			name:    "1 component, left justify",
			justify: core.JustifyLeft,
			componentSizes: [][2]int{
				{
					2, 2,
				},
			},
			expectedSizes: [][2]int{
				{
					2, 2,
				},
			},
		},
		{
			name:    "1 component, right justify",
			justify: core.JustifyRight,
			componentSizes: [][2]int{
				{
					2, 2,
				},
			},
			expectedSizes: [][2]int{
				{
					2, 2,
				},
			},
		},
		{
			name:    "1 component, center justify",
			justify: core.JustifyCenter,
			componentSizes: [][2]int{
				{
					2, 2,
				},
			},
			expectedSizes: [][2]int{
				{
					2, 2,
				},
			},
		},
		{
			name:    "1 component, fill justify",
			justify: core.JustifyFill,
			componentSizes: [][2]int{
				{
					2, 2,
				},
			},
			expectedSizes: [][2]int{
				{
					100, 2,
				},
			},
		},
		{
			name:    "2 components, left justify",
			justify: core.JustifyLeft,
			componentSizes: [][2]int{
				{2, 2},
				{4, 4},
			},
			expectedSizes: [][2]int{
				{2, 2},
				{4, 4},
			},
		},
		{
			name:    "2 components, right justify",
			justify: core.JustifyRight,
			componentSizes: [][2]int{
				{2, 2},
				{4, 4},
			},
			expectedSizes: [][2]int{
				{2, 2},
				{4, 4},
			},
		},
		{
			name:    "2 components, center justify",
			justify: core.JustifyCenter,
			componentSizes: [][2]int{
				{2, 2},
				{4, 4},
			},
			expectedSizes: [][2]int{
				{2, 2},
				{4, 4},
			},
		},
		{
			name:    "2 components, fill justify",
			justify: core.JustifyFill,
			componentSizes: [][2]int{
				{2, 2},
				{4, 4},
			},
			expectedSizes: [][2]int{
				{50, 2},
				{50, 4},
			},
		},
		{
			name:    "2 components, fill justify, one bigger than half room",
			justify: core.JustifyFill,
			componentSizes: [][2]int{
				{60, 2},
				{4, 4},
			},
			expectedSizes: [][2]int{
				{60, 2},
				{40, 4},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			layout := NewColumnLayout()
			canvas := newTestCanvas(canvasWidth, canvasHeight)

			var components []core.Component

			var expectedWidth, expectedHeight int

			for _, componentSize := range test.componentSizes {
				component := newTestComponent(
					componentSize[0],
					componentSize[1],
					canvas,
				)
				expectedWidth += componentSize[0]
				if componentSize[1] > expectedHeight {
					expectedHeight = componentSize[1]
				}
				layout.Add(component)
				components = append(components, component)
			}

			w, h := layout.Size()
			assert.Equal(t, expectedWidth, w)
			assert.Equal(t, expectedHeight, h)

		})
	}
}
