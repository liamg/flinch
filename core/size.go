package core

type Size struct {
	W, H int
}

func (s Size) Add(a Size) Size {
	return Size{
		W: s.W + a.W,
		H: s.H + a.H,
	}
}

func (s Size) Minus(a Size) Size {
	return Size{
		W: s.W - a.W,
		H: s.H - a.H,
	}
}

type SizeStrategy func(available Size, required Size) Size

func SizeStrategyMinimum() SizeStrategy {
	return func(available Size, required Size) Size {
		return required
	}
}

func SizeStrategyMaximum() SizeStrategy {
	return func(available Size, required Size) Size {
		return available
	}
}

func SizeStrategyMaximumWidth() SizeStrategy {
	return func(available Size, required Size) Size {
		return Size{
			W: available.W,
			H: required.H,
		}
	}
}

func SizeStrategyMaximumHeight() SizeStrategy {
	return func(available Size, required Size) Size {
		return Size{
			W: required.W,
			H: available.H,
		}
	}
}

func SizeStrategyPercentage(wPc, hPc int) SizeStrategy {
	return func(available Size, required Size) Size {
		size := Size{
			W: (available.W * wPc) / 100,
			H: (available.H * hPc) / 100,
		}
		if size.W < required.W {
			size.W = required.W
		}
		if size.H < required.H {
			size.H = required.H
		}
		return size
	}
}

func SizeStrategyAtLeast(min Size) SizeStrategy {
	return func(available Size, required Size) Size {
		if available.W < min.W {
			available.W = min.W
		}
		if available.H < min.H {
			available.H = min.H
		}
		return available
	}
}

func SizeStrategyAtMost(max Size) SizeStrategy {
	return func(available Size, required Size) Size {
		if available.W > max.W {
			available.W = max.W
		}
		if available.H > max.H {
			available.H = max.H
		}
		return available
	}
}

func SizeStrategyMultiple(strategies ...SizeStrategy) SizeStrategy {
	return func(available Size, required Size) Size {
		for _, strat := range strategies {
			available = strat(available, required)
		}
		return available
	}
}

type StrategicSizer interface {
	SetSizeStrategy(strategy SizeStrategy)
	GetSizeStrategy() SizeStrategy
}

type Sizer struct {
	strategy SizeStrategy
}

func (s *Sizer) SetSizeStrategy(strategy SizeStrategy) {
	s.strategy = strategy
}

func (s Sizer) GetSizeStrategy() SizeStrategy {
	if s.strategy == nil {
		return SizeStrategyMinimum()
	}
	return s.strategy
}

func CalculateSize(component Component, available Size) Size {

	if sizer, ok := component.(StrategicSizer); ok {
		return sizer.GetSizeStrategy()(available, component.MinimumSize())
	}

	return component.MinimumSize()
}
