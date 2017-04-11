package gallifrey

// Interval represents an interval of integers
type Interval interface {
	Start() int64
	End() int64

	Contains(Interval) bool
	Overlaps(Interval) bool
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
