package gallifrey_test

import (
	"math/rand"

	. "github.com/ghostlang/gallifrey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func sum(ns []int64) (total int64) {
	for _, n := range ns {
		total += n
	}
	return
}

func randIntn64(n int) int64 {
	return int64(rand.Intn(n))
}

var _ = Describe("Calendar", func() {

	var calendar Calendar

	AssertInterval := func() {

		It("returns the correct slice size", func() {
			b := randIntn64(100)
			slice := calendar.Slice(0, b)
			Ω(slice).Should(HaveLen(int(b)))
		})

		It("starts with 0 when asked for the first interval", func() {
			b := randIntn64(100)
			slice := calendar.Slice(0, b)
			Ω(slice).Should(HaveLen(int(b)))
			Ω(slice[0].Lower()).Should(BeNumerically("==", 0))
		})
	}

	Context("created from delta values", func() {

		var (
			lower  int64
			deltas []int64
		)

		JustBeforeEach(func() {
			lower = randIntn64(1000)
			calendar = NewDeltaCalendar(lower, deltas...)
		})

		AssertDeltas := func() {

			It("maintains the correct deltas", func() {
				l := len(deltas)
				slice := calendar.Slice(0, randIntn64(100))
				for i, interval := range slice {
					Ω(interval.Span()).Should(Equal(deltas[i%l]))
				}

			})

		}

		Context("with a single delta value", func() {

			BeforeEach(func() {
				deltas = []int64{randIntn64(20)}
			})

			AssertInterval()
			AssertDeltas()

		})

		Context("with multiple delta values", func() {

			BeforeEach(func() {
				deltas = []int64{}
				for i := 0; i < rand.Intn(12); i++ {
					deltas = append(deltas, randIntn64(31))
				}
			})

			AssertInterval()
			AssertDeltas()

		})

	})

	/*
		Context("created from another calendar", func() {

			var (
				days            = NewDeltaCalendar(0, 60*60*24)
				from   Calendar = days
				slices []int64
			)

			JustBeforeEach(func() {
				calendar = NewCalendarFromCalendar(from, slices...)
			})

			Context("with a single slice argument", func() {
				BeforeEach(func() {
					slices = []int64{randIntn64(31)}
				})
				AssertInterval()
			})

			Context("with multiple slice arguments", func() {
				BeforeEach(func() {
					slices = []int64{}
					for i := 0; i < rand.Intn(12); i++ {
						slices = append(slices, randIntn64(31))
					}
				})
				AssertInterval()
			})
		})
	*/

})
