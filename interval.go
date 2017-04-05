package gallifrey

import (
	"time"
)

type Interval interface {
	Start() time.Time
	End() time.Time
	Duration() time.Duration

	Contains(Interval) bool
	Overlaps(Interval) bool
}

type interval struct {
	start time.Time
	end   time.Time
}

func atOrBefore(a, b time.Time) bool {
	return a == b || a.Before(b)
}

func atOrAfter(a, b time.Time) bool {
	return a == b || a.After(b)
}

func NewInterval(start, end time.Time) Interval {
	return &interval{
		MinTime(start, end),
		MaxTime(start, end),
	}
}

func (i *interval) Start() time.Time {
	return i.start
}

func (i *interval) End() time.Time {
	return i.end
}

func (i *interval) Duration() time.Duration {
	return i.end.Sub(i.start)
}

func (i *interval) Contains(other Interval) bool {
	return !i.end.Before(other.End()) && !i.start.After(other.Start())
}

func (i *interval) startsAtOrBefore(other Interval) bool {
	o := other.Start()
	return i.start == o || i.start.Before(o)
}

func (i *interval) endsAtOrAfter(other Interval) bool {
	o := other.End()
	return i.end == o || i.end.After(o)
}

func MaxTime(a, b time.Time) time.Time {
	if a.Before(b) {
		return b
	}
	return a
}

func MinTime(a, b time.Time) time.Time {
	if a.After(b) {
		return b
	}
	return a
}

func (i *interval) Overlaps(other Interval) bool {
	return i.start.Before(other.End()) && i.end.After(other.Start())
}
