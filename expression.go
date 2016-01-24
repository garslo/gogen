package gogen

import "go/ast"

type Expression interface {
	Ast() ast.Expr
}
