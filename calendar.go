package gallifrey

func sum(ns []int64) int64 {
	var total int64
	for _, n := range ns {
		total += n
	}
	return total
}

type Calendar interface {
	Slice(a, b int64) []Interval
}

func NewDeltaCalendar(lower int64, deltas ...int64) Calendar {
	var sum int64
	for _, i := range deltas {
		sum += i
	}
	return &deltaCalendar{lower, deltas, sum}
}

func NewCalendarFromCalendar(from Calendar, slices ...int64) Calendar {
	return &superCalendar{from, slices, sum(slices)}
}

type deltaCalendar struct {
	lower  int64
	deltas []int64
	sum    int64
}

type superCalendar struct {
	from   Calendar
	slices []int64
	sum    int64
}

func circularSum(vals []int64, idx int64) int64 {
	s := sum(vals)
	l := int64(len(vals))
	return s*(idx/l) + sum(vals[:idx%l])
}

func (c *deltaCalendar) Slice(a, b int64) (result []Interval) {
	l := int64(len(c.deltas))
	last := circularSum(c.deltas, a)
	for i := a; i < b; i++ {
		next := last + c.deltas[i%l]
		result = append(result, NewInterval(last, next))
		last = next
	}
	return
}

func (c *superCalendar) Slice(a, b int64) (result []Interval) {
	return []Interval{}
}
