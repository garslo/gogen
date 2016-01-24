package gogen

import (
	"fmt"
	"go/ast"
	"strconv"
)

type Var struct {
	Name string
}

func (me Var) Ast() ast.Expr {
	return &ast.Ident{
		Name: me.Name,
		Obj: &ast.Object{
			Kind: ast.Var,
			Name: me.Name,
		},
	}
}

// Things that are like Var but either deserve their own name, or have
// slightly different behaviors

func String(value string) Var {
	return Var{fmt.Sprintf(`"%s"`, value)}
}

func Int(value int) Var {
	return Var{strconv.Itoa(value)}
}

func Pkg(value string) Var {
	return Var{value}
}
