package gogen_test

import (
	"go/ast"
	. "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CallFunction", func() {
	var (
		fnCall CallFunction
		params []Expression
		name   Var
		tree   ast.Stmt
	)

	BeforeEach(func() {
		name = Var{"fooFunction"}
		params = []Expression{
			Var{"x"}, Var{"y"},
		}
		fnCall = CallFunction{name, params}
		tree = fnCall.Ast()
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
