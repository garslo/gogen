package gogen

import "go/ast"

type Dotted struct {
	Receiver Expression
	Name     string
}

func (me Dotted) Ast() ast.Expr {
	return &ast.SelectorExpr{
		X: me.Receiver.Ast(),
		Sel: &ast.Ident{
			Name: me.Name,
		},
	}
}
