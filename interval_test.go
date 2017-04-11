package gallifrey_test

import (
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
		interval TimeInterval
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
			func(sdiff, ediff time.Duration, nduration int, expected bool) {
				d := time.Duration(nduration) * duration
				Ω(interval.Overlaps(NewInterval(
					start.Add(sdiff+d),
					end.Add(ediff+d),
				))).Should(Equal(expected))
			},
			Entry("[   ](   )", notime, notime, 1, true),
			Entry("(   )[   ]", notime, notime, -1, true),
			Entry("[  ]  (  )", time.Second, time.Second, 1, false),
			Entry("(  )  [  ]", -time.Second, -time.Second, -1, false),
			Entry("(  [  ]  )", -time.Second, time.Second, 0, true),
			Entry("[  (  )  ]", time.Second, -time.Second, 0, true),
			Entry("[  (  ]  )", time.Second, time.Second, 0, true),
			Entry("(  [  )  ]", -time.Second, -time.Second, 0, true),
			Entry("[(    ]  )", notime, time.Second, 0, true),
			Entry("[(      ])", notime, notime, 0, true),
			Entry("[  (    ])", time.Second, notime, 0, true),
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

})
