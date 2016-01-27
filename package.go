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
	Structs   Structs
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

func (me *Package) Struct(s Struct) *Package {
	me.Structs.Add(s)
	return me
}

func (me *Package) Ast() ast.Node {
	decls := []ast.Decl{}
	for _, imp := range me.Imports {
		decls = append(decls, imp.Ast())
	}
	for _, st := range me.Structs {
		decls = append(decls, st.Ast())
		for _, method := range st.Methods {
			decls = append(decls, method.Ast())
		}
	}
	for _, fn := range me.Functions {
		decls = append(decls, fn.Ast())
	}
	return &ast.File{
		Name: &ast.Ident{
			Name: me.Name,
		},
		Decls: decls,
	}
}

func (me *Package) WriteTo(w io.Writer) error {
	fset := token.NewFileSet()
	return format.Node(w, fset, me.Ast())
}
