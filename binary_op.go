package gogen

import (
	"go/ast"
	"go/token"
)

type LessThan struct {
	Lhs Expression
	Rhs Expression
}

func (me LessThan) Ast() ast.Expr {
	return &ast.BinaryExpr{
		Op: token.LSS,
		X:  me.Lhs.Ast(),
		Y:  me.Rhs.Ast(),
	}
}

type LessThanOrEqual struct {
	Lhs Expression
	Rhs Expression
}

func (me LessThanOrEqual) Ast() ast.Expr {
	return &ast.BinaryExpr{
		Op: token.LEQ,
		X:  me.Lhs.Ast(),
		Y:  me.Rhs.Ast(),
	}
}

type GreaterThan struct {
	Lhs Expression
	Rhs Expression
}

func (me GreaterThan) Ast() ast.Expr {
	return &ast.BinaryExpr{
		Op: token.GTR,
		X:  me.Lhs.Ast(),
		Y:  me.Rhs.Ast(),
	}
}

type GreaterThanOrEqual struct {
	Lhs Expression
	Rhs Expression
}

func (me GreaterThanOrEqual) Ast() ast.Expr {
	return &ast.BinaryExpr{
		Op: token.GEQ,
		X:  me.Lhs.Ast(),
		Y:  me.Rhs.Ast(),
	}
}

type Equals struct {
	Lhs Expression
	Rhs Expression
}

func (me Equals) Ast() ast.Expr {
	return &ast.BinaryExpr{
		Op: token.EQL,
		X:  me.Lhs.Ast(),
		Y:  me.Rhs.Ast(),
	}
}
