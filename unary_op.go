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
