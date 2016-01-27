package gogen

import "go/ast"

type Star struct {
	Value Expression
}

func (me Star) Ast() ast.Expr {
	return &ast.StarExpr{
		X: me.Value.Ast(),
	}
}
