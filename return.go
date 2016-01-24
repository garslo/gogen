package gogen

import "go/ast"

type Return struct {
	Values []Expression
}

func (me Return) Ast() ast.Stmt {
	ret := make([]ast.Expr, len(me.Values))
	for i, val := range me.Values {
		ret[i] = val.Ast()
	}
	return &ast.ReturnStmt{
		Results: ret,
	}
}
