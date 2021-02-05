package components

import (
	"testing"

	"github.com/liamg/flinch/core"
	"github.com/stretchr/testify/assert"
)

func TestColumnLayout(t *testing.T) {

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
					2, 50,
				},
			},
		},
		{
			name:    "1 component, oversized, left justify",
			justify: core.JustifyLeft,
			componentSizes: [][2]int{
				{
					200, 2,
				},
			},
			expectedSizes: [][2]int{
				{
					100, 50,
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
					2, 50,
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
					2, 50,
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
					100, 50,
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
				{2, 50},
				{4, 50},
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
				{2, 50},
				{4, 50},
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
				{2, 50},
				{4, 50},
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
				{49, 50},
				{51, 50},
			},
		},
		{
			name:    "2 components, fill justify, one bigger than half room",
			justify: core.JustifyFill,
			componentSizes: [][2]int{
				{60, 50},
				{4, 50},
			},
			expectedSizes: [][2]int{
				{78, 50},
				{22, 50},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			layout := NewColumnLayout()
			layout.WithJustification(test.justify)
			canvas := newTestCanvas(0, 0, canvasWidth, canvasHeight)

			var components []*testComponent

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

			layout.Render(canvas)

			var usedWidth int
			for i, component := range components {
				cW, cH := component.canvas.Size()
				tc := component.canvas.(*testCanvas)
				cX := tc.x
				cY := tc.y
				assert.Equal(t, test.expectedSizes[i][0], cW)
				assert.Equal(t, test.expectedSizes[i][1], cH)
				switch test.justify {
				case core.JustifyLeft, core.JustifyFill:
					assert.Equal(t, usedWidth, cX)
				case core.JustifyRight:
					var afterWidth int
					for j := i; j < len(components); j++ {
						aW, _ := components[j].canvas.Size()
						afterWidth += aW
					}
					assert.Equal(t, canvasWidth-afterWidth, cX)
				case core.JustifyCenter:
					var allWidth int
					for j := 0; j < len(components); j++ {
						aW, _ := components[j].canvas.Size()
						allWidth += aW
					}
					startX := (canvasWidth - allWidth) / 2
					assert.Equal(t, usedWidth+startX, cX)
				}
				assert.Equal(t, 0, cY)
				usedWidth += cW
			}

		})
	}
}

func TestColumnLayoutDuplicateColumns(t *testing.T) {
	layout := NewColumnLayout()
	canvas := newTestCanvas(0, 0, 100, 50)
	component := newTestComponent(10, 10, canvas)
	layout.Add(component)
	layout.Add(component)
	w, _ := layout.Size()
	assert.Equal(t, 10, w)
}

func TestColumnLayoutRemoveColumn(t *testing.T) {
	layout := NewColumnLayout()
	canvas := newTestCanvas(0, 0, 100, 50)
	component := newTestComponent(10, 10, canvas)
	layout.Add(component)
	layout.Remove(component)
	w, _ := layout.Size()
	assert.Equal(t, 0, w)
}
