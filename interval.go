package gallifrey

type Interval interface {
	Lower() int64
	Upper() int64
	Span() int64
}

// NewInterval returns an interval with the given limits
func NewInterval(l, u int64) Interval {
	if l > u {
		u, l = l, u
	}
	return interval{l, u}
}

// NewIntervalOfSpan returns an interval with the given lower limit and span
func NewIntervalOfSpan(l, s int64) Interval {
	return NewInterval(l, l+s)
}

type interval struct {
	l int64
	u int64
}

func (i interval) Lower() int64 {
	return i.l
}

func (i interval) Upper() int64 {
	return i.u
}

func (i interval) Span() int64 {
	return i.u - i.l
}
