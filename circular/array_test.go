package circular_test

import (
	. "github.com/ghostlang/gallifrey/circular"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Circular array", func() {

	It("gets by index", func() {
		circle := []int64{0, 1, 2, 3, 4}

		var i, n int64

		for ; i < 5; i++ {
			Ω(Get(circle, i)).Should(BeNumerically("==", n))
			n++
		}

		n = 0
		for ; i < 5; i++ {
			Ω(Get(circle, i)).Should(BeNumerically("==", n))
			n++
		}

	})

	It("sums from 0 to 1", func() {
		circle := []int64{1, 2, 3}
		Ω(Sum(circle, 0, 1)).Should(BeNumerically("==", 1))
	})

	It("sums from 0 to len(circle)", func() {
		circle := []int64{1, 2, 3}
		Ω(Sum(circle, 0, 3)).Should(BeNumerically("==", 6))
	})

	It("sums from 1 to len(circle)+1", func() {
		circle := []int64{1, 2, 3}
		Ω(Sum(circle, 1, 4)).Should(BeNumerically("==", 6))
	})

	It("sums from 0 to 2(len(circle))", func() {
		circle := []int64{1, 2, 3}
		Ω(Sum(circle, 0, 6)).Should(BeNumerically("==", 12))
	})

	It("makes a sum slice starting at 0", func() {
		circle := []int64{1, 2, 3}
		Ω(SumSlice(circle, 0, 0, 5)).Should(BeEquivalentTo([]int64{1, 3, 6, 7, 9}))
	})

	It("makes a sum slice starting at not-0", func() {
		circle := []int64{1, 2, 3}
		Ω(SumSlice(circle, 1, 2, 7)).Should(BeEquivalentTo([]int64{5, 6, 8, 11, 12}))
	})

})
