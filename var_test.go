package gogen_test

import (
	"go/ast"

	g "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Var", func() {
	var (
		name string
		v    g.Var
		tree ast.Expr
	)

	BeforeEach(func() {
		v = g.Var{name}
		tree = v.Expression()
	})

	Context("ast", func() {
		It("should generate an *ast.Ident", func() {
			_, ok := tree.(*ast.Ident)
			Expect(ok).To(BeTrue())
		})

		It("should have the right name", func() {
			ident, _ := tree.(*ast.Ident)
			Expect(ident.Name).To(Equal(name))
		})

		It("should have the right object name", func() {
			ident, _ := tree.(*ast.Ident)
			Expect(ident.Obj.Name).To(Equal(name))
		})

		It("should have the right object kind", func() {
			ident, _ := tree.(*ast.Ident)
			Expect(ident.Obj.Kind).To(Equal(ast.Var))
		})
	})
})
