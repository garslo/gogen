package gogen

import "go/ast"

type Statement interface {
	Ast() ast.Stmt
}
