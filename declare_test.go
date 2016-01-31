package gogen_test

import (
	"go/ast"
	"go/token"

	g "github.com/garslo/gogen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Declare", func() {
	var (
		declare  g.Declare
		name     string
		typeName string
		tree     ast.Stmt
	)

	BeforeEach(func() {
		name = "x"
		typeName = "int"
		declare = g.Declare{name, typeName}
		tree = declare.Statement()
	})

	Context("Ast()", func() {
		It("should generate a *ast.DeclStmt", func() {
			_, ok := tree.(*ast.DeclStmt)
			Expect(ok).To(BeTrue())
		})

		It("should have the correct token", func() {
			decl, _ := tree.(*ast.DeclStmt)
			gDecl, _ := decl.Decl.(*ast.GenDecl)
			Expect(gDecl.Tok).To(Equal(token.VAR))
		})

		It("should have the correct var name", func() {
			decl, _ := tree.(*ast.DeclStmt)
			gDecl, _ := decl.Decl.(*ast.GenDecl)
			spec, _ := gDecl.Specs[0].(*ast.ValueSpec)
			Expect(spec.Names).To(HaveLen(1))
			astName := spec.Names[0]
			Expect(astName.Name).To(Equal(name))
			Expect(astName.Obj.Name).To(Equal(name))
			Expect(astName.Obj.Kind).To(Equal(ast.Var))
		})

		It("should have the correct type", func() {
			decl, _ := tree.(*ast.DeclStmt)
			gDecl, _ := decl.Decl.(*ast.GenDecl)
			spec, _ := gDecl.Specs[0].(*ast.ValueSpec)
			t, _ := spec.Type.(*ast.Ident)
			Expect(t.Name).To(Equal(typeName))
		})
	})
})
