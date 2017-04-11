package gallifrey

import "time"

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
