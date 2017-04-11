package gallifrey

import "time"

// Interval represents an interval of integers.
type Interval interface {
	Start() int64
	End() int64

	Contains(Interval) bool
	Overlaps(Interval) bool
}

// TimeInterval is an interval
type TimeInterval interface {
	Start() time.Time
	End() time.Time

	Contains(TimeInterval) bool
	Overlaps(TimeInterval) bool
}

type timeInterval struct {
	start time.Time
	end   time.Time
}

// NewTimeInterval returns a new time interval
func NewTimeInterval(start, end time.Time) TimeInterval {
	if end.Before(start) {
		start, end = end, start
	}
	return &timeInterval{start, end}
}

func (i *timeInterval) Start() time.Time {
	return i.start
}

func (i *timeInterval) End() time.Time {
	return i.end
}

func (i *timeInterval) interval() *interval {
	return &interval{i.start.Unix(), i.end.Unix()}
}

func (i *timeInterval) Contains(other TimeInterval) bool {
	return i.interval().contains(other.Start().Unix(), other.End().Unix())
}

func (i *timeInterval) Overlaps(other TimeInterval) bool {
	return i.interval().overlaps(other.Start().Unix(), other.End().Unix())
}

type interval struct {
	start int64
	end   int64
}

func atOrBefore(a, b int64) bool {
	return a <= b
}

func atOrAfter(a, b int64) bool {
	return a >= b
}

// NewInterval gives you a new interval
func NewInterval(start, end int64) Interval {
	return &interval{
		min(start, end),
		max(start, end),
	}
}

func (i *interval) Start() int64 {
	return i.start
}

func (i *interval) End() int64 {
	return i.end
}

func (i *interval) contains(start, end int64) bool {
	return start >= i.start && end <= i.end
}

func (i *interval) Contains(other Interval) bool {
	return i.contains(other.Start(), other.End())
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

func (i *interval) overlaps(start, end int64) bool {
	return i.start <= end && start <= i.end
}

func (i *interval) Overlaps(other Interval) bool {
	return i.overlaps(other.Start(), other.End())
}
