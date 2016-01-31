package gogen_test

import (
	"go/ast"
	"go/token"

	g "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeclareAndAssign", func() {
	var (
		declAndAssign g.DeclareAndAssign
		lhs           g.Var
		rhs           g.Var
		tree          ast.Stmt
	)
	BeforeEach(func() {
		lhs = g.Var{"x"}
		rhs = g.Var{"1"}
		declAndAssign = g.DeclareAndAssign{lhs, rhs}
		tree = declAndAssign.Statement()
	})

	Context("Ast()", func() {
		It("should generate a *ast.AssignStmt", func() {
			_, ok := tree.(*ast.AssignStmt)
			Expect(ok).To(BeTrue())
		})

		It("should have the right Tok", func() {
			stmt, _ := tree.(*ast.AssignStmt)
			Expect(stmt.Tok).To(Equal(token.DEFINE))
		})

		It("should have the correct number of lhs", func() {
			// TODO: test this with > 1 lhs, rhs
			stmt, _ := tree.(*ast.AssignStmt)
			Expect(stmt.Lhs).To(HaveLen(1))
		})

		It("should have the correct number of rhs", func() {
			// TODO: test this with > 1 lhs, rhs
			stmt, _ := tree.(*ast.AssignStmt)
			Expect(stmt.Lhs).To(HaveLen(1))
		})
	})
})
