package gogen

import (
	"go/ast"
	"go/token"
)

type Range struct {
	Key          Expression
	Value        Expression
	RangeValue   Expression
	Body         []Statement
	DoNotDeclare bool
}

func (me Range) Ast() ast.Stmt {
	body := make([]ast.Stmt, len(me.Body))
	for i, bodyPart := range me.Body {
		body[i] = bodyPart.Ast()
	}
	var (
		key   Expression = Var{"_"}
		value Expression = Var{"_"}
	)

	if me.Key != nil {
		key = me.Key
	}
	if me.Value != nil {
		value = me.Value
	}
	tok := token.DEFINE
	if me.DoNotDeclare || (me.Key == nil && me.Value == nil) {
		tok = token.ASSIGN
	}

	return &ast.RangeStmt{
		Key:   key.Ast(),
		Value: value.Ast(),
		X:     me.RangeValue.Ast(),
		Tok:   tok,
		Body: &ast.BlockStmt{
			List: body,
		},
	}
}
