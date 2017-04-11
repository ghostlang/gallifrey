package gallifrey

import "time"

// TimeInterval is an interval
type TimeInterval interface {
	Start() time.Time
	End() time.Time

	Contains(TimeInterval) bool
	Overlaps(TimeInterval) bool
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

// NewInterval gives you a new interval
func NewInterval(start, end time.Time) TimeInterval {
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

func (i *interval) Contains(other TimeInterval) bool {
	return !i.end.Before(other.End()) && !i.start.After(other.Start())
}

func (i *interval) startsAtOrBefore(other TimeInterval) bool {
	o := other.Start()
	return i.start == o || i.start.Before(o)
}

func (i *interval) endsAtOrAfter(other TimeInterval) bool {
	o := other.End()
	return i.end == o || i.end.After(o)
}

// MaxTime returns whichever time is larger
func MaxTime(a, b time.Time) time.Time {
	if a.Before(b) {
		return b
	}
	return a
}

// MinTime returns whichever time is lesser
func MinTime(a, b time.Time) time.Time {
	if a.After(b) {
		return b
	}
	return a
}

func (i *interval) Overlaps(other TimeInterval) bool {
	return atOrBefore(i.start, other.End()) && atOrAfter(i.end, other.Start())
}
