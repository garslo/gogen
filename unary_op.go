package gogen

import (
	"go/ast"
	"go/token"
)

type Increment struct {
	Value Expression
}

func (me Increment) Ast() ast.Stmt {
	return &ast.IncDecStmt{
		X:   me.Value.Ast(),
		Tok: token.INC,
	}
}

type Decrement struct {
	Value Expression
}

func (me Decrement) Ast() ast.Stmt {
	return &ast.IncDecStmt{
		X:   me.Value.Ast(),
		Tok: token.DEC,
	}
}

type Not struct {
	Value Expression
}

func (me Not) Ast() ast.Expr {
	return &ast.UnaryExpr{
		X:  me.Value.Ast(),
		Op: token.NOT,
	}
}
