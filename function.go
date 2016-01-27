package gogen

import "go/ast"

type Receiver struct {
	Name string
	Type Expression
}

func (me Receiver) Ast() *ast.FieldList {
	if me.Type == nil {
		return nil
	}
	return &ast.FieldList{
		List: []*ast.Field{
			&ast.Field{
				Names: []*ast.Ident{
					&ast.Ident{
						Name: me.Name,
						Obj: &ast.Object{
							Kind: ast.Var,
							Name: me.Name,
						},
					},
				},
				Type: me.Type.Ast(),
			},
		},
	}
}

type Function struct {
	Receiver    Receiver
	Name        string
	ReturnTypes Types
	Parameters  Types
	Body        []Statement
}

func (me Function) Ast() ast.Decl {
	paramFields := make([]*ast.Field, len(me.Parameters))
	for j, param := range me.Parameters {
		var names []*ast.Ident
		if param.Name != "" {
			names = []*ast.Ident{
				&ast.Ident{
					Name: param.Name,
					Obj: &ast.Object{
						Kind: ast.Var,
						Name: param.Name,
					},
				},
			}
		}
		paramFields[j] = &ast.Field{
			Names: names,
			Type: &ast.Ident{
				Name: param.TypeName,
			},
		}
	}
	returnFields := make([]*ast.Field, len(me.ReturnTypes))
	for j, ret := range me.ReturnTypes {
		var names []*ast.Ident
		if ret.Name != "" {
			names = []*ast.Ident{
				&ast.Ident{
					Name: ret.Name,
					Obj: &ast.Object{
						Kind: ast.Var,
						Name: ret.Name,
					},
				},
			}
		}
		returnFields[j] = &ast.Field{
			Names: names,
			Type: &ast.Ident{
				Name: ret.TypeName,
			},
		}
	}
	stmts := make([]ast.Stmt, len(me.Body))
	for j, stmt := range me.Body {
		stmts[j] = stmt.Ast()
	}
	return &ast.FuncDecl{
		Recv: me.Receiver.Ast(),
		Name: &ast.Ident{
			Name: me.Name,
			Obj: &ast.Object{
				Kind: ast.Fun,
				Name: me.Name,
			},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: paramFields,
			},
			Results: &ast.FieldList{
				List: returnFields,
			},
		},
		Body: &ast.BlockStmt{
			List: stmts,
		},
	}
}

// Totally unhiegenic
func (me *Function) UpdateReceiver(receiver Receiver) {
	me.Receiver = receiver
}

type Functions []Function

func (me *Functions) Add(fn Function) {
	*me = append(*me, fn)
}
