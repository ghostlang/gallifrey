package gallifrey_test

import (
	"math/rand"

	. "github.com/ghostlang/gallifrey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func randIntn64(n int) int64 {
	return int64(rand.Intn(n))
}

func randArray(minSize, maxSize, minVal, maxVal int) []int64 {
	size := randIntn64(maxSize-minSize) + int64(minSize)
	result := make([]int64, size)
	for i := int64(0); i < size; i++ {
		result[i] = randIntn64(maxVal-minVal) + int64(minVal)
	}
	return result
}

var _ = Describe("Calendar", func() {

	var calendar Calendar

	Context("created from delta values", func() {

		var (
			lower  int64
			deltas []int64
		)

		JustBeforeEach(func() {
			lower = randIntn64(100)
			deltas = randArray(3, 10, 1, 10)
			calendar = NewDeltaCalendar(lower, deltas...)
		})

		It("can get the first item", func() {
			interval := calendar.Get(0)
			Ω(interval.Lower()).Should(Equal(lower))
			Ω(interval.Upper()).Should(Equal(interval.Lower() + deltas[0]))
		})

		It("can get the second item", func() {
			interval := calendar.Get(1)
			Ω(interval.Lower()).Should(Equal(lower + deltas[0]))
			Ω(interval.Upper()).Should(Equal(interval.Lower() + deltas[1]))
		})

		It("can get the third item", func() {
			interval := calendar.Get(2)
			Ω(interval.Lower()).Should(Equal(lower + deltas[0] + deltas[1]))
			Ω(interval.Upper()).Should(Equal(interval.Lower() + deltas[2]))
		})

	})

	Context("created from another calendar", func() {

		var (
			lower          int64
			deltas, slices []int64
			from, calendar Calendar
		)

		BeforeEach(func() {
			lower = 0
			deltas = []int64{2, 6, 1, 8, 3}
			from = NewDeltaCalendar(lower, deltas...)

			slices = []int64{2, 6, 7, 6}
			calendar = NewGroupingCalendar(from, slices...)
		})

		It("can get the first item", func() {
			interval := calendar.Get(0)
			Ω(interval.Lower()).Should(BeNumerically("==", 0))
			Ω(interval.Upper()).Should(BeNumerically("==", 8))
		})

		It("can get the second item", func() {
			interval := calendar.Get(1)
			Ω(interval.Lower()).Should(BeNumerically("==", 8))
			Ω(interval.Upper()).Should(BeNumerically("==", 29))
		})

		It("can get the third item", func() {
			interval := calendar.Get(2)
			Ω(interval.Lower()).Should(BeNumerically("==", 29))
			Ω(interval.Upper()).Should(BeNumerically("==", 60))
		})

	})
})
