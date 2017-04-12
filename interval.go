package gallifrey

import (
	"sort"
)

// Interval represents an interval of integers
type Interval interface {
	Start() int64
	End() int64
	LessThan(Interval) bool
	GreaterThan(Interval) bool
	Gap(Interval) int64
	Adjacent(Interval, int64) bool
	StartsBefore(Interval) bool
	EndsAfter(Interval) bool
	Contains(Interval) bool
	Overlaps(Interval) bool
	Extend(...Interval) Interval
}

// NewInterval creates a new interval wrapping the range specified
func NewInterval(start, end int64) Interval {
	return &interval{
		min(start, end),
		max(start, end),
	}
}

type interval struct {
	start int64
	end   int64
}

func (i *interval) Start() int64 {
	return i.start
}

func (i *interval) End() int64 {
	return i.end
}

func (i *interval) Contains(other Interval) bool {
	return i.contains(other.Start(), other.End())
}

func (i *interval) Overlaps(other Interval) bool {
	return i.overlaps(other.Start(), other.End())
}

func (i *interval) StartsBefore(other Interval) bool {
	return i.Start() < other.Start()
}

func (i *interval) EndsAfter(other Interval) bool {
	return i.End() > other.End()
}

func (i *interval) LessThan(other Interval) bool {
	return i.End() < other.Start()
}

func (i *interval) GreaterThan(other Interval) bool {
	return i.Start() > other.End()
}

func (i *interval) Gap(other Interval) int64 {
	if i.Overlaps(other) {
		return 0
	}
	if i.LessThan(other) {
		return other.Start() - i.End()
	}
	return i.Start() - other.End()
}

func (i *interval) Adjacent(other Interval, margin int64) bool {
	return i.Gap(other) <= margin
}

func (i *interval) Extend(intervals ...Interval) Interval {
	is := append([]Interval{i}, intervals...)
	sort.Slice(is, func(i, j int) bool {
		return is[i].Start() < is[j].Start()
	})
	s := is[0].Start()
	var e int64
	for _, ival := range is {
		if e < ival.End() {
			e = ival.End()
		}
	}
	return &interval{s, e}
}

func (i *interval) contains(start, end int64) bool {
	return start >= i.start && end <= i.end
}

func (i *interval) overlaps(start, end int64) bool {
	return i.start <= end && start <= i.end
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
