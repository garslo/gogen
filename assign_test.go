package gogen_test

import (
	"go/ast"
	"go/token"
	. "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Assign", func() {
	var (
		assign Assign
		lhs    Var
		rhs    Var
		tree   ast.Stmt
	)

	BeforeEach(func() {
		lhs = Var{"x"}
		rhs = Var{"1"}
		assign = Assign{lhs, rhs}
		tree = assign.Ast()
	})

	Context("Ast()", func() {
		It("should generate a *ast.AssignStmt", func() {
			_, ok := tree.(*ast.AssignStmt)
			Expect(ok).To(BeTrue())
		})

		It("should have the right Tok", func() {
			stmt, _ := tree.(*ast.AssignStmt)
			Expect(stmt.Tok).To(Equal(token.ASSIGN))
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
