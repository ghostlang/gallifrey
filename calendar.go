package gallifrey

import (
	"github.com/ghostlang/gallifrey/circular"
)

var (
	//ReferenceTime, _          = time.Parse("Sun Jan 1 00:00:00 GMT 1905")
	Minutes    Calendar = NewDeltaCalendar(0, 60)
	Hours               = NewGroupingCalendar(Minutes, 60)
	Days                = NewGroupingCalendar(Hours, 24)
	Weeks               = NewGroupingCalendar(Days, 7)
	monthsBase          = NewGroupingCalendar(Days, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31)
)

type Calendar interface {
	Get(idx int64) Interval
}

func NewDeltaCalendar(lower int64, deltas ...int64) Calendar {
	return &deltaCalendar{lower, deltas}
}

type deltaCalendar struct {
	lower  int64
	deltas []int64
}

func (c *deltaCalendar) Get(idx int64) Interval {
	lower := c.lower
	if idx > 0 {
		lower += circular.Sum(c.deltas, 0, idx)
	}
	return NewInterval(lower, lower+circular.Get(c.deltas, idx))
}

func NewGroupingCalendar(from Calendar, slices ...int64) Calendar {
	return &groupingCalendar{from, slices}
}

type groupingCalendar struct {
	from   Calendar
	slices []int64
}

func (c *groupingCalendar) Get(idx int64) Interval {
	var x int64
	if idx > 0 {
		x += circular.Sum(c.slices, 0, idx)
	}
	lower := c.from.Get(x).Lower()
	diff := circular.Get(c.slices, idx)
	upper := c.from.Get(x + diff).Lower()
	return NewInterval(lower, upper)
}
