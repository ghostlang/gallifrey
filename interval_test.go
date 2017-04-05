package gallifrey_test

import (
	"math"
	"time"

	. "github.com/ghostlang/gallifrey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

const notime = time.Duration(0)

var _ = Describe("An interval", func() {

	var (
		start    time.Time
		end      time.Time
		duration time.Duration
		interval Interval
	)

	JustBeforeEach(func() {
		start = time.Now()
		end = start.Add(duration)
		interval = NewInterval(start, end)
	})

	EvaluateInterval := func() {
		It("should return the start time given", func() {
			Ω(interval.Start()).Should(Equal(MinTime(start, end)))
		})
		It("should return the end time given", func() {
			Ω(interval.End()).Should(Equal(MaxTime(start, end)))
		})
		It("should accurately describe the interval's duration", func() {
			Ω(interval.Duration()).Should(Equal(time.Duration(math.Abs(float64(duration)))))
		})
	}

	EvaluateComparisons := func() {
		DescribeTable("i1.Contains(i2)",
			func(sdiff, ediff time.Duration, expected bool) {
				Ω(interval.Contains(NewInterval(start.Add(sdiff), end.Add(ediff)))).Should(Equal(expected))
			},
			Entry("[(      )]", notime, notime, true),
			Entry("[(    ]  )", notime, time.Second, false),
			Entry("[(    )  ]", notime, -time.Second, true),
			Entry("[  (    )]", time.Second, notime, true),
			Entry("[  (  ]  )", time.Second, time.Second, false),
			Entry("[  (  )  ]", time.Second, -time.Second, true),
			Entry("(  [    )]", -time.Second, notime, false),
			Entry("(  [  ]  )", -time.Second, time.Second, false),
			Entry("(  [  )  ]", -time.Second, -time.Second, false),
		)

		DescribeTable("i1.Overlaps(i2)",
			func(sdiff, ediff time.Duration, expected bool) {
				Ω(interval.Overlaps(NewInterval(start.Add(sdiff), end.Add(ediff)))).Should(Equal(expected))
			},
			Entry("[   ](   )", duration, duration, false),
			Entry("(   )[   ]", -duration, -duration, false),
			Entry("[  ]  (  )", duration+time.Second, duration+time.Second, false),
			Entry("(  )  [  ]", -(duration+time.Second), -(duration+time.Second), false),
			Entry("(  [  ]  )", -time.Second, time.Second, true),
			Entry("[  (  )  ]", time.Second, -time.Second, true),
			Entry("[  (  ]  )", time.Second, time.Second, true),
			Entry("(  [  )  ]", -time.Second, -time.Second, true),
			Entry("[(    ]  )", notime, time.Second, true),
			Entry("[(      ])", notime, notime, true),
			Entry("[  (    ])", time.Second, notime, true),
		)
	}

	Context("with start preceding end", func() {
		BeforeEach(func() {
			duration = time.Minute
		})
		EvaluateInterval()
		EvaluateComparisons()
	})

	Context("with end preceding start", func() {
		BeforeEach(func() {
			duration = -time.Minute
		})
		EvaluateInterval()
	})

	Context("with end equaling start", func() {
		BeforeEach(func() {
			duration = 0
		})
		EvaluateInterval()
		EvaluateComparisons()
	})

})
