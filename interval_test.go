package gallifrey_test

import (
	"math/rand"

	. "github.com/ghostlang/gallifrey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

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

var _ = Describe("An integer interval", func() {

	var (
		start    int64
		end      int64
		duration int64
		interval Interval
	)

	JustBeforeEach(func() {
		start = rand.Int63()
		end = start + duration
		interval = NewInterval(start, end)
	})

	EvaluateInterval := func() {
		It("should return the start boundary given", func() {
			Ω(interval.Start()).Should(Equal(min(start, end)))
		})
		It("should return the end time given", func() {
			Ω(interval.End()).Should(Equal(max(start, end)))
		})
	}

	const (
		zero   int64 = 0
		one    int64 = 1
		negone int64 = -1
	)

	EvaluateComparisons := func() {
		DescribeTable("i1.Contains(i2)",
			func(sdiff, ediff int64, expected bool) {
				Ω(interval.Contains(NewInterval(start+sdiff, end+ediff))).Should(Equal(expected))
			},
			Entry("[(      )]", zero, zero, true),
			Entry("[(    ]  )", zero, one, false),
			Entry("[(    )  ]", zero, negone, true),
			Entry("[  (    )]", one, zero, true),
			Entry("[  (  ]  )", one, one, false),
			Entry("[  (  )  ]", one, negone, true),
			Entry("(  [    )]", negone, zero, false),
			Entry("(  [  ]  )", negone, one, false),
			Entry("(  [  )  ]", negone, negone, false),
		)

		DescribeTable("i1.Overlaps(i2)",
			func(sdiff, ediff int64, nduration int64, expected bool) {
				d := nduration * duration
				Ω(interval.Overlaps(NewInterval(
					start+sdiff+d,
					end+ediff+d,
				))).Should(Equal(expected))
			},
			Entry("[   ](   )", zero, zero, one, true),
			Entry("(   )[   ]", zero, zero, negone, true),
			Entry("[  ]  (  )", one, one, one, false),
			Entry("(  )  [  ]", negone, negone, negone, false),
			Entry("(  [  ]  )", negone, one, zero, true),
			Entry("[  (  )  ]", one, negone, zero, true),
			Entry("[  (  ]  )", one, one, zero, true),
			Entry("(  [  )  ]", negone, negone, zero, true),
			Entry("[(    ]  )", zero, one, zero, true),
			Entry("[(      ])", zero, zero, zero, true),
			Entry("[  (    ])", one, zero, zero, true),
		)
	}

	Context("with start preceding end", func() {
		BeforeEach(func() {
			duration = 60
		})
		EvaluateInterval()
		EvaluateComparisons()
	})

	Context("with end preceding start", func() {
		BeforeEach(func() {
			duration = -60
		})
		EvaluateInterval()
	})

})
