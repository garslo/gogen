package gogen

import (
	"go/ast"
	"go/format"
	"go/token"
	"io"
)

type Package struct {
	Name      string
	Imports   Imports
	Functions Functions
}

// TODO: make this take an Import{}. It's weird right now.
func (me *Package) Import(pkg string) *Package {
	me.Imports.Add(Import{pkg})
	return me
}

func (me *Package) Function(fn Function) *Package {
	me.Functions.Add(fn)
	return me
}

func (me *Package) Ast() ast.Node {
	importDecls := make([]ast.Decl, len(me.Imports))
	for i, imp := range me.Imports {
		importDecls[i] = imp.Ast()
	}
	funcDecls := make([]ast.Decl, len(me.Functions))
	for i, fn := range me.Functions {
		funcDecls[i] = fn.Ast()
	}
	return &ast.File{
		Name: &ast.Ident{
			Name: me.Name,
		},
		Decls: append(importDecls, funcDecls...),
	}
}

func (me *Package) WriteTo(w io.Writer) error {
	fset := token.NewFileSet()
	return format.Node(w, fset, me.Ast())
}
