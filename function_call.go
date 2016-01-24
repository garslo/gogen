package gogen

import "go/ast"

type CallFunction struct {
	Func   Expression
	Params []Expression
}

func (me CallFunction) Ast() ast.Stmt {
	params := make([]ast.Expr, len(me.Params))
	for i, param := range me.Params {
		params[i] = param.Ast()
	}
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun:  me.Func.Ast(),
			Args: params,
		},
	}
}
