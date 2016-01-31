package gogen_test

import (
	"go/ast"
	"go/token"

	g "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Import", func() {
	var (
		imp  g.Import
		name string
		tree ast.Decl
	)

	BeforeEach(func() {
		name = "someimport"
		imp = g.Import{name}
		tree = imp.Declaration()
	})

	Context("Ast()", func() {
		It("should have the correct types all the way down", func() {
			gDecl, ok := tree.(*ast.GenDecl)
			Expect(ok).To(BeTrue())
			Expect(gDecl.Specs).To(HaveLen(1))
			_, ok = gDecl.Specs[0].(*ast.ImportSpec)
			Expect(ok).To(BeTrue())
		})

		It("should have an import token", func() {
			gDecl, _ := tree.(*ast.GenDecl)
			Expect(gDecl.Tok).To(Equal(token.IMPORT))
		})

		It("should have quotes around it's name", func() {
			gDecl, _ := tree.(*ast.GenDecl)
			spec, _ := gDecl.Specs[0].(*ast.ImportSpec)
			Expect(spec.Path.Kind).To(Equal(token.STRING))
			Expect(spec.Path.Value).To(Equal(`"someimport"`))
		})

		// TODO: named imports, etc.
	})
})
