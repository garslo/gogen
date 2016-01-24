package gogen

import "go/ast"

type Declaration interface {
	Ast() ast.Decl
}
