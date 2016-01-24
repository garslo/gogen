package gogen

import (
	"go/ast"
	"go/token"
)

type DeclareAndAssign struct {
	Lhs Expression
	Rhs Expression
}

func (me DeclareAndAssign) Ast() ast.Stmt {
	return &ast.AssignStmt{
		Tok: token.DEFINE,
		Lhs: []ast.Expr{me.Lhs.Ast()},
		Rhs: []ast.Expr{me.Rhs.Ast()},
	}
}
