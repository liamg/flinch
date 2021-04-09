package core

type Colour [3]int32

var ColourFgFaint = NewColour(128, 128, 128)

func NewColour(r, g, b int32) Colour {
	return Colour([3]int32{r, g, b})
}

func FaintColour(c Colour) Colour {
	if !c.IsDefault() {
		return NewColour(
			c.Red()/2,
			c.Green()/2,
			c.Blue()/2,
		)
	}

	return ColourFgFaint
}

func (c Colour) Red() int32 {
	return c[0]
}

func (c Colour) Green() int32 {
	return c[1]
}

func (c Colour) Blue() int32 {
	return c[2]
}

func (c Colour) IsDefault() bool {
	return c[0] == 0 && c[1] == 0 && c[2] == 0
}
