package core

type Colour [3]int32

func NewColour(r, g, b int32) Colour {
	return Colour([3]int32{r, g, b})
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
