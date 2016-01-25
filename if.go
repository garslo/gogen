package gogen

import "go/ast"

type If struct {
	Init      Statement
	Condition Expression
	Body      []Statement
}

func (me If) Ast() ast.Stmt {
	var (
		init ast.Stmt
	)
	if me.Init != nil {
		init = me.Init.Ast()
	}
	body := make([]ast.Stmt, len(me.Body))
	for j, stmt := range me.Body {
		body[j] = stmt.Ast()
	}
	return &ast.IfStmt{
		Init: init,
		Cond: me.Condition.Ast(),
		Body: &ast.BlockStmt{
			List: body,
		},
	}
}
