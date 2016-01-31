package gogen_test

import (
	"go/ast"

	g "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("For", func() {
	var (
		forStatement g.For
		tree         ast.Stmt
	)

	BeforeEach(func() {
		forStatement = g.For{}
		tree = forStatement.Statement()
	})

	Context("Ast()", func() {
		It("should generate a *ast.ForStmt", func() {
			_, ok := tree.(*ast.ForStmt)
			Expect(ok).To(BeTrue())
		})

		It("should not provide an Init when not given one", func() {
			stmt, _ := tree.(*ast.ForStmt)
			Expect(stmt.Init).To(BeNil())
		})

		It("should not provide a Cond when not given one", func() {
			stmt, _ := tree.(*ast.ForStmt)
			Expect(stmt.Cond).To(BeNil())
		})

		It("should not provide a Post when not given one", func() {
			stmt, _ := tree.(*ast.ForStmt)
			Expect(stmt.Post).To(BeNil())
		})

		It("should not provide a Body when not given one", func() {
			stmt, _ := tree.(*ast.ForStmt)
			Expect(stmt.Body.List).To(HaveLen(0))
		})

		// TODO: should add more
	})
})
