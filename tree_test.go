package gallifrey_test

import (
	. "github.com/ghostlang/gallifrey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tree", func() {

	t := NewIntervalTree()

	It("should allow insert", func() {
		t.Insert(NewInterval(10, 20))
		Ω(t.Contains(NewInterval(10, 20))).Should(BeTrue())
		Ω(t.Contains(NewInterval(15, 15))).Should(BeTrue())
		Ω(t.Contains(NewInterval(15, 18))).Should(BeTrue())
		Ω(t.Contains(NewInterval(9, 19))).Should(BeFalse())
		Ω(t.Contains(NewInterval(0, 5))).Should(BeFalse())

		t.Insert(NewInterval(5, 9))
		Ω(t.Contains(NewInterval(0, 5))).Should(BeFalse())
		Ω(t.Contains(NewInterval(9, 19))).Should(BeTrue())
	})

})
