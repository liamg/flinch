package core

type Colour [3]uint8

func NewColour(r, g, b uint8) Colour {
	return Colour([3]uint8{r, g, b})
}

func (c Colour) Red() uint8 {
	return c[0]
}

func (c Colour) Green() uint8 {
	return c[1]
}

func (c Colour) Blue() uint8 {
	return c[2]
}
