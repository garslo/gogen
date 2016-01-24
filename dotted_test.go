package gogen_test

import (
	"go/ast"
	. "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dotted", func() {
	var (
		dot      Dotted
		receiver Var
		name     string
		tree     ast.Expr
	)

	BeforeEach(func() {
		receiver = Var{"receiver"}
		name = "somename"
		dot = Dotted{receiver, name}
		tree = dot.Ast()
	})

	Context("Ast()", func() {
		It("should return a *ast.SelectorExpr", func() {
			_, ok := tree.(*ast.SelectorExpr)
			Expect(ok).To(BeTrue())
		})

		It("should have the correct name", func() {
			expr, _ := tree.(*ast.SelectorExpr)
			Expect(expr.Sel.Name).To(Equal(name))
		})

		//// This is hard to test...the receiver is just a ast.Expr.
		//It("should have the correct receiver", func() {
		//	expr, _ := tree.(*ast.SelectorExpr)
		//	Expect(expr.Sel.Name).To(Equal(name))
		//})
	})

})
