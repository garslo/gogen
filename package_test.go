package gogen_test

import (
	g "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Package", func() {
	var (
		pkg g.Package
	)

	BeforeEach(func() {
		pkg = g.Package{}
	})

	It("should add functions", func() {
		pkg.Declare(g.Function{})
		Expect(pkg.Declarations).To(HaveLen(1))
	})

	It("should add imports", func() {
		pkg.Declare(g.Import{"os"})
		Expect(pkg.Declarations).To(HaveLen(1))
	})
})
