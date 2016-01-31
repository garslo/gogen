package gogen_test

import (
	"go/ast"

	g "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CallFunction", func() {
	var (
		fnCall g.CallFunction
		params []g.Expression
		name   g.Var
		tree   ast.Stmt
	)

	BeforeEach(func() {
		name = g.Var{"fooFunction"}
		params = []g.Expression{
			g.Var{"x"}, g.Var{"y"},
		}
		fnCall = g.CallFunction{name, params}
		tree = fnCall.Statement()
	})

	Context("Ast()", func() {
		It("should generate a &ast.ExprStmt", func() {
			_, ok := tree.(*ast.ExprStmt)
			Expect(ok).To(BeTrue())
		})

		It("should have the correct number of params", func() {
			stmt, _ := tree.(*ast.ExprStmt)
			expr, _ := stmt.X.(*ast.CallExpr)
			Expect(expr.Args).To(HaveLen(len(params)))
		})
	})
})
