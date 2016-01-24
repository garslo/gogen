package gogen

import "go/ast"

type Var struct {
	Name string
}

func (me Var) Ast() ast.Expr {
	return &ast.Ident{
		Name: me.Name,
		Obj: &ast.Object{
			Kind: ast.Var,
			Name: me.Name,
		},
	}
}
