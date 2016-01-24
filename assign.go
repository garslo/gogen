package gogen

import (
	"go/ast"
	"go/token"
)

type Assign struct {
	Lhs Expression
	Rhs Expression
}

func (me Assign) Ast() ast.Stmt {
	return &ast.AssignStmt{
		Tok: token.ASSIGN,
		Lhs: []ast.Expr{me.Lhs.Ast()},
		Rhs: []ast.Expr{me.Rhs.Ast()},
	}
}
