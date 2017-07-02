package gallifrey_test

import (
	"math/rand"

	. "github.com/ghostlang/gallifrey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Interval", func() {

	rand.Seed(GinkgoRandomSeed())

	var (
		interval Interval
		lower    int64
		upper    int64
		span     int64
	)

	BeforeEach(func() {
		lower = int64(rand.Intn(1000))
		span = int64(rand.Intn(100))
		upper = lower + span
	})

	AssertIntervalConsistency := func() {
		It("should return the given lower limit", func() {
			Ω(interval.Lower()).Should(Equal(lower))
		})
		It("should return the given upper limit", func() {
			Ω(interval.Upper()).Should(Equal(upper))
		})
		It("should return the correct span", func() {
			Ω(interval.Span()).Should(Equal(span))
		})
		It("should return lower less than upper", func() {
			Ω(interval.Lower()).Should(BeNumerically("<", interval.Upper()))
		})
	}

	Context("created from limits", func() {

		Context("with the first limit lower than the second", func() {
			JustBeforeEach(func() {
				interval = NewInterval(lower, upper)
			})
			AssertIntervalConsistency()
		})

		Context("with the first limit higher than the second", func() {
			JustBeforeEach(func() {
				interval = NewInterval(upper, lower)
			})
			AssertIntervalConsistency()
		})
	})

	Context("created from a lower limit and a span", func() {

		Context("with a positive span", func() {
			JustBeforeEach(func() {
				interval = NewIntervalOfSpan(lower, span)
			})
			AssertIntervalConsistency()
		})

		Context("with a negative span", func() {
			JustBeforeEach(func() {
				interval = NewIntervalOfSpan(upper, -span)
			})
			AssertIntervalConsistency()
		})
	})
})
