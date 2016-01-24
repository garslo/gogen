package gogen_test

import (
	. "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Package", func() {
	var (
		pkg Package
	)

	It("should add functions", func() {
		pkg.Function(Function{})
		Expect(pkg.Functions).To(HaveLen(1))
	})

	It("should add imports", func() {
		pkg.Import("os")
		Expect(pkg.Imports).To(HaveLen(1))
	})
})
